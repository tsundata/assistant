package controllers

import (
	"context"
	"encoding/json"
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
}

func NewWebController(opt *web.Options, logger *zap.Logger, midClient pb.MiddleClient) *WebController {
	return &WebController{opt: opt, logger: logger, midClient: midClient}
}

func (wc *WebController) Index(c *fasthttp.RequestCtx) {
	c.Response.SetBody([]byte("Web"))
}

func (wc *WebController) Robots(c *fasthttp.RequestCtx) {
	txt := `User-agent: *
Disallow: /`

	c.Response.SetBody([]byte(txt))
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
	if err != nil || reply.GetContent() == "" {
		c.Response.SetStatusCode(http.StatusNotFound)
		return
	}

	var list []string
	err = json.Unmarshal([]byte(reply.GetContent()), &list)
	if err != nil {
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
		c.Response.SetBodyString("error text")
		return
	}

	png, err := qrcode.Encode(txt, qrcode.Medium, 512)
	if err != nil {
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

func (wc *WebController) Credentials(c *fasthttp.RequestCtx) {
	var items []components.Component
	items = append(items, &components.Link{
		Title: "pocket (Pocket)",
		Name:  "pocket (Pocket)",
		URL:   "#1",
	})
	items = append(items, &components.Link{
		Title: "my's github (Github)",
		Name:  "my's github (Github)",
		URL:   "#2",
	})
	items = append(items, &components.Link{
		Title: "my's facebook (Facebook)",
		Name:  "my's facebook (Facebook)",
		URL:   "#3",
	})
	comp := components.Html{
		Title:   "Credentials",
		UseIcon: true,
		Page: &components.Page{
			Title: "Credentials",
			Action: &components.Button{
				Title: "Add Credentials",
				URL:   "/credentials/xxx/create",
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
		Title: "Name",
		Type:  "text",
	})
	items = append(items, &components.Input{
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
				URL:   "/credentials/xxx",
			},
			Content: &components.Form{
				Action: "/demo",
				Method: "POST",
				Inputs: items,
			},
		},
	}

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) CredentialsStore(c *fasthttp.RequestCtx) {

}

func (wc *WebController) Setting(c *fasthttp.RequestCtx) {
	var items []components.Component
	items = append(items, &components.Text{
		Title: "foo: bar",
	})
	items = append(items, &components.Text{
		Title: "open: false",
	})
	items = append(items, &components.Text{
		Title: "my's facebook (Facebook): true",
	})
	comp := components.Html{
		Title:   "Setting",
		UseIcon: true,
		Page: &components.Page{
			Title: "Setting",
			Action: &components.Button{
				Title: "Add Setting",
				URL:   "/setting/xxx/create",
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
	var items []components.Component
	items = append(items, &components.Input{
		Title: "Key",
		Type:  "text",
	})
	items = append(items, &components.Input{
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
				URL:   "/setting/xxx",
			},
			Content: &components.Form{
				Action: "/demo",
				Method: "POST",
				Inputs: items,
			},
		},
	}

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) SettingStore(c *fasthttp.RequestCtx) {

}

func checkFlag(path string) bool {
	return true
}
