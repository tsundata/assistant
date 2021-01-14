package rules

import (
	"encoding/json"
	"fmt"
	"github.com/influxdata/cron"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"log"
	"strings"
	"sync"
	"time"
)

type Rule struct {
	When   string
	Action func() []model.Message
}

type cronRuleset struct {
	outCh     chan model.Message
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
func (r *cronRuleset) Boot(b *rulebot.RuleBot) {
	r.start(b)
	r.send(b)
}

func (r *cronRuleset) HelpMessage(b *rulebot.RuleBot, _ model.Message) string {
	helpMsg := fmt.Sprintln("cron attach <job name>- attach one cron job")
	helpMsg = fmt.Sprintln(helpMsg, "cron detach <job name> - detach one cron job")
	helpMsg = fmt.Sprintln(helpMsg, "cron list - list all available crons")
	helpMsg = fmt.Sprintln(helpMsg, "cron start - start all crons")
	helpMsg = fmt.Sprintln(helpMsg, "cron stop - stop all crons")

	return helpMsg
}

func (r *cronRuleset) ParseMessage(b *rulebot.RuleBot, in model.Message) []model.Message {
	if strings.HasPrefix(in.Text, "cron attach") {
		ruleName := strings.TrimSpace(strings.TrimPrefix(in.Text, "cron attach"))
		ret := []model.Message{{
			Text: r.attach(b, ruleName, "in.Room"),
		}}
		r.start(b)
		return ret
	}

	if strings.HasPrefix(in.Text, "cron detach") {
		ruleName := strings.TrimSpace(strings.TrimPrefix(in.Text, "cron detach"))
		return []model.Message{{
			Text: r.attach(b, ruleName, "in.Room"),
		}}
	}

	if in.Text == "cron list" {
		var ret []model.Message
		for ruleName, rule := range r.cronRules {
			ret = append(ret, model.Message{
				Text: "@" + rule.When + " " + ruleName,
			})
		}
		return ret
	}

	if in.Text == "cron start" {
		r.start(b)
		return []model.Message{{
			Text: "all cron jobs started",
		}}
	}

	if in.Text == "cron stop" {
		r.stop()
		return []model.Message{{
			Text: "all cron jobs started",
		}}
	}

	return []model.Message{}
}

func (r *cronRuleset) attach(_ *rulebot.RuleBot, ruleName, room string) string {
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

	a, err := json.Marshal(r.attachedCrons)
	if err != nil {
		return fmt.Sprintf("error attaching %s: %v", ruleName, err)
	}

	// b.MemorySave("cron", "attached", a)
	log.Println(a)
	return ruleName + " attached to this room"
}

// nolint:unused
func (r *cronRuleset) detach(_ *rulebot.RuleBot, ruleName, room string) string {
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

	a, err := json.Marshal(r.attachedCrons)
	if err != nil {
		return fmt.Sprintf("error detaching %s: %v", ruleName, err)
	}
	// b.MemorySave("cron", "attached", a)
	log.Println(a)
	return ruleName + " detached to this room"
}

func (r *cronRuleset) start(_ *rulebot.RuleBot) {
	r.stop()

	r.mu.Lock()
	defer r.mu.Unlock()

	// process cron
	for rule := range r.cronRules {
		c := make(chan struct{})
		r.stopChan = append(r.stopChan, c)
		go processCronRule(r.cronRules[rule], c, r.outCh, "room")
	}
}

// send message
func (r *cronRuleset) send(b *rulebot.RuleBot) {
	go func() {
		for out := range r.outCh {
			b.Send(out)
		}
	}()
}

func processCronRule(rule Rule, stop chan struct{}, outCh chan model.Message, _ string) {
	p, err := cron.ParseUTC(rule.When)
	if err != nil {
		log.Println(err)
		return
	}
	nextTime, err := p.Next(time.Now())
	if err != nil {
		log.Println(err)
		return
	}
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
			nextTime, err = p.Next(time.Now())
			if err != nil {
				log.Println(err)
				continue
			}
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
		outCh:         make(chan model.Message, 10),
	}
	return r
}
