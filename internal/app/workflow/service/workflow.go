package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/influxdata/cron"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action"
	"github.com/tsundata/assistant/internal/app/workflow/action/opcode"
	"github.com/tsundata/assistant/internal/app/workflow/script"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"go.etcd.io/etcd/clientv3"
	"strconv"
	"strings"
	"time"
)

type Workflow struct {
	etcd       *clientv3.Client
	db         *sqlx.DB
	rdb        *redis.Client
	midClient  pb.MiddleClient
	msgClient  pb.MessageClient
	taskClient pb.TaskClient
}

func NewWorkflow(etcd *clientv3.Client, db *sqlx.DB, rdb *redis.Client, midClient pb.MiddleClient, msgClient pb.MessageClient, taskClient pb.TaskClient) *Workflow {
	return &Workflow{etcd: etcd, db: db, rdb: rdb, midClient: midClient, msgClient: msgClient, taskClient: taskClient}
}

func (s *Workflow) SyntaxCheck(_ context.Context, payload *pb.WorkflowRequest) (*pb.StateReply, error) {
	switch payload.Type {
	case model.MessageTypeAction:
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
	case model.MessageTypeScript:
		p, err := script.NewParser(script.NewLexer([]rune(payload.GetText())))
		if err != nil {
			return &pb.StateReply{State: false}, err
		}
		tree, err := p.Parse()
		if err != nil {
			return &pb.StateReply{State: false}, err
		}

		sa := script.NewSemanticAnalyzer()
		err = sa.Visit(tree)
		if err != nil {
			return &pb.StateReply{State: false}, err
		}

		return &pb.StateReply{State: true}, nil
	}

	return &pb.StateReply{State: false}, nil
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
	i.SetClient(s.rdb, s.midClient, s.msgClient, nil, s.taskClient)
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

		if symbolTable.Cron == nil && symbolTable.Webhook == nil {
			return &pb.StateReply{State: false}, nil
		}

		if symbolTable.Cron != nil {
			trigger.Type = "cron"
			trigger.When = symbolTable.Cron.When

			// store
			_, err := s.db.NamedExec("INSERT INTO `triggers` (`type`, `kind`, `when`, `message_id`, `time`) VALUES (:type, :kind, :when, :message_id, :time)", trigger)
			if err != nil {
				return nil, err
			}
		}

		if symbolTable.Webhook != nil {
			trigger.Type = "webhook"
			trigger.Flag = symbolTable.Webhook.Flag
			trigger.Secret = symbolTable.Webhook.Secret

			var find model.Trigger
			err = s.db.Get(&find, "SELECT id  FROM `triggers` WHERE `type` = ? AND `flag` = ?", trigger.Type, trigger.Flag)
			if err != nil && err != sql.ErrNoRows {
				return nil, err
			}

			if find.ID > 0 {
				return nil, errors.New("exist flag: " + trigger.Flag)
			}

			// store
			_, err = s.db.NamedExec("INSERT INTO `triggers` (`type`, `kind`, `flag`, `secret`, `message_id`, `time`) VALUES (:type, :kind, :flag, :secret, :message_id, :time)", trigger)
			if err != nil {
				return nil, err
			}
		}
	default:
		return &pb.StateReply{State: false}, nil
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Workflow) DeleteTrigger(_ context.Context, payload *pb.TriggerRequest) (*pb.StateReply, error) {
	result, err := s.db.Exec("DELETE FROM triggers WHERE message_id = ?", payload.MessageId)
	if err != nil {
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows > 0 {
		return &pb.StateReply{State: true}, nil
	}

	return &pb.StateReply{State: false}, nil
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
