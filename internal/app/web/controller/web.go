package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/web/component"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/sdk"
	"github.com/tsundata/assistant/internal/pkg/util"
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
)

type WebController struct {
	opt     *config.AppConfig
	rdb     *redis.Client
	logger  *logger.Logger
	gateway *sdk.GatewayClient
}

func NewWebController(opt *config.AppConfig, rdb *redis.Client, logger *logger.Logger, gateway *sdk.GatewayClient) *WebController {
	return &WebController{opt: opt, rdb: rdb, logger: logger, gateway: gateway}
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

	reply, err := wc.gateway.GetPage(uuid)
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	if reply.GetContent() == "" {
		return c.Status(http.StatusBadRequest).SendString("content empty")
	}

	if reply.GetType() == "html" {
		c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
		return c.SendString(reply.GetContent())
	}

	if reply.GetType() != "json" {
		return c.Status(http.StatusBadRequest).SendString("error type")
	}

	var list []string
	err = json.Unmarshal([]byte(reply.GetContent()), &list)
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	var items []component.Component

	for _, item := range list {
		re, _ := regexp.Compile(util.UrlRegex)
		s := re.FindString(item)
		if s != "" {
			item = strings.ReplaceAll(item, s, fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, s, s))
		}
		item = strings.ReplaceAll(item, "\n", "<br>")

		items = append(items, &component.Text{
			Title: item,
		})
	}

	comp := component.Html{
		Title: reply.GetTitle(),
		Page: &component.Page{
			Title: reply.GetTitle(),
			Content: &component.List{
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
	var items []component.Component

	reply, err := wc.gateway.GetApps()
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	for _, app := range reply.GetApps() {
		authStr := "Unauthorized"
		authorizedURL := fmt.Sprintf("/app/%s?uuid=%s", app.GetType(), uuid)
		if app.GetIsAuthorized() {
			authStr = "Authorized"
			authorizedURL = "javascript:void(0);"
		}
		items = append(items, &component.App{
			Name: app.GetTitle(),
			Icon: "rocket",
			Text: fmt.Sprintf("%s (%s)", app.GetType(), authStr),
			URL:  authorizedURL,
		})
	}

	comp := component.Html{
		Title:   "Apps",
		UseIcon: true,
		Page: &component.Page{
			Title: "Apps",
			Content: &component.List{
				Items: items,
			},
		},
	}

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) Memo(c *fiber.Ctx) error {
	var items []component.Component

	reply, err := wc.gateway.GetMessages()
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
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
		err := md.Convert(util.StringToByte(item.GetText()), &buf)
		if err != nil {
			wc.logger.Error(err)
		} else {
			text = buf.String()
		}

		items = append(items, &component.Memo{
			Time: item.GetTime(),
			Content: &component.Text{
				Title: text,
			},
		})
	}

	comp := component.Html{
		Title: "Memo",
		Page: &component.Page{
			Title: "Memo",
			Content: &component.List{
				Items: items,
			},
		},
	}

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) Credentials(c *fiber.Ctx) error {
	var items []component.Component

	reply, err := wc.gateway.GetMaskingCredentials()
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	for _, item := range reply.GetItems() {
		items = append(items, &component.LinkButton{
			Title: item.GetKey(),
			Name:  item.GetValue(),
			URL:   "javascript:void(0)",
		})
	}

	comp := component.Html{
		Title: "Credentials",
		Page: &component.Page{
			Title: "Credentials",
			Action: &component.Link{
				Title: "Add Credentials",
				URL:   fmt.Sprintf("/credentials/%s/create", c.Params("uuid")),
			},
			Content: &component.List{
				Items: items,
			},
		},
	}

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) CredentialsCreate(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	options := vendors.ProviderCredentialOptions

	selectOption := make(map[string]string)
	selectOption[""] = "-"
	for k := range options {
		selectOption[k] = strings.Title(k)
	}

	var items []component.Component
	items = append(items, &component.Input{
		Name:  "name",
		Title: "Name",
		Type:  "text",
	})
	items = append(items, &component.Select{
		Name:  "type",
		Title: "Type",
		Value: selectOption,
	})
	comp := component.Html{
		Title: "Create Credentials",
		Page: &component.Page{
			Title: "Create Credentials",
			Action: &component.Link{
				Title: "Go Back",
				URL:   fmt.Sprintf("/credentials/%s", uuid),
			},
			Content: &component.Form{
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
			Key:   util.ByteToString(k),
			Value: util.ByteToString(v),
		})
	})
	_, err := wc.gateway.CreateCredential(&pb.KVsRequest{
		Kvs: kvs,
	})
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Redirect(fmt.Sprintf("/credentials/%s", c.Params("uuid")), http.StatusFound)
}

