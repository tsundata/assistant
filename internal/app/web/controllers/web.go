package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/skip2/go-qrcode"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/web"
	"github.com/tsundata/assistant/internal/app/web/components"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
	"github.com/tsundata/assistant/internal/pkg/vendors/pocket"
	"github.com/valyala/fasthttp"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type WebController struct {
	opt       *web.Options
	rdb       *redis.Client
	logger    *zap.Logger
	midClient pb.MiddleClient
	msgClient pb.MessageClient
	wfClient  pb.WorkflowClient
}

func NewWebController(opt *web.Options, rdb *redis.Client, logger *zap.Logger,
	midClient pb.MiddleClient, msgClient pb.MessageClient, wfClient pb.WorkflowClient) *WebController {
	return &WebController{opt: opt, rdb: rdb, logger: logger, midClient: midClient, msgClient: msgClient, wfClient: wfClient}
}

func (wc *WebController) Index(c *fasthttp.RequestCtx) {
	c.Response.SetBody([]byte("Web"))
}

func (wc *WebController) Echo(c *fasthttp.RequestCtx) {
	text := c.FormValue("text")
	c.Response.SetBody(text)
}

func (wc *WebController) Robots(c *fasthttp.RequestCtx) {
	txt := `User-agent: *
Disallow: /`

	c.Response.SetBody(utils.StringToByte(txt))
}

func (wc *WebController) Page(c *fasthttp.RequestCtx) {
	pageRe := regexp.MustCompile(`([\w\-]+)$`)
	r := pageRe.FindSubmatch(c.Path())
	if len(r) < 1 {
		c.Response.SetStatusCode(http.StatusNotFound)
		return
	}

	reply, err := wc.midClient.GetPage(context.Background(), &pb.PageRequest{
		Uuid: utils.ByteToString(r[0]),
	})
	if err != nil {
		wc.logger.Error(err.Error())
		c.Response.SetStatusCode(http.StatusNotFound)
		return
	}
	if reply.GetContent() == "" {
		c.Response.SetStatusCode(http.StatusNotFound)
		return
	}

	var list []string
	err = json.Unmarshal([]byte(reply.GetContent()), &list)
	if err != nil {
		wc.logger.Error(err.Error())
		c.Response.SetStatusCode(http.StatusNotFound)
		return
	}

	var items []components.Component

	for _, item := range list {
		re, _ := regexp.Compile(utils.UrlRegex)
		s := re.FindString(item)
		if s != "" {
			item = strings.ReplaceAll(item, s, fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, s, s))
		}
		item = strings.ReplaceAll(item, "\n", "<br>")

		items = append(items, &components.Text{
			Title: item,
		})
	}

	comp := components.Html{
		Title: reply.GetTitle(),
		Page: &components.Page{
			Title: reply.GetTitle(),
			Content: &components.List{
				Items: items,
			},
		},
	}

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) Qr(c *fasthttp.RequestCtx) {
	path := c.URI().PathOriginal()
	qrRe := regexp.MustCompile(`^/qr/(.*)$`)
	r := qrRe.FindSubmatch(path)
	if len(r) < 1 {
		c.Response.SetStatusCode(http.StatusNotFound)
		return
	}

	txt, err := url.QueryUnescape(utils.ByteToString(r[1]))
	if err != nil {
		wc.logger.Error(err.Error())
		c.Response.SetBodyString("error text")
		return
	}

	png, err := qrcode.Encode(txt, qrcode.Medium, 512)
	if err != nil {
		wc.logger.Error(err.Error())
		c.Response.SetBodyString("error qr")
		return
	}

	c.Response.Header.Set("Content-Type", "image/png")
	c.Response.SetBody(png)
}

