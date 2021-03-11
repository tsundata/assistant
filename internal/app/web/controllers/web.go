package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/web"
	"github.com/tsundata/assistant/internal/app/web/components"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
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
	logger    *logger.Logger
	midClient pb.MiddleClient
	msgClient pb.MessageClient
	wfClient  pb.WorkflowClient
}

func NewWebController(opt *web.Options, rdb *redis.Client, logger *logger.Logger,
	midClient pb.MiddleClient, msgClient pb.MessageClient, wfClient pb.WorkflowClient) *WebController {
	return &WebController{opt: opt, rdb: rdb, logger: logger, midClient: midClient, msgClient: msgClient, wfClient: wfClient}
}

func (wc *WebController) Index(c *fiber.Ctx) error {
	return c.SendString("Web")
}

func (wc *WebController) Echo(c *fiber.Ctx) error {
	return c.SendString(c.FormValue("text"))
}

func (wc *WebController) Robots(c *fiber.Ctx) error {
	txt := `User-agent: *
Disallow: /`

	return c.SendString(txt)
}

func (wc *WebController) Page(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	reply, err := wc.midClient.GetPage(context.Background(), &pb.PageRequest{
		Uuid: uuid,
	})
	if err != nil {
		wc.logger.Error(err)
		return c.SendStatus(http.StatusBadRequest)
	}
	if reply.GetContent() == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	var list []string
	err = json.Unmarshal([]byte(reply.GetContent()), &list)
	if err != nil {
		wc.logger.Error(err)
		return c.SendStatus(http.StatusBadRequest)
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

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) Qr(c *fiber.Ctx) error {
	text := c.Params("text", "")
	if text == "" {
		return c.SendStatus(http.StatusNotFound)
	}

	txt, err := url.QueryUnescape(text)
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusNotFound).SendString("error text")
	}

	png, err := qrcode.Encode(txt, qrcode.Medium, 512)
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusNotFound).SendString("error qr")
	}

	c.Response().Header.Set("Content-Type", "image/png")
	return c.Send(png)
}

func (wc *WebController) Apps(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	var items []components.Component

	reply, err := wc.midClient.GetApps(context.Background(), &pb.TextRequest{})
	if err != nil {
		wc.logger.Error(err)
		return c.SendStatus(http.StatusBadRequest)
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

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) Memo(c *fiber.Ctx) error {
	var items []components.Component

	reply, err := wc.msgClient.List(context.Background(), &pb.MessageRequest{})
	if err != nil {
		wc.logger.Error(err)
		return c.SendStatus(http.StatusBadRequest)
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
			wc.logger.Error(err)
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

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) Credentials(c *fiber.Ctx) error {
	var items []components.Component

	reply, err := wc.midClient.GetMaskingCredentials(context.Background(), &pb.TextRequest{})
	if err != nil {
		wc.logger.Error(err)
		return c.SendStatus(http.StatusBadRequest)
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
				URL:   fmt.Sprintf("/credentials/%s/create", c.Params("uuid")),
			},
			Content: &components.List{
				Items: items,
			},
		},
	}

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) CredentialsCreate(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	options := vendors.OAuthProvidersOptions

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

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) CredentialsStore(c *fiber.Ctx) error {
	var kvs []*pb.KV
	c.Request().PostArgs().VisitAll(func(k, v []byte) {
		kvs = append(kvs, &pb.KV{
			Key:   utils.ByteToString(k),
			Value: utils.ByteToString(v),
		})
	})
	_, err := wc.midClient.CreateCredential(context.Background(), &pb.KVsRequest{
		Kvs: kvs,
	})
	if err != nil {
		wc.logger.Error(err)
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.Redirect(fmt.Sprintf("/credentials/%s", c.Params("uuid")), http.StatusFound)
}

func (wc *WebController) Setting(c *fiber.Ctx) error {
	var items []components.Component

	reply, err := wc.midClient.GetSetting(context.Background(), &pb.TextRequest{})
	if err != nil {
		wc.logger.Error(err)
		return c.SendStatus(http.StatusBadRequest)
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
				URL:   fmt.Sprintf("/setting/%s/create", c.Params("uuid")),
			},
			Content: &components.List{
				Items: items,
			},
		},
	}

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) SettingCreate(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
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

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) SettingStore(c *fiber.Ctx) error {
	key := c.FormValue("key")
	value := c.FormValue("value")

	_, err := wc.midClient.CreateSetting(context.Background(), &pb.KVRequest{
		Key:   key,
		Value: value,
	})
	if err != nil {
		wc.logger.Error(err)
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.Redirect(fmt.Sprintf("/setting/%s", c.Params("uuid")), http.StatusFound)
}

func (wc *WebController) Scripts(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	var items []components.Component

	reply, err := wc.msgClient.GetScriptMessages(context.Background(), &pb.TextRequest{})
	if err != nil {
		wc.logger.Error(err)
		return c.SendStatus(http.StatusBadRequest)
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
				URL:   fmt.Sprintf("/script/%s/create", c.Params("uuid")),
			},
			Content: &components.List{
				Items: items,
			},
		},
	}

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) ScriptCreate(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
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

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) ScriptRun(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if err != nil {
		return c.Redirect(fmt.Sprintf("%s/echo?text=%s", wc.opt.URL, "error id"), http.StatusFound)
	}

	clientDeadline := time.Now().Add(time.Minute)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	reply, err := wc.msgClient.Run(ctx, &pb.MessageRequest{Id: id})
	if err != nil {
		return c.Redirect(fmt.Sprintf("%s/echo?text=failed: %s", wc.opt.URL, err), http.StatusFound)
	}

	_, _ = wc.msgClient.Send(context.Background(), &pb.MessageRequest{Text: reply.Text})

	return c.Redirect(fmt.Sprintf("%s/echo?text=%s", wc.opt.URL, "success"), http.StatusFound)
}

