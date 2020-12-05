package cron

import (
	"encoding/json"
	"fmt"
	"github.com/gorhill/cronexpr"
	"github.com/tsundata/assistant/internal/app/message/bot"
	"github.com/tsundata/assistant/internal/pkg/model"
	"strings"
	"sync"
	"time"
)

type Rule struct {
	When   string
	Action func() []model.Event
}

type cronRuleset struct {
	outCh     chan model.Event
	cronRules map[string]Rule

	mu            sync.Mutex
	attachedCrons map[string][]string
	stopChan      []chan struct{}
}

// Name returns this rules name - meant for debugging.
func (r *cronRuleset) Name() string {
	return "Cron Ruleset"
}

// Boot runs preparatory steps for ruleset execution
func (r *cronRuleset) Boot(self *bot.Bot) {
	r.start(self)
}

func (r cronRuleset) HelpMessage(self *bot.Bot, _ model.Event) string {
	helpMsg := fmt.Sprintln("cron attach <job name>- attach one cron job")
	helpMsg = fmt.Sprintln(helpMsg, "cron detach <job name> - detach one cron job")
	helpMsg = fmt.Sprintln(helpMsg, "cron list - list all available crons")
	helpMsg = fmt.Sprintln(helpMsg, "cron start - start all crons")
	helpMsg = fmt.Sprintln(helpMsg, "cron stop - stop all crons")

	return helpMsg
}

func (r *cronRuleset) ParseMessage(self *bot.Bot, in model.Event) []model.Event {
	if strings.HasPrefix(in.Data.Message.Text, "cron attach") {
		ruleName := strings.TrimSpace(strings.TrimPrefix(in.Data.Message.Text, "cron attach"))
		ret := []model.Event{{
			Data: model.EventData{Message: model.Message{
				Text: r.attach(self, ruleName, "in.Room"),
			}},
		}}
		r.start(self)
		return ret
	}

	if strings.HasPrefix(in.Data.Message.Text, "cron detach") {
		ruleName := strings.TrimSpace(strings.TrimPrefix(in.Data.Message.Text, "cron detach"))
		return []model.Event{{
			Data: model.EventData{Message: model.Message{
				Text: r.attach(self, ruleName, "in.Room"),
			}},
		}}
	}

	if in.Data.Message.Text == "cron list" {
		var ret []model.Event
		for ruleName, rule := range r.cronRules {
			ret = append(ret, model.Event{
				Data: model.EventData{Message: model.Message{
					Text: "@" + rule.When + " " + ruleName,
				}},
			})
		}
		return ret
	}

	if in.Data.Message.Text == "cron start" {
		r.start(self)
		return []model.Event{
			{
				Data: model.EventData{Message: model.Message{
					Text: "all cron jobs started",
				}},
			},
		}
	}

	if in.Data.Message.Text == "cron stop" {
		r.stop()
		return []model.Event{
			{
				Data: model.EventData{Message: model.Message{
					Text: "all cron jobs stopped",
				}},
			},
		}
	}

	return []model.Event{}
}

func (r *cronRuleset) attach(self *bot.Bot, ruleName, room string) string {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.cronRules[ruleName]; !ok {
		return ruleName + " not found"
	}

	for _, rn := range r.attachedCrons[room] {
		if rn == ruleName {
			return ruleName + " already attached to this room"
		}
	}
	r.attachedCrons[room] = append(r.attachedCrons[room], ruleName)

	b, err := json.Marshal(r.attachedCrons)
	if err != nil {
		return fmt.Sprintf("error attaching %s: %v", ruleName, err)
	}

	// self.MemorySave("cron", "attached", b)
	fmt.Println(b)
	return ruleName + " attached to this room"
}

func (r *cronRuleset) detach(self *bot.Bot, ruleName, room string) string {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.attachedCrons[room]; !ok {
		return "room not found in cron memory"
	}

	var newRoom []string
	for _, rn := range r.attachedCrons[room] {
		if rn == ruleName {
			continue
		}
		newRoom = append(newRoom, rn)
	}
	r.attachedCrons[room] = newRoom

	b, err := json.Marshal(r.attachedCrons)
	if err != nil {
		return fmt.Sprintf("error detaching %s: %v", ruleName, err)
	}
	// self.MemorySave("cron", "attached", b)
	fmt.Println(b)
	return ruleName + " detached to this room"
}

func (r *cronRuleset) start(self *bot.Bot) {
	r.stop()

	r.mu.Lock()
	defer r.mu.Unlock()

	// process cron
	for rule, _ := range r.cronRules {
		c := make(chan struct{})
		r.stopChan = append(r.stopChan, c)
		go processCronRule(r.cronRules[rule], c, r.outCh, "room")
	}

	// send message
	go func() {
		for {
			select {
			case out := <-r.outCh:
				self.Send(out)
			default:

			}
		}
	}()
}

func processCronRule(rule Rule, stop chan struct{}, outCh chan model.Event, cronRoom string) {
	nextTime := cronexpr.MustParse(rule.When).Next(time.Now())
	for {
		select {
		case <-stop:
			return
		default:
			if nextTime.Format("2006-01-02 15:04") == time.Now().Format("2006-01-02 15:04") {
				msgs := rule.Action()
				for _, msg := range msgs {
					// msg.Room = cronRoom
					outCh <- msg
				}
			}
			nextTime = cronexpr.MustParse(rule.When).Next(time.Now())
			time.Sleep(2 * time.Second)
		}
	}
}

func (r *cronRuleset) stop() {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, c := range r.stopChan {
		c <- struct{}{}
	}
	r.stopChan = []chan struct{}{}
}

// New returns a cron rule set
func New(rules map[string]Rule) *cronRuleset {
	r := &cronRuleset{
		attachedCrons: make(map[string][]string),
		cronRules:     rules,
		outCh:         make(chan model.Event, 10),
	}
	return r
}