func (wc *WebController) Apps(c *fasthttp.RequestCtx) {
	uuid := utils.ExtractUUID(utils.ByteToString(c.Path()))
	var items []components.Component

	reply, err := wc.midClient.Apps(context.Background(), &pb.TextRequest{})
	if err != nil {
		wc.logger.Error(err.Error())
		c.Response.SetStatusCode(http.StatusNotFound)
		return
	}

	for _, app := range reply.GetApps() {
		authStr := "Unauthorized"
		authorizedURL := fmt.Sprintf("/app/%s?uuid=%s", app.GetType(), uuid)
		if app.GetIsAuthorized() {
			authStr = "Authorized"
			authorizedURL = "javascript:void(0);"
		}
		items = append(items, &components.App{
			Name: app.GetTitle(),
			Icon: "rocket",
			Text: fmt.Sprintf("%s (%s)", app.GetType(), authStr),
			URL:  authorizedURL,
		})
	}

	comp := components.Html{
		Title:   "Apps",
		UseIcon: true,
		Page: &components.Page{
			Title: "Apps",
			Content: &components.List{
				Items: items,
			},
		},
	}

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) Memo(c *fasthttp.RequestCtx) {
	var items []components.Component

	reply, err := wc.msgClient.List(context.Background(), &pb.MessageRequest{})
	if err != nil {
		wc.logger.Error(err.Error())
		c.Response.SetStatusCode(http.StatusNotFound)
		return
	}

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.Linkify,
			extension.Table,
			extension.TaskList,
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	var buf bytes.Buffer
	for _, item := range reply.GetMessages() {
		// markdown
		text := item.GetText()
		buf.Reset()
		err := md.Convert(utils.StringToByte(item.GetText()), &buf)
		if err != nil {
			wc.logger.Error(err.Error())
		} else {
			text = buf.String()
		}

		items = append(items, &components.Memo{
			Time: item.GetTime(),
			Content: &components.Text{
				Title: text,
			},
		})
	}

	comp := components.Html{
		Title: "Memo",
		Page: &components.Page{
			Title: "Memo",
			Content: &components.List{
				Items: items,
			},
		},
	}

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) Credentials(c *fasthttp.RequestCtx) {
	var items []components.Component

	reply, err := wc.midClient.GetCredentials(context.Background(), &pb.TextRequest{})
	if err != nil {
		wc.logger.Error(err.Error())
		c.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	for _, item := range reply.GetItems() {
		items = append(items, &components.LinkButton{
			Title: item.GetKey(),
			Name:  item.GetValue(),
			URL:   "javascript:void(0)",
		})
	}

	comp := components.Html{
		Title: "Credentials",
		Page: &components.Page{
			Title: "Credentials",
			Action: &components.Link{
				Title: "Add Credentials",
				URL:   fmt.Sprintf("/credentials/%s/create", utils.ExtractUUID(utils.ByteToString(c.Path()))),
			},
			Content: &components.List{
				Items: items,
			},
		},
	}

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) CredentialsCreate(c *fasthttp.RequestCtx) {
	uuid := utils.ExtractUUID(utils.ByteToString(c.Path()))
	options := map[string]interface{}{
		"github": map[string]string{
			"client_id":     "Client ID",
			"client_secret": "Client secrets",
		},
		"pocket": map[string]string{
			"consumer_key": "Consumer Key",
		},
		"pushover": map[string]string{
			"token": "API Token",
			"user":  "User Key",
		},
	}

	selectOption := make(map[string]string)
	selectOption[""] = "-"
	for k := range options {
		selectOption[k] = strings.Title(k)
	}

	var items []components.Component
	items = append(items, &components.Input{
		Name:  "name",
		Title: "Name",
		Type:  "text",
	})
	items = append(items, &components.Select{
		Name:  "type",
		Title: "Type",
		Value: selectOption,
	})
	comp := components.Html{
		Title: "Create Credentials",
		Page: &components.Page{
			Title: "Create Credentials",
			Action: &components.Link{
				Title: "Go Back",
				URL:   fmt.Sprintf("/credentials/%s", uuid),
			},
			Content: &components.Form{
				Action: fmt.Sprintf("/credentials/%s/store", uuid),
				Method: "POST",
				Inputs: items,
			},
		},
	}

	d, _ := json.Marshal(options)
	h := "`<div class='input option-input'>\n<label for='input-${key}'>${options[e.target.value][key]}:</label>\n<input type='text' id='input-${key}' name='${key}'>\n</div>`"

	comp.SetJs(template.JS(fmt.Sprintf(`const options = %s
    document.querySelector("select[name=type]").addEventListener("change", function (e) {
        if (e.target.value !== "") {
            let o = ""
            Object.keys(options[e.target.value]).forEach(function (key) {
                o += %s
            })
            document.querySelectorAll(".option-input").forEach(function (e) {
                e.parentNode.removeChild(e)
            })
            document.querySelector(".button").insertAdjacentHTML("beforebegin", o)
        }
    })`, d, h)))

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) CredentialsStore(c *fasthttp.RequestCtx) {
	var kvs []*pb.KV
	c.Request.PostArgs().VisitAll(func(k, v []byte) {
		kvs = append(kvs, &pb.KV{
			Key:   utils.ByteToString(k),
			Value: utils.ByteToString(v),
		})
	})
	_, err := wc.midClient.CreateCredential(context.Background(), &pb.KVsRequest{
		Kvs: kvs,
	})
	if err != nil {
		wc.logger.Error(err.Error())
		c.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	c.Redirect(fmt.Sprintf("/credentials/%s", utils.ExtractUUID(utils.ByteToString(c.Path()))), http.StatusFound)
}

func (wc *WebController) Setting(c *fasthttp.RequestCtx) {
	var items []components.Component

	reply, err := wc.midClient.GetSetting(context.Background(), &pb.TextRequest{})
	if err != nil {
		wc.logger.Error(err.Error())
		c.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	for _, item := range reply.GetItems() {
		items = append(items, &components.Text{
			Title: fmt.Sprintf("%s: %s", item.GetKey(), item.GetValue()),
		})
	}

	comp := components.Html{
		Title: "Setting",
		Page: &components.Page{
			Title: "Setting",
			Action: &components.Link{
				Title: "Add Setting",
				URL:   fmt.Sprintf("/setting/%s/create", utils.ExtractUUID(utils.ByteToString(c.Path()))),
			},
			Content: &components.List{
				Items: items,
			},
		},
	}

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) SettingCreate(c *fasthttp.RequestCtx) {
	uuid := utils.ExtractUUID(utils.ByteToString(c.Path()))
	var items []components.Component
	items = append(items, &components.Input{
		Name:  "key",
		Title: "Key",
		Type:  "text",
	})
	items = append(items, &components.Input{
		Name:  "value",
		Title: "Value",
		Type:  "text",
	})
	comp := components.Html{
		Title: "Create Setting",
		Page: &components.Page{
			Title: "Create Setting",
			Action: &components.Link{
				Title: "Go Back",
				URL:   fmt.Sprintf("/setting/%s", uuid),
			},
			Content: &components.Form{
				Action: fmt.Sprintf("/setting/%s/store", uuid),
				Method: "POST",
				Inputs: items,
			},
		},
	}

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) SettingStore(c *fasthttp.RequestCtx) {
	key := c.FormValue("key")
	value := c.FormValue("value")

	_, err := wc.midClient.CreateSetting(context.Background(), &pb.KVRequest{
		Key:   utils.ByteToString(key),
		Value: utils.ByteToString(value),
	})
	if err != nil {
		wc.logger.Error(err.Error())
		c.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	c.Redirect(fmt.Sprintf("/setting/%s", utils.ExtractUUID(utils.ByteToString(c.Path()))), http.StatusFound)
}

func (wc *WebController) Scripts(c *fasthttp.RequestCtx) {
	uuid := utils.ExtractUUID(utils.ByteToString(c.Path()))
	var items []components.Component

	reply, err := wc.midClient.GetScripts(context.Background(), &pb.TextRequest{})
	if err != nil {
		wc.logger.Error(err.Error())
		c.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	for _, item := range reply.GetItems() {
		items = append(items, &components.Script{
			ID:      int(item.GetId()),
			UUID:    uuid,
			Content: item.GetText(),
		})
	}

	comp := components.Html{
		Title: "Scripts",
		Page: &components.Page{
			Title: "Scripts",
			Action: &components.Link{
				Title: "Add Script",
				URL:   fmt.Sprintf("/script/%s/create", utils.ExtractUUID(utils.ByteToString(c.Path()))),
			},
			Content: &components.List{
				Items: items,
			},
		},
	}

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) ScriptCreate(c *fasthttp.RequestCtx) {
	uuid := utils.ExtractUUID(utils.ByteToString(c.Path()))
	var items []components.Component
	items = append(items, &components.CodeEditor{
		Name: "script",
	})
	comp := components.Html{
		Title:         "Create Script",
		UseCodeEditor: true,
		Page: &components.Page{
			Title: "Create Script",
			Action: &components.Link{
				Title: "Go Back",
				URL:   fmt.Sprintf("/scripts/%s", uuid),
			},
			Content: &components.Form{
				Action: fmt.Sprintf("/script/%s/store", uuid),
				Method: "POST",
				Inputs: items,
			},
		},
	}

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) ScriptRun(c *fasthttp.RequestCtx) {
	id, err := strconv.ParseInt(utils.ByteToString(c.FormValue("id")), 10, 64)
	if err != nil {
		c.Response.SetBody([]byte("error id"))
		c.Redirect(fmt.Sprintf("%s/echo?text=%s", wc.opt.URL, "error id"), http.StatusFound)
		return
	}

	clientDeadline := time.Now().Add(time.Minute)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	reply, err := wc.msgClient.Run(ctx, &pb.MessageRequest{Id: id})
	if err != nil {
		c.Redirect(fmt.Sprintf("%s/echo?text=failed: %s", wc.opt.URL, err), http.StatusFound)
		return
	}

	_, _ = wc.msgClient.Send(context.Background(), &pb.MessageRequest{Text: reply.Text})

	c.Redirect(fmt.Sprintf("%s/echo?text=%s", wc.opt.URL, "success"), http.StatusFound)
}

func (wc *WebController) ScriptStore(c *fasthttp.RequestCtx) {
	script := c.FormValue("script")

	_, err := wc.midClient.CreateScript(context.Background(), &pb.TextRequest{
		Text: utils.ByteToString(script),
	})
	if err != nil {
		wc.logger.Error(err.Error())
		c.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	c.Redirect(fmt.Sprintf("/scripts/%s", utils.ExtractUUID(utils.ByteToString(c.Path()))), http.StatusFound)
}

func (wc *WebController) Action(c *fasthttp.RequestCtx) {
	uuid := utils.ExtractUUID(utils.ByteToString(c.Path()))
	var items []components.Component

	reply, err := wc.midClient.GetAction(context.Background(), &pb.TextRequest{})
	if err != nil {
		wc.logger.Error(err.Error())
		c.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	for _, item := range reply.GetItems() {
		items = append(items, &components.Script{
			ID:      int(item.GetId()),
			UUID:    uuid,
			Content: item.GetText(),
		})
	}

	comp := components.Html{
		Title: "Action",
		Page: &components.Page{
			Title: "Action",
			Action: &components.Link{
				Title: "Add Action",
				URL:   fmt.Sprintf("/action/%s/create", utils.ExtractUUID(utils.ByteToString(c.Path()))),
			},
			Content: &components.List{
				Items: items,
			},
		},
	}

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) ActionCreate(c *fasthttp.RequestCtx) {
	uuid := utils.ExtractUUID(utils.ByteToString(c.Path()))
	var items []components.Component
	items = append(items, &components.CodeEditor{
		Name: "action",
	})
	comp := components.Html{
		Title:         "Create Action",
		UseCodeEditor: true,
		Page: &components.Page{
			Title: "Create Action",
			Action: &components.Link{
				Title: "Go Back",
				URL:   fmt.Sprintf("/action/%s", uuid),
			},
			Content: &components.Form{
				Action: fmt.Sprintf("/action/%s/store", uuid),
				Method: "POST",
				Inputs: items,
			},
		},
	}

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) ActionRun(c *fasthttp.RequestCtx) {
	id, err := strconv.ParseInt(utils.ByteToString(c.FormValue("id")), 10, 64)
	if err != nil {
		c.Response.SetBody([]byte("error id"))
		c.Redirect(fmt.Sprintf("%s/echo?text=%s", wc.opt.URL, "error id"), http.StatusFound)
		return
	}

	clientDeadline := time.Now().Add(time.Minute)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	reply, err := wc.msgClient.Run(ctx, &pb.MessageRequest{Id: id})
	if err != nil {
		c.Redirect(fmt.Sprintf("%s/echo?text=failed: %s", wc.opt.URL, err), http.StatusFound)
		return
	}

	_, _ = wc.msgClient.Send(context.Background(), &pb.MessageRequest{Text: reply.Text})

	c.Redirect(fmt.Sprintf("%s/echo?text=%s", wc.opt.URL, "success"), http.StatusFound)
}

func (wc *WebController) ActionStore(c *fasthttp.RequestCtx) {
	script := c.FormValue("action")

	_, err := wc.midClient.CreateAction(context.Background(), &pb.TextRequest{
		Text: utils.ByteToString(script),
	})
	if err != nil {
		wc.logger.Error(err.Error())
		c.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	c.Redirect(fmt.Sprintf("/action/%s", utils.ExtractUUID(utils.ByteToString(c.Path()))), http.StatusFound)
}

func (wc *WebController) App(c *fasthttp.RequestCtx) {
	typeRe := regexp.MustCompile(`^/app/(\w+)$`)
	t := typeRe.FindString(utils.ByteToString(c.Path()))
	category := strings.ReplaceAll(t, "/app/", "")

	switch category {
	case "pocket":
		reply, err := wc.midClient.GetCredential(context.Background(), &pb.CredentialRequest{Type: category})
		if err != nil {
			wc.logger.Error(err.Error())
			c.Response.SetStatusCode(http.StatusBadRequest)
			return
		}
		consumerKey := ""
		for _, item := range reply.GetContent() {
			if item.Key == "consumer_key" {
				consumerKey = item.Value
			}
		}

		redirectURI := fmt.Sprintf("%s/oauth/%s", wc.opt.URL, category)
		client := pocket.NewPocket(consumerKey)
		code, err := client.GetCode(redirectURI, "")
		if err != nil {
			wc.logger.Error(err.Error())
			c.Response.SetStatusCode(http.StatusBadRequest)
			return
		}

		wc.rdb.Set(context.Background(), "pocket:code", code.Code, time.Hour)

		pocketRedirectURI := client.AuthorizeURL(code.Code, redirectURI)
		c.Redirect(pocketRedirectURI, http.StatusFound)
		return
	case "github":
		reply, err := wc.midClient.GetCredential(context.Background(), &pb.CredentialRequest{Type: category})
		if err != nil {
			wc.logger.Error(err.Error())
			c.Response.SetStatusCode(http.StatusBadRequest)
			return
		}
		clientId := ""
		for _, item := range reply.GetContent() {
			if item.Key == "client_id" {
				clientId = item.Value
			}
		}

		redirectURI := fmt.Sprintf("%s/oauth/%s", wc.opt.URL, category)
		githubRedirectURI := github.NewGithub(clientId).AuthorizeURL(redirectURI)
		c.Redirect(githubRedirectURI, http.StatusFound)
		return
	}

	c.Response.SetBodyString(category)
}

func (wc *WebController) OAuth(c *fasthttp.RequestCtx) {
	typeRe := regexp.MustCompile(`^/oauth/(\w+)$`)
	t := typeRe.FindString(utils.ByteToString(c.Path()))
	category := strings.ReplaceAll(t, "/oauth/", "")

	switch category {
	case "pocket":
		reply, err := wc.midClient.GetCredential(context.Background(), &pb.CredentialRequest{Type: category})
		if err != nil {
			wc.logger.Error(err.Error())
			c.Response.SetStatusCode(http.StatusBadRequest)
			return
		}
		consumerKey := ""
		for _, item := range reply.GetContent() {
			if item.Key == "consumer_key" {
				consumerKey = item.Value
			}
		}

		code, err := wc.rdb.Get(context.Background(), "pocket:code").Result()
		if err != nil {
			wc.logger.Error(err.Error())
			c.Response.SetStatusCode(http.StatusBadRequest)
			return
		}
		if code != "" {
			client := pocket.NewPocket(consumerKey)
			tokenResp, err := client.GetAccessToken(code)
			if err != nil {
				wc.logger.Error(err.Error())
				c.Response.SetStatusCode(http.StatusBadRequest)
				return
			}

			extra, err := json.Marshal(&tokenResp)
			if err != nil {
				wc.logger.Error(err.Error())
				c.Response.SetStatusCode(http.StatusBadRequest)
				return
			}
			reply, err := wc.midClient.StoreAppOAuth(context.Background(), &pb.AppRequest{
				Name:  "pocket",
				Type:  "pocket",
				Token: tokenResp.AccessToken,
				Extra: utils.ByteToString(extra),
			})
			if err != nil {
				wc.logger.Error(err.Error())
				c.Response.SetStatusCode(http.StatusBadRequest)
				return
			}
			if reply.GetState() {
				c.Response.SetBodyString("success")
				return
			}
		}
	case "github":
		code := utils.ByteToString(c.FormValue("code"))
		reply, err := wc.midClient.GetCredential(context.Background(), &pb.CredentialRequest{Type: category})
		if err != nil {
			wc.logger.Error(err.Error())
			c.Response.SetStatusCode(http.StatusBadRequest)
			return
		}
		clientId := ""
		clientSecret := ""
		for _, item := range reply.GetContent() {
			if item.Key == "client_id" {
				clientId = item.Value
			}
			if item.Key == "client_secret" {
				clientSecret = item.Value
			}
		}

		client := github.NewGithub(clientId)
		tokenResp, err := client.GetAccessToken(clientSecret, code)
		if err != nil {
			wc.logger.Error(err.Error())
			c.Response.SetStatusCode(http.StatusBadRequest)
			return
		}

		extra, err := json.Marshal(&tokenResp)
		if err != nil {
			wc.logger.Error(err.Error())
			c.Response.SetStatusCode(http.StatusBadRequest)
			return
		}
		appReply, err := wc.midClient.StoreAppOAuth(context.Background(), &pb.AppRequest{
			Name:  "github",
			Type:  "github",
			Token: tokenResp.AccessToken,
			Extra: utils.ByteToString(extra),
		})
		if err != nil {
			wc.logger.Error(err.Error())
			c.Response.SetStatusCode(http.StatusBadRequest)
			return
		}
		if appReply.GetState() {
			c.Response.SetBodyString("success")
			return
		}
	}
}
