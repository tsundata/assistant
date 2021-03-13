package rules

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/influxdata/cron"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/tsundata/assistant/internal/pkg/version"
	"log"
	"strings"
	"time"
)

type Rule struct {
	Name   string
	When   string
	Action func(b *rulebot.RuleBot) []string
}

type Result struct {
	Name   string
	Result []string
}

type cronRuleset struct {
	outCh     chan Result
	cronRules []Rule
}

// New returns a cron rule set
func New(rules []Rule) *cronRuleset {
	r := &cronRuleset{
		cronRules: rules,
		outCh:     make(chan Result, 10),
	}
	return r
}

// Name returns this rules name - meant for debugging.
func (r *cronRuleset) Name() string {
	return "Cron Ruleset"
}

// Boot runs preparatory steps for ruleset execution
func (r *cronRuleset) Boot(b *rulebot.RuleBot) {
	r.daemon(b)
}

func (r *cronRuleset) HelpMessage(_ *rulebot.RuleBot, _ string) string {
	return ""
}

func (r *cronRuleset) ParseMessage(_ *rulebot.RuleBot, _ string) []string {
	return []string{}
}

func (r *cronRuleset) daemon(b *rulebot.RuleBot) {
	log.Println("cron starting...")

	// process cron
	for rule := range r.cronRules {
		log.Println("cron " + r.cronRules[rule].Name + ": start...")
		go r.ruleWorker(b, r.cronRules[rule])
	}

	// send message
	go r.resultWorker(b)
}

func (r *cronRuleset) ruleWorker(b *rulebot.RuleBot, rule Rule) {
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
		if nextTime.Format("2006-01-02 15:04") == time.Now().Format("2006-01-02 15:04") {
			msgs := func() []string {
				defer func() {
					if r := recover(); r != nil {
						log.Println("ruleWorker recover ", rule.Name)
						if v, ok := r.(error); ok {
							log.Println(v)
						}
					}
				}()
				return rule.Action(b)
			}()
			if len(msgs) > 0 {
				r.outCh <- Result{
					Name:   rule.Name,
					Result: msgs,
				}
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

func (r *cronRuleset) resultWorker(b *rulebot.RuleBot) {
	for out := range r.outCh {
		// filter
		diff := r.filter(b, out.Name, out.Result)
		// send
		r.send(b, out.Name, diff)
	}
}

func (r *cronRuleset) filter(b *rulebot.RuleBot, name string, latest []string) []string {
	ctx := context.Background()
	sentKey := fmt.Sprintf("cron:%s:sent", name)
	todoKey := fmt.Sprintf("cron:%s:todo", name)
	sendTimeKey := fmt.Sprintf("cron:%s:sendtime", name)

	// sent
	smembers := b.RDB.SMembers(ctx, sentKey)
	old, err := smembers.Result()
	if err != nil && err != redis.Nil {
		return []string{}
	}

	// to do
	smembers = b.RDB.SMembers(ctx, todoKey)
	todo, err := smembers.Result()
	if err != nil && err != redis.Nil {
		return []string{}
	}

	// merge
	tobeCompared := append(old, todo...)

	// diff
	diff := utils.StringSliceDiff(latest, tobeCompared)

	// record
	b.RDB.Set(ctx, sendTimeKey, time.Now().Unix(), 0)

	// add data
	for _, item := range diff {
		b.RDB.SAdd(ctx, sentKey, item)
	}
	b.RDB.Expire(ctx, sentKey, 7*24*time.Hour)

	// clear to do
	b.RDB.Del(ctx, todoKey)

	return diff
}

func (r *cronRuleset) send(b *rulebot.RuleBot, name string, out []string) {
	if len(out) == 0 {
		return
	}

	text := fmt.Sprintf("Cron %s (v%s)\n%s", name, version.Version, strings.Join(out, "\n"))

	_, err := b.MsgClient.Send(context.Background(), &pb.MessageRequest{
		Text: text,
	})
	if err != nil {
		return
	}
}
