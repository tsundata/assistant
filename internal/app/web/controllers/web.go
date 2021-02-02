package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/web"
	"github.com/tsundata/assistant/internal/app/web/components"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/valyala/fasthttp"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type WebController struct {
	opt       *web.Options
	logger    *zap.Logger
	midClient pb.MiddleClient
	msgClient pb.MessageClient
}

func NewWebController(opt *web.Options, logger *zap.Logger,
	midClient pb.MiddleClient, msgClient pb.MessageClient) *WebController {
	return &WebController{opt: opt, logger: logger, midClient: midClient, msgClient: msgClient}
}

func (wc *WebController) Index(c *fasthttp.RequestCtx) {
	c.Response.SetBody([]byte("Web"))
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
	var items []components.Component

	reply, err := wc.midClient.Apps(context.Background(), &pb.TextRequest{})
	if err != nil {
		wc.logger.Error(err.Error())
		c.Response.SetStatusCode(http.StatusNotFound)
		return
	}

	for _, app := range reply.GetApps() {
		authStr := "Unauthorized"
		if app.GetIsAuthorized() {
			authStr = "Authorized"
		}
		items = append(items, &components.App{
			Name: app.GetTitle(),
			Icon: "rocket",
			Text: authStr,
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
		c.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	for _, item := range reply.GetItems() {
		items = append(items, &components.LinkButton{
			Title: item.Key,
			Name:  item.Value,
			URL:   "javascript:void(0)",
		})
	}

	comp := components.Html{
		Title:   "Credentials",
		UseIcon: true,
		Page: &components.Page{
			Title: "Credentials",
			Action: &components.Link{
				Title: "Add Credentials",
				URL:   fmt.Sprintf("/credentials/%s/create", extractUUID(c.Path())),
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
	var items []components.Component
	items = append(items, &components.Input{
		Name:  "name",
		Title: "Name",
		Type:  "text",
	})
	items = append(items, &components.Input{
		Name:  "k1",
		Title: "App Key",
		Type:  "text",
	})
	comp := components.Html{
		Title:   "Create Credentials",
		UseIcon: true,
		Page: &components.Page{
			Title: "Create Credentials",
			Action: &components.Link{
				Title: "Go Back",
				URL:   fmt.Sprintf("/credentials/%s", extractUUID(c.Path())),
			},
			Content: &components.Form{
				Action: fmt.Sprintf("/credentials/%s/store", extractUUID(c.Path())),
				Method: "POST",
				Inputs: items,
			},
		},
	}

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
		c.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	c.Redirect(fmt.Sprintf("/credentials/%s", extractUUID(c.Path())), http.StatusFound)
}

func (wc *WebController) Setting(c *fasthttp.RequestCtx) {
	var items []components.Component

	reply, err := wc.midClient.GetSetting(context.Background(), &pb.TextRequest{})
	if err != nil {
		c.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	for _, item := range reply.GetItems() {
		items = append(items, &components.Text{
			Title: fmt.Sprintf("%s: %s", item.Key, item.Value),
		})
	}

	comp := components.Html{
		Title:   "Setting",
		UseIcon: true,
		Page: &components.Page{
			Title: "Setting",
			Action: &components.Link{
				Title: "Add Setting",
				URL:   fmt.Sprintf("/setting/%s/create", extractUUID(c.Path())),
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
	uuid := extractUUID(c.Path())
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
		Title:   "Create Setting",
		UseIcon: true,
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
		c.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	c.Redirect(fmt.Sprintf("/setting/%s", extractUUID(c.Path())), http.StatusFound)
}
