package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/web"
	"github.com/tsundata/assistant/internal/app/web/components"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"regexp"
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
		items = append(items, &components.Text{
			Name:  item,
			Title: item,
		})
	}

	comp := components.Html{
		Title: reply.Title,
		Page: &components.Page{
			Title: reply.Title,
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
	items = append(items, &components.App{
		Name: "Pocket",
		Icon: "get-pocket",
		Text: "Authorized",
	})
	items = append(items, &components.App{
		Name: "Github",
		Icon: "github",
		Text: "Unauthorized",
	})
	items = append(items, &components.App{
		Name: "Facebook",
		Icon: "facebook",
		Text: "Unauthorized",
	})
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

	for _, item := range reply.Messages {
		items = append(items, &components.Memo{
			Time: item.GetTime(),
			Content: &components.Text{
				Title: item.GetText(),
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

	for _, item := range reply.Items {
		items = append(items, &components.Link{
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
			Action: &components.Button{
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
			Action: &components.Button{
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

	for _, item := range reply.Items {
		items = append(items, &components.Text{
			Title: fmt.Sprintf("%s: %s", item.Key, item.Value),
		})
	}

	comp := components.Html{
		Title:   "Setting",
		UseIcon: true,
		Page: &components.Page{
			Title: "Setting",
			Action: &components.Button{
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
			Action: &components.Button{
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
