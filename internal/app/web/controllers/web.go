package controllers

import (
	"context"
	"encoding/json"
	"github.com/skip2/go-qrcode"
	"github.com/tsundata/assistant/internal/app/web"
	"github.com/tsundata/assistant/internal/app/web/components"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"regexp"
)

type WebController struct {
	o         *web.Options
	logger    *zap.Logger
	webClient *rpc.Client
}

func NewWebController(o *web.Options, logger *zap.Logger, webClient *rpc.Client) *WebController {
	return &WebController{o: o, logger: logger, webClient: webClient}
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

	payload := model.Page{
		UUID: string(r[0]),
	}
	var reply model.Page
	err := wc.webClient.Call(context.Background(), "GetPage", &payload, &reply)
	if err != nil {
		c.Response.SetStatusCode(http.StatusNotFound)
		return
	}

	var list []string
	err = json.Unmarshal([]byte(reply.Content), &list)
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

	txt, err := url.QueryUnescape(string(r[1]))
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
