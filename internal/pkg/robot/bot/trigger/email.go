package trigger

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors/email"
	"regexp"
	"strconv"
	"strings"
)

type Email struct {
	text  string
	email []string
}

func NewEmail() *Email {
	return &Email{}
}

func (t *Email) Cond(message *pb.Message) bool {
	re := regexp.MustCompile(`(?m)` + util.EmailRegex)
	ts := re.FindAllString(message.GetText(), -1)

	if len(ts) == 0 {
		return false
	}

	t.text = message.GetText()
	for _, item := range ts {
		t.text = strings.ReplaceAll(t.text, item, "")
		t.text = strings.ReplaceAll(t.text, "<mailto:|>", "") // special
		t.email = append(t.email, item)
	}

	t.email = clear(t.email)

	return true
}

func (t *Email) Handle(ctx context.Context, botCtx bot.Context, comp component.Component) {
	reply, err := comp.Middle().GetCredential(ctx, &pb.CredentialRequest{Type: email.ID})
	if err != nil {
		return
	}
	host := ""
	port := ""
	username := ""
	password := ""
	for _, item := range reply.GetContent() {
		if item.Key == email.Host {
			host = item.Value
		}
		if item.Key == email.Port {
			port = item.Value
		}
		if item.Key == email.Username {
			username = item.Value
		}
		if item.Key == email.Password {
			password = item.Value
		}
	}
	if host == "" || port == "" || username == "" || password == "" {
		return
	}

	p, err := strconv.Atoi(port)
	if err != nil {
		return
	}

	config := email.Config{
		Host:     host,
		Port:     p,
		Username: username,
		Password: password,
	}

	for _, mail := range t.email {
		err = email.SendEmail(&config, mail, t.text, fmt.Sprintf(`%s <br>
---------------- <br>
Sended by Assistant
`, t.text))
		if err != nil {
			comp.GetLogger().Error(err)
			return
		}
		err = comp.GetBus().Publish(ctx, enum.Message, event.MessageChannelSubject, pb.Message{
			UserId:   botCtx.Message.GetUserId(),
			GroupId:  botCtx.Message.GetGroupId(),
			Text:     fmt.Sprintf("Sended to Mail: %s", mail),
			Type:     string(enum.MessageTypeText),
			Sequence: botCtx.Message.GetSequence(),
			SendTime: util.Now(),
		})
		if err != nil {
			return
		}
	}
}
