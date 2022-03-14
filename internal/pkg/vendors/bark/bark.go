package bark

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tsundata/assistant/internal/pkg/push"
	"github.com/tsundata/assistant/internal/pkg/util"
	"net/http"
	"time"
)

const (
	ID        = "bark"
	ServerUrl = "server_url"
	TokenKey  = "token"
)

type Bark struct {
	c     *resty.Client
	token string
}

func NewBark(serverUrl, token string) *Bark {
	v := &Bark{token: token}
	v.c = resty.New()
	if util.IsUrl(serverUrl) {
		v.c.SetBaseURL(serverUrl)
	} else {
		v.c.SetBaseURL("https://api.day.app")
	}
	v.c.SetTimeout(time.Minute)
	return v
}

func (b *Bark) Send(message push.Message) error {
	resp, err := b.c.R().
		SetQueryParams(map[string]string{
			"url":   message.Url,
			"sound": message.Sound,
		}).
		Post(fmt.Sprintf("/%s/%s", message.Title, message.Content))
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	} else {
		return fmt.Errorf("pushover error %d", resp.StatusCode())
	}
}
