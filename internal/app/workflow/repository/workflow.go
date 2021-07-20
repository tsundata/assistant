package repository

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/util"
)

type WorkflowRepository interface {
	GetTriggerByFlag(t, flag string) (pb.Trigger, error)
	ListTriggersByType(t string) ([]pb.Trigger, error)
	CreateTrigger(trigger pb.Trigger) (int64, error)
	DeleteTriggerByMessageID(messageID int64) error
}

type RqliteWorkflowRepository struct {
	db *rqlite.Conn
}

func NewRqliteWorkflowRepository(db *rqlite.Conn) WorkflowRepository {
	return &RqliteWorkflowRepository{db: db}
}

func (r *RqliteWorkflowRepository) GetTriggerByFlag(t, flag string) (pb.Trigger, error) {
	rows, err := r.db.QueryOne("SELECT * FROM `triggers` WHERE `type` = '%s' AND `flag` = '%s'", t, flag)
	if err != nil {
		return pb.Trigger{}, err
	}

	var trigger pb.Trigger
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.Trigger{}, err
		}
		util.Inject(&trigger, m)
	}

	return trigger, nil
}

func (r *RqliteWorkflowRepository) ListTriggersByType(t string) ([]pb.Trigger, error) {
	rows, err := r.db.QueryOne("SELECT * FROM `triggers` WHERE `type` = '%s'", t)
	if err != nil {
		return nil, err
	}

	var triggers []pb.Trigger
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return nil, err
		}
		var item pb.Trigger
		util.Inject(&item, m)
		triggers = append(triggers, item)
	}

	return triggers, nil
}

func (r *RqliteWorkflowRepository) CreateTrigger(trigger pb.Trigger) (int64, error) {
	trigger.CreatedAt = util.Now()
	res, err := r.db.WriteOne("INSERT INTO `triggers` (`type`, `kind`, `flag`, `secret`, `when`, `message_id`, `created_at`) VALUES ('%s', '%s', '%s', '%s', '%s', '%d', '%s')",
		trigger.Type, trigger.Kind, trigger.Flag, trigger.Secret, trigger.When, trigger.MessageId, trigger.CreatedAt)
	if err != nil {
		return 0, err
	}
	return res.LastInsertID, nil
}

func (r *RqliteWorkflowRepository) DeleteTriggerByMessageID(messageID int64) error {
	_, err := r.db.WriteOne("DELETE FROM triggers WHERE message_id = %d", messageID)
	return err
}
