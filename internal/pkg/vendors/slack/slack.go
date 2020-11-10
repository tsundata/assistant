package slack

import (
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
