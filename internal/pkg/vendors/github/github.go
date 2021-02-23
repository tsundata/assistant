package github

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
}

type Github struct {
	c        *resty.Client
	ClientId string
}

func NewGithub(clientId string) *Github {
	v := &Github{ClientId: clientId}

	v.c = resty.New()
	v.c.SetHostURL("https://api.github.com")
	v.c.SetTimeout(time.Minute)

	return v
}

func (v *Github) AuthorizeURL(redirectURI string) string {
	return fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", v.ClientId, redirectURI)
}

func (v *Github) GetAccessToken(clientSecret, code string) (*TokenResponse, error) {
	resp, err := v.c.R().
		SetResult(&TokenResponse{}).
		SetHeader("Accept", "application/json").
		SetBody(map[string]interface{}{
			"client_id":     v.ClientId,
			"client_secret": clientSecret,
			"code":          code,
		}).
		Post("https://github.com/login/oauth/access_token")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.Result().(*TokenResponse), nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)", resp.StatusCode(), resp.Header().Get("X-Error-Code"), resp.Header().Get("X-Error"))
	}
}