func (wc *WebController) Setting(c *fiber.Ctx) error {
	var items []component.Component

	reply, err := wc.gateway.GetSettings()
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	for _, item := range reply.GetItems() {
		items = append(items, &component.Text{
			Title: fmt.Sprintf("%s: %s", item.GetKey(), item.GetValue()),
		})
	}

	comp := component.Html{
		Title: "Setting",
		Page: &component.Page{
			Title: "Setting",
			Action: &component.Link{
				Title: "Add Setting",
				URL:   fmt.Sprintf("/setting/%s/create", c.Params("uuid")),
			},
			Content: &component.List{
				Items: items,
			},
		},
	}

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) SettingCreate(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	var items []component.Component
	items = append(items, &component.Input{
		Name:  "key",
		Title: "Key",
		Type:  "text",
	})
	items = append(items, &component.Input{
		Name:  "value",
		Title: "Value",
		Type:  "text",
	})
	comp := component.Html{
		Title: "Create Setting",
		Page: &component.Page{
			Title: "Create Setting",
			Action: &component.Link{
				Title: "Go Back",
				URL:   fmt.Sprintf("/setting/%s", uuid),
			},
			Content: &component.Form{
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

	_, err := wc.gateway.CreateSetting(&pb.KVRequest{
		Key:   key,
		Value: value,
	})
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Redirect(fmt.Sprintf("/setting/%s", c.Params("uuid")), http.StatusFound)
}

func (wc *WebController) Action(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	var items []component.Component

	reply, err := wc.gateway.GetActionMessages()
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	for _, item := range reply.GetItems() {
		items = append(items, &component.Action{
			ID:      int(item.GetId()),
			UUID:    uuid,
			Content: item.GetText(),
		})
	}

	comp := component.Html{
		Title: "Action",
		Page: &component.Page{
			Title: "Action",
			Action: &component.Link{
				Title: "Add Action",
				URL:   fmt.Sprintf("/action/%s/create", c.Params("uuid")),
			},
			Content: &component.List{
				Items: items,
			},
		},
	}

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(string(comp.GetContent()))
}

func (wc *WebController) ActionCreate(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	var items []component.Component
	items = append(items, &component.CodeEditor{
		Name: "action",
	})
	comp := component.Html{
		Title:         "Create Action",
		UseCodeEditor: true,
		Page: &component.Page{
			Title: "Create Action",
			Action: &component.Link{
				Title: "Go Back",
				URL:   fmt.Sprintf("/action/%s", uuid),
			},
			Content: &component.Form{
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
		return c.Redirect(fmt.Sprintf("%s/echo?text=%s", wc.opt.Web.Url, "error id"), http.StatusFound)
	}

	reply, err := wc.gateway.RunMessage(&pb.MessageRequest{Id: id})
	if err != nil {
		return c.Redirect(fmt.Sprintf("%s/echo?text=failed: %s", wc.opt.Web.Url, err), http.StatusFound)
	}

	_, _ = wc.gateway.SendMessage(&pb.MessageRequest{Text: reply.GetText()})

	return c.Redirect(fmt.Sprintf("%s/echo?text=%s", wc.opt.Web.Url, "ok"), http.StatusFound)
}

func (wc *WebController) ActionStore(c *fiber.Ctx) error {
	action := c.FormValue("action")

	_, err := wc.gateway.CreateActionMessage(&pb.TextRequest{
		Text: action,
	})
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Redirect(fmt.Sprintf("/action/%s", c.Params("uuid")), http.StatusFound)
}

func (wc *WebController) WorkflowDelete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if err != nil {
		return c.Redirect(fmt.Sprintf("%s/echo?text=%s", wc.opt.Web.Url, "error id"), http.StatusFound)
	}

	_, err = wc.gateway.DeleteWorkflowMessage(&pb.MessageRequest{Id: id})
	if err != nil {
		return c.Redirect(fmt.Sprintf("%s/echo?text=failed: %s", wc.opt.Web.Url, err), http.StatusFound)
	}

	return c.Redirect(fmt.Sprintf("%s/echo?text=%s", wc.opt.Web.Url, "ok"), http.StatusFound)
}

func (wc *WebController) Role(c *fiber.Ctx) error {
	reply, err := wc.gateway.GetRoleImage()
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	c.Response().Header.Set("Content-Type", "image/png")
	return c.SendString(reply.GetText())
}

func (wc *WebController) App(c *fiber.Ctx) error {
	category := c.Params("category")
	provider := vendors.NewOAuthProvider(wc.rdb, category, wc.opt.Web.Url)
	return provider.Redirect(c, wc.gateway)
}

func (wc *WebController) OAuth(c *fiber.Ctx) error {
	category := c.Params("category")
	provider := vendors.NewOAuthProvider(wc.rdb, category, wc.opt.Web.Url)
	err := provider.StoreAccessToken(c, wc.gateway)
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	return c.SendString("ok")
}

func (wc *WebController) Webhook(c *fiber.Ctx) error {
	flag := c.Params("flag", "")

	// Headers(Authorization: Base ?) -> query(secret)
	secret := c.Get("Authorization", "")
	secret = strings.ReplaceAll(secret, "Base ", "")
	if secret == "" {
		secret = c.Query("secret", "")
	}

	_, err := wc.gateway.WebhookTrigger(&pb.TriggerRequest{
		Type:   "webhook",
		Flag:   flag,
		Secret: secret,
		Header: c.Request().Header.String(),
		Body:   util.ByteToString(c.Request().Body()),
	})
	if err != nil {
		wc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.SendString("ok")
}
