package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/influxdata/cron"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action"
	"github.com/tsundata/assistant/internal/app/workflow/action/opcode"
	"github.com/tsundata/assistant/internal/app/workflow/repository"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"strings"
	"time"
)

type Workflow struct {
	bus     event.Bus
	rdb     *redis.Client
	logger  log.Logger
	message pb.MessageSvcClient
	middle  pb.MiddleSvcClient
	repo    repository.WorkflowRepository
}

func NewWorkflow(
	bus event.Bus,
	rdb *redis.Client,
	repo repository.WorkflowRepository,
	message pb.MessageSvcClient,
	middle pb.MiddleSvcClient,
	logger log.Logger) *Workflow {
	return &Workflow{bus: bus, rdb: rdb, repo: repo, logger: logger, message: message, middle: middle}
}

func (s *Workflow) SyntaxCheck(_ context.Context, payload *pb.WorkflowRequest) (*pb.StateReply, error) {
	switch payload.Type {
	case enum.MessageTypeAction:
		if payload.GetText() == "" {
			return nil, errors.New("empty action")
		}
		p, err := action.NewParser(action.NewLexer([]rune(payload.GetText())))
		if err != nil {
			return &pb.StateReply{State: false}, err
		}
		tree, err := p.Parse()
		if err != nil {
			return &pb.StateReply{State: false}, err
		}

		symbolTable := action.NewSemanticAnalyzer()
		err = symbolTable.Visit(tree)
		if err != nil {
			return &pb.StateReply{State: false}, err
		}

		return &pb.StateReply{State: true}, nil
	default:
		return &pb.StateReply{State: false}, nil
	}
}

func (s *Workflow) RunAction(ctx context.Context, payload *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
	if payload.GetText() == "" {
		return nil, errors.New("empty action")
	}
	p, err := action.NewParser(action.NewLexer([]rune(payload.GetText())))
	if err != nil {
		return nil, err
	}
	tree, err := p.Parse()
	if err != nil {
		return nil, err
	}

	symbolTable := action.NewSemanticAnalyzer()
	err = symbolTable.Visit(tree)
	if err != nil {
		return nil, err
	}

	i := action.NewInterpreter(ctx, tree)
	i.SetComponent(s.bus, s.rdb, s.message, s.middle, s.logger)
	_, err = i.Interpret()
	if err != nil {
		return nil, err
	}

	var result string
	if i.Comp.Debug {
		result = fmt.Sprintf("Tracing\n-------\n %s", i.Stdout())
	}

	return &pb.WorkflowReply{
		Text: result,
	}, nil
}

func (s *Workflow) WebhookTrigger(ctx context.Context, payload *pb.TriggerRequest) (*pb.WorkflowReply, error) {
	trigger, err := s.repo.GetTriggerByFlag(ctx, payload.Trigger.GetType(), payload.Trigger.GetFlag())
	if err != nil {
		return nil, err
	}

	// Authorization
	if trigger.Secret != "" && payload.Trigger.GetSecret() != trigger.Secret {
		return nil, errors.New("error secret")
	}

	if trigger.MessageId <= 0 {
		return nil, errors.New("error trigger")
	}

	// publish event
	err = s.bus.Publish(ctx, enum.Workflow, event.WorkflowRunSubject, pb.Message{
		Id: trigger.MessageId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.WorkflowReply{}, nil
}

func (s *Workflow) CronTrigger(ctx context.Context, _ *pb.TriggerRequest) (*pb.WorkflowReply, error) {
	triggers, err := s.repo.ListTriggersByType(ctx, "cron")
	if err != nil {
		return nil, err
	}

	for _, trigger := range triggers {
		var lastTime time.Time
		key := fmt.Sprintf("workflow:cron:%d:time", trigger.MessageId)
		t := s.rdb.Get(ctx, key).Val()
		if t == "" {
			lastTime = time.Time{}
		} else {
			lastTime, err = time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
			if err != nil {
				return nil, err
			}
		}

		p, err := cron.ParseUTC(trigger.When)
		if err != nil {
			return nil, err
		}
		nextTime, err := p.Next(lastTime)
		if err != nil {
			return nil, err
		}

		now := time.Now()
		if nextTime.Before(now) {
			// time
			s.rdb.Set(ctx, key, now.Format("2006-01-02 15:04:05"), 0)

			// publish event
			err = s.bus.Publish(ctx, enum.Workflow, event.WorkflowRunSubject, pb.Message{Id: trigger.MessageId})
			if err != nil {
				return nil, err
			}
		}
	}

	return &pb.WorkflowReply{}, nil
}

func (s *Workflow) CreateTrigger(ctx context.Context, payload *pb.TriggerRequest) (*pb.StateReply, error) {
	var trigger pb.Trigger
	trigger.Type = payload.Trigger.GetType()
	trigger.Kind = payload.Trigger.GetKind()
	trigger.MessageId = payload.Trigger.GetMessageId()

	switch payload.Trigger.GetKind() {
	case enum.MessageTypeAction:
		if payload.Info.GetMessageText() == "" {
			return nil, errors.New("empty action")
		}
		p, err := action.NewParser(action.NewLexer([]rune(payload.Info.GetMessageText())))
		if err != nil {
			return nil, err
		}
		tree, err := p.Parse()
		if err != nil {
			return nil, err
		}

		symbolTable := action.NewSemanticAnalyzer()
		err = symbolTable.Visit(tree)
		if err != nil {
			return nil, err
		}

		if symbolTable.Cron == nil && symbolTable.Webhook == nil {
			return &pb.StateReply{State: false}, nil
		}

		if symbolTable.Cron != nil {
			trigger.Type = "cron"
			trigger.When = symbolTable.Cron.When

			// store
			_, err = s.repo.CreateTrigger(ctx, &trigger)
			if err != nil {
				return nil, err
			}
		}

		if symbolTable.Webhook != nil {
			trigger.Type = "webhook"
			trigger.Flag = symbolTable.Webhook.Flag
			trigger.Secret = symbolTable.Webhook.Secret

			find, err := s.repo.GetTriggerByFlag(ctx, trigger.Type, trigger.Flag)
			if err != nil {
				return nil, err
			}

			if find.Id > 0 {
				return nil, errors.New("exist flag: " + trigger.Flag)
			}

			// store
			_, err = s.repo.CreateTrigger(ctx, &trigger)
			if err != nil {
				return nil, err
			}
		}

		return &pb.StateReply{State: true}, nil
	default:
		return &pb.StateReply{State: false}, nil
	}
}

func (s *Workflow) DeleteTrigger(ctx context.Context, payload *pb.TriggerRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteTriggerByMessageID(ctx, payload.Trigger.GetMessageId())
	if err != nil {
		return &pb.StateReply{State: false}, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Workflow) ActionDoc(_ context.Context, payload *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
	var docs string
	if payload.GetText() == "" {
		docs = strings.Join(opcode.Docs(), "\n")
	} else {
		docs = opcode.Doc(payload.GetText())
	}
	return &pb.WorkflowReply{
		Text: docs,
	}, nil
}

func (s *Workflow) ListWebhook(ctx context.Context, _ *pb.WorkflowRequest) (*pb.WebhooksReply, error) {
	triggers, err := s.repo.ListTriggersByType(ctx, "webhook")
	if err != nil {
		return nil, err
	}
	var flags []string
	for _, item := range triggers {
		flags = append(flags, item.Flag)
	}
	return &pb.WebhooksReply{Flag: flags}, nil
}
