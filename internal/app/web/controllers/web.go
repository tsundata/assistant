package controllers

import (
	"github.com/tsundata/assistant/internal/app/web"
	"github.com/tsundata/assistant/internal/app/web/components"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type WebController struct {
	o         *web.Options
	logger    *zap.Logger
	subClient *rpc.Client
	msgClient *rpc.Client
}

func NewWebController(o *web.Options, logger *zap.Logger, subClient *rpc.Client, msgClient *rpc.Client) *WebController {
	return &WebController{o: o, logger: logger, subClient: subClient, msgClient: msgClient}
}

func (wc *WebController) Index(c *fasthttp.RequestCtx) {
	comp := components.Html{
		Title: "Title...",
		Page: &components.Page{
			Title: "Title...",
			Content: &components.List{
				Items: []components.Component{
					&components.Link{
						Name:  "Link1...",
						Title: "Link1...",
						URL:   "https://www.demo.com/t/1",
					},
					&components.Link{
						Name:  "Link2...",
						Title: "Link2...",
						URL:   "https://www.demo.com/t/2",
					},
				},
			},
		},
	}

	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response.SetBody([]byte(comp.GetContent()))
}

func (wc *WebController) Robots(c *fasthttp.RequestCtx) {
	txt := `User-agent: *
Disallow: /`

	c.Response.SetBody([]byte(txt))
}
