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

type ListResponse struct {
	Status int             `json:"status"`
	List   map[string]Item `json:"list"`
}

type Item struct {
	Id            string `json:"item_id"`
	ResolvedId    string `json:"resolved_id"`
	GivenUrl      string `json:"given_url"`
	GivenTitle    string `json:"given_title"`
	Favorite      string `json:"favorite"`
	Status        string `json:"status"`
	TimeAdded     string `json:"time_added"`
	TimeUpdated   string `json:"time_updated"`
	TimeRead      string `json:"time_read"`
	TimeFavorited string `json:"time_favorited"`
	ResolvedTitle string `json:"resolved_title"`
	ResolvedUrl   string `json:"resolved_url"`
	Excerpt       string `json:"excerpt"`
	IsArticle     string `json:"is_article"`
	IsIndex       string `json:"is_index"`
	HasVideo      string `json:"has_video"`
	HasImage      string `json:"has_image"`
	WordCount     string `json:"word_count"`
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

func (p *Pocket) Retrieve(accessToken string, count int) (*ListResponse, error) {
	resp, err := p.c.R().
		SetResult(&ListResponse{}).
		SetBody(map[string]interface{}{
			"consumer_key": p.ConsumerKey,
			"access_token": accessToken,
			"count":        count,
			"detailType":   "simple",
			"state":        "all",
			"sort":         "newest",
		}).
		Post("/v3/get")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.Result().(*ListResponse), nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)", resp.StatusCode(), resp.Header().Get("X-Error-Code"), resp.Header().Get("X-Error"))
	}
}
