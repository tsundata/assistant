package sdk

import (
	"github.com/go-resty/resty/v2"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/config"
	"time"
)

type GatewayClient struct {
	r *resty.Client
}

func NewGatewayClient(c *config.AppConfig) *GatewayClient {
	client := resty.New()
	client.SetHostURL(c.Gateway.Url)
	client.SetTimeout(time.Minute)
	return &GatewayClient{r: client}
}

func (c *GatewayClient) AuthToken(token string) *GatewayClient {
	c.r.SetAuthToken(token)
	return c
}

func (c *GatewayClient) GetPage(uuid string) (result *pb.PageReply, err error) {
	resp, err := c.r.R().
		SetQueryParam("uuid", uuid).
		SetResult(result).
		Get("page")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.PageReply), nil
}
