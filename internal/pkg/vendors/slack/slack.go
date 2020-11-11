package slack

import (
	"bytes"
	"encoding/json"
	"github.com/slack-go/slack"
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
func SlashShortcutParse(r *http.Request) (s SlashShortcut, err error) {
	if err = r.ParseForm(); err != nil {
		return s, err
	}

	payload := r.PostForm.Get("payload")

	err = json.Unmarshal([]byte(payload), &s)
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
