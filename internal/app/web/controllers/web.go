package controllers

import (
	"context"
	"encoding/json"
	"github.com/tsundata/assistant/internal/app/web"
	"github.com/tsundata/assistant/internal/app/web/components"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net/http"
	"regexp"
)

type WebController struct {
	o          *web.Options
	logger     *zap.Logger
	pageClient *rpc.Client
}

func NewWebController(o *web.Options, logger *zap.Logger, pageClient *rpc.Client) *WebController {
	return &WebController{o: o, logger: logger, pageClient: pageClient}
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
	err := wc.pageClient.Call(context.Background(), "Get", &payload, &reply)
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
