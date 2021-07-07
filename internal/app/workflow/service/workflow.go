package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/influxdata/cron"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action"
	"github.com/tsundata/assistant/internal/app/workflow/action/opcode"
	"github.com/tsundata/assistant/internal/app/workflow/repository"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/model"
	"strings"
	"time"
)

type Workflow struct {
	bus     *event.Bus
	rdb     *redis.Client
	logger  *logger.Logger
	message pb.MessageClient
	middle  pb.MiddleClient
	repo    repository.WorkflowRepository
}

func NewWorkflow(
	bus *event.Bus,
	rdb *redis.Client,
	repo repository.WorkflowRepository,
	message pb.MessageClient,
	middle pb.MiddleClient,
	logger *logger.Logger) *Workflow {
	return &Workflow{bus: bus, rdb: rdb, repo: repo, logger: logger, message: message, middle: middle}
}

func (s *Workflow) SyntaxCheck(_ context.Context, payload *pb.WorkflowRequest) (*pb.StateReply, error) {
	switch payload.Type {
	case model.MessageTypeAction:
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

func (s *Workflow) RunAction(_ context.Context, payload *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
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

	i := action.NewInterpreter(tree)
	i.SetClient(s.bus, s.rdb, s.message, s.middle, s.logger)
	_, err = i.Interpret()
	if err != nil {
		return nil, err
	}

	var result string
	if i.Ctx.Debug {
		result = fmt.Sprintf("Tracing\n-------\n %s", i.Stdout())
	}

	return &pb.WorkflowReply{
		Text: result,
	}, nil
}

func (s *Workflow) WebhookTrigger(ctx context.Context, payload *pb.TriggerRequest) (*pb.WorkflowReply, error) {
	trigger, err := s.repo.GetTriggerByFlag(payload.GetType(), payload.GetFlag())
	if err != nil {
		return nil, err
	}

	// Authorization
	if trigger.Secret != "" && payload.GetSecret() != trigger.Secret {
		return nil, errors.New("error secret")
	}

	// publish event
	err = s.bus.Publish(event.RunWorkflowSubject, model.Message{
		ID: trigger.MessageID,
	})
	if err != nil {
		return nil, err
	}

	return &pb.WorkflowReply{}, nil
}

func (s *Workflow) CronTrigger(ctx context.Context, _ *pb.TriggerRequest) (*pb.WorkflowReply, error) {
	triggers, err := s.repo.ListTriggersByType("cron")
	if err != nil {
		return nil, err
	}

	for _, trigger := range triggers {
		var lastTime time.Time
		key := fmt.Sprintf("workflow:cron:%d:time", trigger.MessageID)
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
			err = s.bus.Publish(event.RunWorkflowSubject, model.Message{ID: trigger.MessageID})
			if err != nil {
				return nil, err
			}
		}
	}

	return &pb.WorkflowReply{}, nil
}

func (s *Workflow) CreateTrigger(_ context.Context, payload *pb.TriggerRequest) (*pb.StateReply, error) {
	var trigger model.Trigger
	trigger.Type = payload.GetType()
	trigger.Kind = payload.GetKind()
	trigger.MessageID = int(payload.GetMessageId())

	switch payload.GetKind() {
	case model.MessageTypeAction:
		if payload.GetMessageText() == "" {
			return nil, errors.New("empty action")
		}
		p, err := action.NewParser(action.NewLexer([]rune(payload.GetMessageText())))
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
			_, err = s.repo.CreateTrigger(trigger)
			if err != nil {
				return nil, err
			}
		}

		if symbolTable.Webhook != nil {
			trigger.Type = "webhook"
			trigger.Flag = symbolTable.Webhook.Flag
			trigger.Secret = symbolTable.Webhook.Secret

			find, err := s.repo.GetTriggerByFlag(trigger.Type, trigger.Flag)
			if err != nil {
				return nil, err
			}

			if find.ID > 0 {
				return nil, errors.New("exist flag: " + trigger.Flag)
			}

			// store
			_, err = s.repo.CreateTrigger(trigger)
			if err != nil {
				return nil, err
			}
		}

		return &pb.StateReply{State: true}, nil
	default:
		return &pb.StateReply{State: false}, nil
	}
}

func (s *Workflow) DeleteTrigger(_ context.Context, payload *pb.TriggerRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteTriggerByMessageID(payload.GetMessageId())
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
