package controllers

import (
	"net/http"

	"github.com/tsundata/framework"
)

type GatewayController struct {
}

func NewGatewayController() *GatewayController {
	return &GatewayController{}
}

func (gc *GatewayController) Index(c *framework.Context) {
	c.JSON(http.StatusOK, framework.H{"path": "ROOT"})
}
