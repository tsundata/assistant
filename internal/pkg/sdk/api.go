package sdk

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
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
	c.r.SetAuthScheme("X-Token")
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

func (c *GatewayClient) GetApps() (result *pb.AppsReply, err error) {
	resp, err := c.r.R().
		SetResult(result).
		Get("apps")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.AppsReply), nil
}

func (c *GatewayClient) GetMessages() (result *pb.MessageListReply, err error) {
	resp, err := c.r.R().
		SetResult(result).
		Get("messages")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.MessageListReply), nil
}

func (c *GatewayClient) GetMaskingCredentials() (result *pb.MaskingReply, err error) {
	resp, err := c.r.R().
		SetResult(result).
		Get("masking_credentials")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.MaskingReply), nil
}

func (c *GatewayClient) CreateCredential(in *pb.KVsRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		SetResult(result).
		Post("credential")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.StateReply), nil
}

func (c *GatewayClient) GetSettings() (result *pb.SettingsReply, err error) {
	resp, err := c.r.R().
		SetResult(result).
		Get("settings")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.SettingsReply), nil
}

func (c *GatewayClient) CreateSetting(in *pb.KVRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		SetResult(result).
		Post("setting")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.StateReply), nil
}

func (c *GatewayClient) GetActionMessages() (result *pb.ActionReply, err error) {
	resp, err := c.r.R().
		SetResult(result).
		Get("action/messages")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.ActionReply), nil
}

func (c *GatewayClient) CreateActionMessage(in *pb.TextRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		SetResult(result).
		Post("action/message")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.StateReply), nil
}

func (c *GatewayClient) DeleteWorkflowMessage(in *pb.MessageRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		SetResult(result).
		Delete("workflow/message")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.StateReply), nil
}

func (c *GatewayClient) RunMessage(in *pb.MessageRequest) (result *pb.TextReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		SetResult(result).
		Post("message/run")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.TextReply), nil
}

func (c *GatewayClient) SendMessage(in *pb.MessageRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		SetResult(result).
		Post("message/send")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.StateReply), nil
}

func (c *GatewayClient) WebhookTrigger(in *pb.TriggerRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		SetResult(result).
		Post("webhook/trigger")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.StateReply), nil
}

func (c *GatewayClient) GetCredential(id string) (result *pb.CredentialReply, err error) {
	resp, err := c.r.R().
		SetQueryParam("type", id).
		SetResult(result).
		Post("credential")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.CredentialReply), nil
}

func (c *GatewayClient) StoreAppOAuth(in *pb.AppRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		SetResult(result).
		Post("app/oauth")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.StateReply), nil
}

func (c *GatewayClient) Authorization(in *pb.TextRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		SetResult(result).
		Post("app/oauth")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.RawBody().Close() }()
	return resp.Result().(*pb.StateReply), nil
}

var ProviderSet = wire.NewSet(NewGatewayClient)
