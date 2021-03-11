package dropbox

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"io"
	"net/http"
	"time"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	UID         string `json:"uid"`
	AccountID   string `json:"account_id"`
	Scope       string `json:"scope"`
}

type Dropbox struct {
	c        *resty.Client
	ClientId string
}

func NewDropbox(clientId string) *Dropbox {
	v := &Dropbox{ClientId: clientId}

	v.c = resty.New()
	v.c.SetHostURL("https://api.dropboxapi.com")
	v.c.SetTimeout(time.Minute)

	return v
}

func (v *Dropbox) AuthorizeURL(redirectURI string) string {
	return fmt.Sprintf("https://www.dropbox.com/oauth2/authorize?client_id=%s&response_type=code&redirect_uri=%s", v.ClientId, redirectURI)
}

func (v *Dropbox) GetAccessToken(clientSecret, redirectURI, code string) (*TokenResponse, error) {
	resp, err := v.c.R().
		SetBasicAuth(v.ClientId, clientSecret).
		SetFormData(map[string]string{
			"code":         code,
			"grant_type":   "authorization_code",
			"redirect_uri": redirectURI,
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
		return &result, nil
	} else {
		return nil, fmt.Errorf("%d, %s", resp.StatusCode(), utils.ByteToString(resp.Body()))
	}
}

func (v *Dropbox) Upload(accessToken string, path string, content io.Reader) error {
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
		SetAuthToken(accessToken).
		SetHeader("Content-Type", "application/octet-stream").
		SetHeader("Dropbox-API-Arg", utils.ByteToString(apiArg)).
		SetContentLength(true).
		SetBody(content).
		Post("https://content.dropboxapi.com/2/files/upload")
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	} else {
		return fmt.Errorf("%d, %s", resp.StatusCode(), utils.ByteToString(resp.Body()))
	}
}
