package bot

import (
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/valyala/fasthttp"
	"log"
	"strings"
)

type Bot struct {
	name        string
	providerIn  model.Event
	providerOut []model.Event
	rules       []RuleParser

	slackWebhook string
}

func New(name string, slackWebhook string, opts ...Option) *Bot {
	s := &Bot{
		name: name,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.slackWebhook = slackWebhook

	return s
}

func (s *Bot) Process(in model.Event) *Bot {
	log.Println("plugin process event")

	s.providerIn = in
	s.providerOut = []model.Event{}
	if strings.HasPrefix(in.Data.Message.Text, s.Name()+" help") {
		helpMsg := fmt.Sprintln("available commands:")
		for _, rule := range s.rules {
			helpMsg = fmt.Sprintln(helpMsg, rule.HelpMessage(s, in))
		}
		s.providerOut = append(s.providerOut, model.Event{
			Data: model.EventData{
				Type: model.EventTypeMessage,
				Message: model.Message{
					Type: model.MessageTypeText,
					Text: helpMsg,
				},
			},
		})
		return s
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic recovered when parsing message: %#v. Panic: %v", in, r)
		}
	}()
	for _, rule := range s.rules {
		responses := rule.ParseMessage(s, in)
		for _, r := range responses {
			s.providerOut = append(s.providerOut, r)
		}
	}
	return s
}

func (s *Bot) MessageProviderOut() []model.Event {
	return s.providerOut
}

func (s *Bot) Name() string {
	return s.name
}

func (s *Bot) Send(out model.Event) {
	fmt.Printf("send event : %v\n", out)

	client := http.NewClient()
	resp, err := client.PostJSON(s.slackWebhook, map[string]interface{}{
		"text": out.Data.Message.Text,
	})

	if err != nil {
		log.Println(err)
		return
	}

	fasthttp.ReleaseResponse(resp)
}

type Option func(*Bot)

type RuleParser interface {
	Name() string
	Boot(*Bot)
	ParseMessage(*Bot, model.Event) []model.Event
	HelpMessage(*Bot, model.Event) string
}

func RegisterRuleset(rule RuleParser) Option {
	return func(s *Bot) {
		log.Printf("registering ruleset %T", rule)
		rule.Boot(s)
		s.rules = append(s.rules, rule)
	}
}

type pluginRuleset struct {
	pluginBins []string
	plugins    []RuleParser
}
