package pocket

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type CodeResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

type Pocket struct {
	c           *resty.Client
	ConsumerKey string
}

func NewPocket(consumerKey string) *Pocket {
	p := &Pocket{ConsumerKey: consumerKey}

	p.c = resty.New()
	p.c.SetHostURL("https://getpocket.com")
	p.c.SetTimeout(time.Minute)

	return p
}

func (p *Pocket) GetCode(redirectURI, state string) (*CodeResponse, error) {
	resp, err := p.c.R().
		SetResult(&CodeResponse{}).
		SetHeader("X-Accept", "application/json").
		SetBody(map[string]interface{}{"consumer_key": p.ConsumerKey, "redirect_uri": redirectURI, "state": state}).
		Post("/v3/oauth/request")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.Result().(*CodeResponse), nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)", resp.StatusCode(), resp.Header().Get("X-Error-Code"), resp.Header().Get("X-Error"))
	}
}

func (p *Pocket) AuthorizeURL(code, redirectURI string) string {
	return fmt.Sprintf("https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=%s", code, redirectURI)
}

func (p *Pocket) GetAccessToken(code string) (*TokenResponse, error) {
	resp, err := p.c.R().
		SetResult(&TokenResponse{}).
		SetHeader("X-Accept", "application/json").
		SetBody(map[string]interface{}{"consumer_key": p.ConsumerKey, "code": code}).
		Post("/v3/oauth/authorize")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.Result().(*TokenResponse), nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)", resp.StatusCode(), resp.Header().Get("X-Error-Code"), resp.Header().Get("X-Error"))
	}
}
