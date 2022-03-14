package pushover

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tsundata/assistant/internal/pkg/push"
	"net/http"
	"strconv"
	"time"
)

const (
	ID       = "pushover"
	TokenKey = "token"
	UserKey  = "user"
)

type Response struct {
	Status  int         `json:"status"`
	Request string      `json:"request"`
	Errors  interface{} `json:"errors"`
}

type Limitation struct {
	Limit     int64
	Remaining int64
	Reset     int64
}

type Pushover struct {
	c     *resty.Client
	token string
	user  string
}

func NewPushover(user, token string) *Pushover {
	v := &Pushover{}
	v.user = user
	v.token = token
	v.c = resty.New()
	v.c.SetBaseURL("https://api.pushover.net/1")
	v.c.SetTimeout(time.Minute)
	return v
}

func (p *Pushover) Send(message push.Message) error {
	resp, err := p.c.R().
		SetResult(&Response{}).
		SetBody(map[string]interface {
		}{
			"token":   p.token,
			"user":    p.user,
			"title":   message.Title,
			"message": message.Content,
		}).
		Post("/messages.json")
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		result := resp.Result().(*Response)
		if result.Status == 1 {
			return nil
		}
		r, _ := json.Marshal(result.Errors)
		return errors.New(string(r))
	} else {
		return fmt.Errorf("pushover error %d", resp.StatusCode())
	}
}

func (p *Pushover) Limitations() (*Limitation, error) {
	resp, err := p.c.R().
		SetResult(&Response{}).
		SetQueryParam("token", p.token).
		Get("/apps/limits.json")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		result := resp.Result().(*Response)
		if result.Status == 1 {
			limit, _ := strconv.ParseInt(resp.Header().Get("X-Limit-App-Limit"), 10, 64)
			remaining, _ := strconv.ParseInt(resp.Header().Get("X-Limit-App-Remaining"), 10, 64)
			reset, _ := strconv.ParseInt(resp.Header().Get("X-Limit-App-Reset"), 10, 64)
			return &Limitation{
				Limit:     limit,
				Remaining: remaining,
				Reset:     reset,
			}, nil
		}

		return nil, errors.New("error status")
	} else {
		return nil, fmt.Errorf("pushover error %d", resp.StatusCode())
	}
}
