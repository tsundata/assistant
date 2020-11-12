package slack

import (
	"bytes"
	"encoding/json"
	"github.com/slack-go/slack"
	"github.com/valyala/fasthttp"
	"net/http"
)

func SecretsVerifier(header http.Header, body []byte, secret string) error {
	sv, err := slack.NewSecretsVerifier(header, secret)
	if err != nil {
		return err
	}
	_, err = sv.Write(body)
	if err != nil {
		return err
	}
	return sv.Ensure()
}

func ResponseText(responseURL, text string) error {
	res := map[string]string{"text": text}
	j, err := json.Marshal(&res)
	if err != nil {
		return err
	}
	_, err = http.Post(responseURL, "application/json", bytes.NewBuffer(j))
	return err
}

// SlashShortcut contains information about a request of the slash command
type SlashShortcut struct {
	Token      string    `json:"token"`
	Type       string    `json:"type"`
	ActionTs   string    `json:"action_ts"`
	Team       SlashTeam `json:"team,omitempty"`
	User       SlashUser `json:"user,omitempty"`
	CallbackID string    `json:"callback_id"`
	TriggerID  string    `json:"trigger_id"`
}

type SlashUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	TeamID   string `json:"team_id"`
}

type SlashTeam struct {
	ID     string `json:"id"`
	Domain string `json:"domain"`
}

// SlashShortcutParse will parse the request of the slash command
func SlashShortcutParse(r *fasthttp.Request) (s SlashShortcut, err error) {
	payload := r.PostArgs().Peek("payload")
	err = json.Unmarshal(payload, &s)
	return
}

// ValidateToken validates verificationTokens
func (s SlashShortcut) ValidateToken(verificationTokens ...string) bool {
	for _, token := range verificationTokens {
		if s.Token == token {
			return true
		}
	}
	return false
}

// SlashCommand contains information about a request of the slash command
type SlashCommand struct {
	Token          string `json:"token"`
	TeamID         string `json:"team_id"`
	TeamDomain     string `json:"team_domain"`
	EnterpriseID   string `json:"enterprise_id,omitempty"`
	EnterpriseName string `json:"enterprise_name,omitempty"`
	ChannelID      string `json:"channel_id"`
	ChannelName    string `json:"channel_name"`
	UserID         string `json:"user_id"`
	UserName       string `json:"user_name"`
	Command        string `json:"command"`
	Text           string `json:"text"`
	ResponseURL    string `json:"response_url"`
	TriggerID      string `json:"trigger_id"`
	APIAppID       string `json:"api_app_id"`
}

// SlashCommandParse will parse the request of the slash command
func SlashCommandParse(r *fasthttp.Request) (s SlashCommand, err error) {
	s.Token = string(r.PostArgs().Peek("token"))
	s.TeamID = string(r.PostArgs().Peek("team_id"))
	s.TeamDomain = string(r.PostArgs().Peek("team_domain"))
	s.EnterpriseID = string(r.PostArgs().Peek("enterprise_id"))
	s.EnterpriseName = string(r.PostArgs().Peek("enterprise_name"))
	s.ChannelID = string(r.PostArgs().Peek("channel_id"))
	s.ChannelName = string(r.PostArgs().Peek("channel_name"))
	s.UserID = string(r.PostArgs().Peek("user_id"))
	s.UserName = string(r.PostArgs().Peek("user_name"))
	s.Command = string(r.PostArgs().Peek("command"))
	s.Text = string(r.PostArgs().Peek("text"))
	s.ResponseURL = string(r.PostArgs().Peek("response_url"))
	s.TriggerID = string(r.PostArgs().Peek("trigger_id"))
	s.APIAppID = string(r.PostArgs().Peek("api_app_id"))
	return s, nil
}

// ValidateToken validates verificationTokens
func (s SlashCommand) ValidateToken(verificationTokens ...string) bool {
	for _, token := range verificationTokens {
		if s.Token == token {
			return true
		}
	}
	return false
}
