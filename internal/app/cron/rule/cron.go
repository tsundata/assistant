package rule

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/influxdata/cron"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"go.uber.org/zap"
	"time"
)

type Rule struct {
	Name   string
	When   string
	Action func(context.Context, component.Component) []result.Result
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

func (r *cronRuleset) HelpRule(_ *rulebot.RuleBot, _ string) string {
	return ""
}

func (r *cronRuleset) ParseRule(_ context.Context, _ *rulebot.RuleBot, _ string) []string {
	return []string{}
}

func (r *cronRuleset) daemon(b *rulebot.RuleBot) {
	b.Comp.GetLogger().Info("cron starting...")

	// process cron
	ctx := md.BuildAuthContext(enum.SuperUserID) // fixme
	for rule := range r.cronRules {
		b.Comp.GetLogger().Info("cron " + r.cronRules[rule].Name + ": start...")
		go r.ruleWorker(ctx, b, r.cronRules[rule])
	}

	// result pipeline
	go r.resultWorker(ctx, b)
	go r.resultWorker(ctx, b)
}

func (r *cronRuleset) ruleWorker(ctx context.Context, b *rulebot.RuleBot, rule Rule) {
	// register cron
	err := b.Comp.GetBus().Publish(ctx, enum.Middle, event.CronRegisterSubject, &pb.CronRequest{Text: rule.Name})
	if err != nil {
		b.Comp.GetLogger().Error(err, zap.String("cron", rule.Name))
		return
	}

	p, err := cron.ParseUTC(rule.When)
	if err != nil {
		b.Comp.GetLogger().Error(err, zap.String("cron", rule.Name))
		return
	}
	nextTime, err := p.Next(time.Now())
	if err != nil {
		b.Comp.GetLogger().Error(err, zap.String("cron", rule.Name))
		return
	}
	for {
		if nextTime.Format("2006-01-02 15:04") == time.Now().Format("2006-01-02 15:04") {
			// check status
			state, err := b.Comp.Middle().GetCronStatus(ctx, &pb.CronRequest{
				Text: rule.Name,
			})
			if err != nil {
				b.Comp.GetLogger().Error(err, zap.String("cron", rule.Name))
				continue
			}
			// stop
			if !state.State {
				time.Sleep(30 * time.Second)
				continue
			}

			b.Comp.GetLogger().Info("cron "+rule.Name+": scheduled", zap.String("cron", rule.Name))
			msgs := func() []result.Result {
				defer func() {
					if r := recover(); r != nil {
						b.Comp.GetLogger().Warn("ruleWorker recover "+rule.Name, zap.String("cron", rule.Name))
						if v, ok := r.(error); ok {
							b.Comp.GetLogger().Error(v, zap.String("cron", rule.Name))
						}
					}
				}()
				return rule.Action(ctx, b.Comp)
			}()
			if len(msgs) > 0 {
				for _, item := range msgs {
					r.outCh <- item
				}
			}
		}
		nextTime, err = p.Next(time.Now())
		if err != nil {
			b.Comp.GetLogger().Error(err, zap.String("cron", rule.Name))
			continue
		}
		time.Sleep(2 * time.Second)
	}
}

func (r *cronRuleset) resultWorker(ctx context.Context, b *rulebot.RuleBot) {
	for out := range r.outCh {
		// filter
		res := r.filter(b.Comp, out)
		// pipeline
		r.pipeline(ctx, b.Comp, res)
	}
}

func (r *cronRuleset) filter(comp component.Component, res result.Result) result.Result {
	compB := context.Background()
	filterKey := fmt.Sprintf("cron:%d:filter", res.Kind)

	// filter
	state := comp.GetRedis().SIsMember(compB, filterKey, res.ID)
	ex, err := state.Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return result.EmptyResult()
	}
	if ex {
		return result.EmptyResult()
	}

	// add
	comp.GetRedis().SAdd(compB, filterKey, res.ID)

	return res
}

func (r *cronRuleset) pipeline(ctx context.Context, comp component.Component, res result.Result) {
	if res.ID == "" {
		return
	}
	pipeline.Workflow(ctx, comp, res)
}
