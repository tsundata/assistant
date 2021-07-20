package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/util"
)

type WorkflowRepository interface {
	GetTriggerByFlag(t, flag string) (pb.Trigger, error)
	ListTriggersByType(t string) ([]pb.Trigger, error)
	CreateTrigger(trigger pb.Trigger) (int64, error)
	DeleteTriggerByMessageID(messageID int64) error
}

type MysqlWorkflowRepository struct {
	logger log.Logger
	db     *sqlx.DB
}

func NewMysqlWorkflowRepository(logger log.Logger, db *sqlx.DB) WorkflowRepository {
	return &MysqlWorkflowRepository{logger: logger, db: db}
}

func (r *MysqlWorkflowRepository) GetTriggerByFlag(t, flag string) (pb.Trigger, error) {
	var trigger pb.Trigger
	err := r.db.Get(&trigger, "SELECT message_id, kind FROM `triggers` WHERE `type` = ? AND `flag` = ?", t, flag)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pb.Trigger{}, nil
		}
		return pb.Trigger{}, err
	}
	return trigger, nil
}

func (r *MysqlWorkflowRepository) ListTriggersByType(t string) ([]pb.Trigger, error) {
	var triggers []pb.Trigger
	err := r.db.Select(&triggers, "SELECT message_id, kind, `when` FROM `triggers` WHERE `type` = ?", t)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return triggers, nil
}

func (r *MysqlWorkflowRepository) CreateTrigger(trigger pb.Trigger) (int64, error) {
	res, err := r.db.NamedExec("INSERT INTO `triggers` (`type`, `kind`, `when`, `message_id`) VALUES (:type, :kind, :when, :message_id)", trigger)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MysqlWorkflowRepository) DeleteTriggerByMessageID(messageID int64) error {
	result, err := r.db.Exec("DELETE FROM triggers WHERE message_id = ?", messageID)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

type RqliteWorkflowRepository struct {
	logger log.Logger
	db     *rqlite.Conn
}

func NewRqliteWorkflowRepository(logger log.Logger, db *rqlite.Conn) WorkflowRepository {
	return &RqliteWorkflowRepository{logger: logger, db: db}
}

func (r *RqliteWorkflowRepository) GetTriggerByFlag(t, flag string) (pb.Trigger, error) {
	rows, err := r.db.QueryOne("SELECT message_id, kind FROM `triggers` WHERE `type` = '%s' AND `flag` = '%s'", t, flag)
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
	rows, err := r.db.QueryOne("SELECT message_id, kind, `when` FROM `triggers` WHERE `type` = '%s'", t)
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
	res, err := r.db.WriteOne("INSERT INTO `triggers` (`type`, `kind`, `when`, `message_id`, `created_at`) VALUES ('%s', '%s', '%s', '%d', '%s')", trigger.Type, trigger.Kind, trigger.When, trigger.MessageId, trigger.CreatedAt)
	if err != nil {
		return 0, err
	}
	return res.LastInsertID, nil
}

func (r *RqliteWorkflowRepository) DeleteTriggerByMessageID(messageID int64) error {
	_, err := r.db.WriteOne("DELETE FROM triggers WHERE message_id = %d", messageID)
	return err
}
