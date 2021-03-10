package dropbox

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	AccountID   string `json:"account_id"`
	UID         string `json:"uid"`
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
		SetResult(&TokenResponse{}).
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
		return resp.Result().(*TokenResponse), nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)", resp.StatusCode(), resp.Header().Get("X-Error-Code"), resp.Header().Get("X-Error"))
	}
}
