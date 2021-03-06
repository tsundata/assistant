package rule

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/influxdata/cron"
	"github.com/tsundata/assistant/internal/app/cron/pipeline"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"time"
)

type Rule struct {
	Name   string
	When   string
	Action func(ctx rulebot.IContext) []result.Result
}

type cronRuleset struct {
	outCh     chan result.Result
	cronRules []Rule
}

// New returns a cron rule set
func New(rules []Rule) *cronRuleset {
	r := &cronRuleset{
		cronRules: rules,
		outCh:     make(chan result.Result, 100),
	}
	return r
}

// Name returns this rule name - meant for debugging.
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
	b.Ctx.GetLogger().Info("cron starting...")

	// process cron
	for rule := range r.cronRules {
		b.Ctx.GetLogger().Info("cron " + r.cronRules[rule].Name + ": start...")
		go r.ruleWorker(b, r.cronRules[rule])
	}

	// result pipeline
	go r.resultWorker(b)
	go r.resultWorker(b)
}

func (r *cronRuleset) ruleWorker(b *rulebot.RuleBot, rule Rule) {
	p, err := cron.ParseUTC(rule.When)
	if err != nil {
		b.Ctx.GetLogger().Error(err)
		return
	}
	nextTime, err := p.Next(time.Now())
	if err != nil {
		b.Ctx.GetLogger().Error(err)
		return
	}
	for {
		if nextTime.Format("2006-01-02 15:04") == time.Now().Format("2006-01-02 15:04") {
			msgs := func() []result.Result {
				defer func() {
					if r := recover(); r != nil {
						b.Ctx.GetLogger().Info("ruleWorker recover " + rule.Name)
						if v, ok := r.(error); ok {
							b.Ctx.GetLogger().Error(v)
						}
					}
				}()
				return rule.Action(b.Ctx)
			}()
			if len(msgs) > 0 {
				for _, item := range msgs {
					r.outCh <- item
				}
			}
		}
		nextTime, err = p.Next(time.Now())
		if err != nil {
			b.Ctx.GetLogger().Error(err)
			continue
		}
		time.Sleep(2 * time.Second)
	}
}

func (r *cronRuleset) resultWorker(b *rulebot.RuleBot) {
	for out := range r.outCh {
		// filter
		res := r.filter(b.Ctx, out)
		// pipeline
		r.pipeline(b.Ctx, res)
	}
}

func (r *cronRuleset) filter(ctx rulebot.IContext, res result.Result) result.Result {
	ctxB := context.Background()
	filterKey := fmt.Sprintf("cron:%d:filter", res.Kind)

	// filter
	state := ctx.GetRedis().SIsMember(ctxB, filterKey, res.ID)
	ex, err := state.Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return result.EmptyResult()
	}
	if ex {
		return result.EmptyResult()
	}

	// add
	ctx.GetRedis().SAdd(ctxB, filterKey, res.ID)

	return res
}

func (r *cronRuleset) pipeline(ctx rulebot.IContext, res result.Result) {
	if res.ID == "" {
		return
	}
	pipeline.Workflow(ctx, res)
}
