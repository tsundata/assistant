package rules

import (
	"github.com/influxdata/cron"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"log"
	"sync"
	"time"
)

type Rule struct {
	When   string
	Action func(b *rulebot.RuleBot) []string
}

type cronRuleset struct {
	outCh     chan string
	cronRules map[string]Rule

	mu       sync.Mutex
	stopChan []chan struct{}
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

func (r *cronRuleset) HelpMessage(_ *rulebot.RuleBot, _ model.Message) string {
	return ""
}

func (r *cronRuleset) ParseMessage(_ *rulebot.RuleBot, _ model.Message) []model.Message {
	return []model.Message{}
}

func (r *cronRuleset) start(b *rulebot.RuleBot) {
	r.stop()

	r.mu.Lock()
	defer r.mu.Unlock()

	// process cron
	for rule := range r.cronRules {
		c := make(chan struct{})
		r.stopChan = append(r.stopChan, c)
		go processCronRule(b, r.cronRules[rule], c, r.outCh)
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

// send message
func (r *cronRuleset) send(b *rulebot.RuleBot) {
	go func() {
		for out := range r.outCh {
			b.Send(out)
		}
	}()
}

func processCronRule(b *rulebot.RuleBot, rule Rule, stop chan struct{}, outCh chan string) {
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
				msgs := rule.Action(b)
				for _, msg := range msgs {
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

// New returns a cron rule set
func New(rules map[string]Rule) *cronRuleset {
	r := &cronRuleset{
		cronRules: rules,
		outCh:     make(chan string, 10),
	}
	return r
}
