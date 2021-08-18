package sdk

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/config"
	"net/http"
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
		Get("page")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) GetApps() (result *pb.AppsReply, err error) {
	resp, err := c.r.R().
		Get("apps")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) GetMessages() (result *pb.MessagesReply, err error) {
	resp, err := c.r.R().
		Get("messages")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) GetMaskingCredentials() (result *pb.MaskingReply, err error) {
	resp, err := c.r.R().
		Get("masking_credentials")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) CreateCredential(in *pb.KVsRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		Post("credential")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) GetSettings() (result *pb.SettingsReply, err error) {
	resp, err := c.r.R().
		Get("settings")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) CreateSetting(in *pb.KVRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		Post("setting")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) GetActionMessages() (result *pb.ActionReply, err error) {
	resp, err := c.r.R().
		Get("action/messages")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) CreateActionMessage(in *pb.TextRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		Post("action/message")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) DeleteWorkflowMessage(in *pb.Message) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		Delete("workflow/message")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) RunMessage(in *pb.MessageRequest) (result *pb.TextReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		Post("message/run")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) SendMessage(in *pb.MessageRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		Post("message/send")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) WebhookTrigger(in *pb.TriggerRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		Post("webhook/trigger")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) GetCredential(id string) (result *pb.CredentialReply, err error) {
	resp, err := c.r.R().
		SetQueryParam("type", id).
		Get("credential")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) GetRoleImage() (result *pb.BytesReply, err error) {
	resp, err := c.r.R().
		Get("role/image")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) GetChartData(uuid string) (result *pb.ChartDataReply, err error) {
	resp, err := c.r.R().
		SetQueryParam("uuid", uuid).
		Get("chart")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) StoreAppOAuth(in *pb.AppRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		Post("app/oauth")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

func (c *GatewayClient) Authorization(in *pb.TextRequest) (result *pb.StateReply, err error) {
	resp, err := c.r.R().
		SetBody(in).
		Post("auth")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.String())
	}

	err = json.Unmarshal(resp.Body(), &result)
	return
}

var ProviderSet = wire.NewSet(NewGatewayClient)