func (wc *WebController) ScriptStore(c *fiber.Ctx) error {
	script := c.FormValue("script")

	_, err := wc.msgClient.CreateScriptMessage(context.Background(), &pb.TextRequest{
		Text: script,
	})
	if err != nil {
		wc.logger.Error(err)
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.Redirect(fmt.Sprintf("/scripts/%s", c.Params("uuid")), http.StatusFound)
}

func (wc *WebController) Action(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	var items []components.Component

	reply, err := wc.msgClient.GetActionMessages(context.Background(), &pb.TextRequest{})
	if err != nil {
		wc.logger.Error(err)
		return c.SendStatus(http.StatusBadRequest)
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
				URL:   fmt.Sprintf("/action/%s/create", c.Params("uuid")),
			},
			Content: &components.List{
				Items: items,
			},
		},
	}

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) ActionCreate(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
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

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) ActionRun(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if err != nil {
		return c.Redirect(fmt.Sprintf("%s/echo?text=%s", wc.opt.URL, "error id"), http.StatusFound)
	}

	clientDeadline := time.Now().Add(time.Minute)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	reply, err := wc.msgClient.Run(ctx, &pb.MessageRequest{Id: id})
	if err != nil {
		return c.Redirect(fmt.Sprintf("%s/echo?text=failed: %s", wc.opt.URL, err), http.StatusFound)
	}

	_, _ = wc.msgClient.Send(context.Background(), &pb.MessageRequest{Text: reply.Text})

	return c.Redirect(fmt.Sprintf("%s/echo?text=%s", wc.opt.URL, "success"), http.StatusFound)
}

func (wc *WebController) ActionStore(c *fiber.Ctx) error {
	action := c.FormValue("action")

	_, err := wc.msgClient.CreateActionMessage(context.Background(), &pb.TextRequest{
		Text: action,
	})
	if err != nil {
		wc.logger.Error(err)
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.Redirect(fmt.Sprintf("/action/%s", c.Params("uuid")), http.StatusFound)
}

func (wc *WebController) App(c *fiber.Ctx) error {
	provider := vendors.NewOAuthProvider(wc.rdb, c, wc.opt.URL)
	return provider.Redirect(c, wc.midClient)
}

func (wc *WebController) OAuth(c *fiber.Ctx) error {
	provider := vendors.NewOAuthProvider(wc.rdb, c, wc.opt.URL)
	err := provider.StoreAccessToken(c, wc.midClient)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	return c.SendString("Success")
}
