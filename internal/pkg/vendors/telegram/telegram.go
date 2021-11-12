package telegram

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type IncomingRequest struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageId int    `json:"message_id"`
	From      User   `json:"from"`
	Chat      Chat   `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}

type User struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type MessageResponse struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type Telegram struct {
	c     *resty.Client
	token string
}

func NewTelegram(token string) *Telegram {
	v := &Telegram{
		token: token,
	}

	v.c = resty.New()
	v.c.SetBaseURL(fmt.Sprintf("https://api.telegram.org/bot%s/", token))
	v.c.SetTimeout(time.Minute)

	return v
}

func (v *Telegram) SendMessage(chatID int, text string) (result *MessageResponse, err error) {
	resp, err := v.c.R().
		SetBody(map[string]interface{}{
			"chat_id": chatID,
			"text":    text,
		}).
		Post("sendMessage")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		err = json.Unmarshal(resp.Body(), &result)
		return
	} else {
		return nil, fmt.Errorf("%d", resp.StatusCode())
	}
}
