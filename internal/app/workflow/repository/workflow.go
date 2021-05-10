package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/model"
)

type WorkflowRepository interface {
	GetTriggerByFlag(t, flag string) (model.Trigger, error)
	ListTriggersByType(t string) ([]model.Trigger, error)
	CreateTrigger(trigger model.Trigger) (int64, error)
	DeleteTriggerByMessageID(messageID int64) error
}

type MysqlWorkflowRepository struct {
	logger *logger.Logger
	db     *sqlx.DB
}

func NewMysqlWorkflowRepository(logger *logger.Logger, db *sqlx.DB) WorkflowRepository {
	return &MysqlWorkflowRepository{logger: logger, db: db}
}

func (r *MysqlWorkflowRepository) GetTriggerByFlag(t, flag string) (model.Trigger, error) {
	var trigger model.Trigger
	err := r.db.Get(&trigger, "SELECT message_id, kind FROM `triggers` WHERE `type` = ? AND `flag` = ?", t, flag)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Trigger{}, nil
		}
		return model.Trigger{}, err
	}
	return trigger, nil
}

func (r *MysqlWorkflowRepository) ListTriggersByType(t string) ([]model.Trigger, error) {
	var triggers []model.Trigger
	err := r.db.Select(&triggers, "SELECT message_id, kind, `when` FROM `triggers` WHERE `type` = ?", t)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return triggers, nil
}

func (r *MysqlWorkflowRepository) CreateTrigger(trigger model.Trigger) (int64, error) {
	res, err := r.db.NamedExec("INSERT INTO `triggers` (`type`, `kind`, `when`, `message_id`, `time`) VALUES (:type, :kind, :when, :message_id, :time)", trigger)
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
