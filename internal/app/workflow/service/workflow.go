package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/influxdata/cron"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action"
	"github.com/tsundata/assistant/internal/app/workflow/script"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"go.etcd.io/etcd/clientv3"
	"strconv"
	"time"
)

type Workflow struct {
	etcd       *clientv3.Client
	db         *sqlx.DB
	midClient  pb.MiddleClient
	msgClient  pb.MessageClient
	taskClient pb.TaskClient
}

func NewWorkflow(etcd *clientv3.Client, db *sqlx.DB, midClient pb.MiddleClient, msgClient pb.MessageClient, taskClient pb.TaskClient) *Workflow {
	return &Workflow{etcd: etcd, db: db, midClient: midClient, msgClient: msgClient, taskClient: taskClient}
}

func (s *Workflow) RunScript(_ context.Context, payload *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
	p, err := script.NewParser(script.NewLexer([]rune(payload.GetText())))
	if err != nil {
		return nil, err
	}
	tree, err := p.Parse()
	if err != nil {
		return nil, err
	}

	sa := script.NewSemanticAnalyzer()
	err = sa.Visit(tree)
	if err != nil {
		return nil, err
	}

	i := script.NewInterpreter(tree)
	i.SetClient(s.midClient)
	_, err = i.Interpret()
	if err != nil {
		return nil, err
	}

	return &pb.WorkflowReply{
		Text: fmt.Sprintf("Tracing\n-------\n %s", i.Stdout()),
	}, nil
}

func (s *Workflow) RunAction(_ context.Context, payload *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
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
	i.SetClient(s.midClient, s.msgClient)
	_, err = i.Interpret()
	if err != nil {
		return nil, err
	}

	return &pb.WorkflowReply{
		Text: fmt.Sprintf("Tracing\n-------\n %s", i.Stdout()),
	}, nil
}

func (s *Workflow) WebhookTrigger(ctx context.Context, payload *pb.TriggerRequest) (*pb.WorkflowReply, error) {
	var trigger model.Trigger
	err := s.db.Get(&trigger, "SELECT * FROM `triggers` WHERE `type` = ? AND `flag` = ?", payload.Type, payload.Flag)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// push task
	j, err := json.Marshal(map[string]string{
		"type": trigger.Kind,
		"id":   strconv.Itoa(trigger.MessageID),
	})
	if err != nil {
		return nil, err
	}
	_, err = s.taskClient.Send(ctx, &pb.JobRequest{Name: "run", Args: utils.ByteToString(j)})
	if err != nil {
		return nil, err
	}

	return &pb.WorkflowReply{
		Text: "",
	}, nil
}

func (s *Workflow) CronTrigger(ctx context.Context, _ *pb.TriggerRequest) (*pb.WorkflowReply, error) {
	var triggers []model.Trigger
	err := s.db.Select(&triggers, "SELECT * FROM `triggers` WHERE `type` = ?", "cron")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	for _, trigger := range triggers {
		p, err := cron.ParseUTC(trigger.When)
		if err != nil {
			return nil, err
		}
		nextTime, err := p.Next(time.Now())
		if err != nil {
			return nil, err
		}

		if nextTime.Format("2006-01-02 15:04") == time.Now().Format("2006-01-02 15:04") {
			// push task
			j, err := json.Marshal(map[string]string{
				"type": trigger.Kind,
				"id":   strconv.Itoa(trigger.MessageID),
			})
			if err != nil {
				return nil, err
			}
			_, err = s.taskClient.Send(ctx, &pb.JobRequest{Name: "run", Args: utils.ByteToString(j)})
			if err != nil {
				return nil, err
			}
		}
	}

	return &pb.WorkflowReply{
		Text: "",
	}, nil
}

func (s *Workflow) CreateTrigger(_ context.Context, payload *pb.TriggerRequest) (*pb.StateReply, error) {
	var trigger model.Trigger
	trigger.Type = payload.GetType()
	trigger.Kind = payload.GetKind()
	trigger.MessageID = int(payload.GetMessageId())
	trigger.Time = time.Now()

	switch payload.GetKind() {
	case model.MessageTypeAction:
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

		if symbolTable.Webhook == nil {
			return nil, nil
		} else {
			trigger.Type = "webhook"
			trigger.Flag = symbolTable.Webhook.Flag
			trigger.Secret = symbolTable.Webhook.Secret
		}

		if symbolTable.Cron == nil {
			return nil, nil
		} else {
			trigger.Type = "cron"
			trigger.When = symbolTable.Cron.When
		}
	case model.MessageTypeScript:
		// TODO
		return nil, nil
	default:
		return nil, nil
	}

	// store
	res, err := s.db.NamedExec("INSERT INTO `triggers` (`type`, `kind`, `flag`, `secret`, `message_id`, `time`) VALUES (:type, :kind, :flag, :secret, :message_id, :time)", trigger)
	if err != nil {
		return nil, err
	}
	_, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{
		State: true,
	}, nil
}
