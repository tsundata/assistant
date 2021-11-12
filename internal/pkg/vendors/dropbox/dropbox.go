package dropbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/sdk"
	"github.com/tsundata/assistant/internal/pkg/util"
	"io"
	"net/http"
	"time"
)

const (
	ID              = "dropbox"
	ClientIdKey     = "key"
	ClientSecretKey = "secret"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	UID         string `json:"uid"`
	AccountID   string `json:"account_id"`
	Scope       string `json:"scope"`
}

type Dropbox struct {
	c            *resty.Client
	clientId     string
	clientSecret string
	redirectURI  string
	accessToken  string
}

func NewDropbox(clientId, clientSecret, redirectURI, accessToken string) *Dropbox {
	v := &Dropbox{clientId: clientId, clientSecret: clientSecret, redirectURI: redirectURI, accessToken: accessToken}

	v.c = resty.New()
	v.c.SetBaseURL("https://api.dropboxapi.com")
	v.c.SetTimeout(time.Minute)

	return v
}

func (v *Dropbox) AuthorizeURL() string {
	return fmt.Sprintf("https://www.dropbox.com/oauth2/authorize?client_id=%s&response_type=code&redirect_uri=%s", v.clientId, v.redirectURI)
}

func (v *Dropbox) GetAccessToken(code string) (interface{}, error) {
	resp, err := v.c.R().
		SetBasicAuth(v.clientId, v.clientSecret).
		SetFormData(map[string]string{
			"code":         code,
			"grant_type":   "authorization_code",
			"redirect_uri": v.redirectURI,
		}).
		Post("/oauth2/token")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		var result TokenResponse
		err = json.Unmarshal(resp.Body(), &result)
		if err != nil {
			return nil, err
		}
		v.accessToken = result.AccessToken
		return &result, nil
	} else {
		return nil, fmt.Errorf("%d, %s", resp.StatusCode(), util.ByteToString(resp.Body()))
	}
}

func (v *Dropbox) Redirect(c *fiber.Ctx, gateway *sdk.GatewayClient) error {
	reply, err := gateway.GetCredential(ID)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	clientId := ""
	for _, item := range reply.GetContent() {
		if item.Key == ClientIdKey {
			clientId = item.Value
		}
	}
	v.clientId = clientId

	appRedirectURI := v.AuthorizeURL()
	return c.Redirect(appRedirectURI, http.StatusFound)
}

func (v *Dropbox) StoreAccessToken(c *fiber.Ctx, gateway *sdk.GatewayClient) error {
	code := c.FormValue("code")
	reply, err := gateway.GetCredential(ID)
	if err != nil {
		return err
	}
	clientId := ""
	clientSecret := ""
	for _, item := range reply.GetContent() {
		if item.Key == ClientIdKey {
			clientId = item.Value
		}
		if item.Key == ClientSecretKey {
			clientSecret = item.Value
		}
	}
	v.clientId = clientId
	v.clientSecret = clientSecret

	tokenResp, err := v.GetAccessToken(code)
	if err != nil {
		return err
	}

	extra, err := json.Marshal(&tokenResp)
	if err != nil {
		return err
	}
	appReply, err := gateway.StoreAppOAuth(&pb.AppRequest{
		App: &pb.App{
			Name:  ID,
			Type:  ID,
			Token: v.accessToken,
			Extra: util.ByteToString(extra),
		},
	})
	if err != nil {
		return err
	}
	if appReply.GetState() {
		return nil
	}
	return errors.New("error")
}

func (v *Dropbox) Upload(path string, content io.Reader) error {
	apiArg, err := json.Marshal(map[string]interface{}{
		"path":            path,
		"mode":            "add",
		"autorename":      true,
		"mute":            false,
		"strict_conflict": false,
	})
	if err != nil {
		return err
	}
	resp, err := v.c.R().
		SetAuthToken(v.accessToken).
		SetHeader("Content-Type", "application/octet-stream").
		SetHeader("Dropbox-API-Arg", util.ByteToString(apiArg)).
		SetContentLength(true).
		SetBody(content).
		Post("https://content.dropboxapi.com/2/files/upload")
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	} else {
		return fmt.Errorf("%d, %s", resp.StatusCode(), util.ByteToString(resp.Body()))
	}
}
