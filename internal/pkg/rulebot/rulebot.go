package rulebot

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/valyala/fasthttp"
	"log"
	"strings"
)

type RuleBot struct {
	name        string
	providerIn  model.Event
	providerOut []model.Event
	rules       []RuleParser

	webhook   string
	SubClient pb.SubscribeClient
	MidClient pb.MiddleClient
}

func New(name string, v *viper.Viper, SubClient pb.SubscribeClient, MidClient pb.MiddleClient, opts ...Option) *RuleBot {
	s := &RuleBot{
		name: name,
	}

	for _, opt := range opts {
		opt(s)
	}

	slack := v.GetStringMapString("slack")
	s.webhook = slack["webhook"]
	s.SubClient = SubClient
	s.MidClient = MidClient

	return s
}

func (s *RuleBot) Process(in model.Event) *RuleBot {
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
		s.providerOut = append(s.providerOut, responses...)
	}
	return s
}

func (s *RuleBot) MessageProviderOut() []model.Event {
	return s.providerOut
}

func (s *RuleBot) Name() string {
	return s.name
}

func (s *RuleBot) Send(out model.Event) {
	client := http.NewClient()
	resp, err := client.PostJSON(s.webhook, map[string]interface{}{
		"text": out.Data.Message.Text,
	})

	if err != nil {
		log.Println(err)
		return
	}

	fasthttp.ReleaseResponse(resp)
}

type Option func(*RuleBot)

type RuleParser interface {
	Name() string
	Boot(*RuleBot)
	ParseMessage(*RuleBot, model.Event) []model.Event
	HelpMessage(*RuleBot, model.Event) string
}

func RegisterRuleset(rule RuleParser) Option {
	return func(s *RuleBot) {
		log.Printf("registering ruleset %T", rule)
		rule.Boot(s)
		s.rules = append(s.rules, rule)
	}
}