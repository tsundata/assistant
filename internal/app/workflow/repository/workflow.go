package repository

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/util"
)

type WorkflowRepository interface {
	GetTriggerByFlag(ctx context.Context, t, flag string) (*pb.Trigger, error)
	ListTriggersByType(ctx context.Context, t string) ([]*pb.Trigger, error)
	CreateTrigger(ctx context.Context, trigger *pb.Trigger) (int64, error)
	DeleteTriggerByMessageID(ctx context.Context, messageID int64) error
}

type MysqlWorkflowRepository struct {
	db *mysql.Conn
}

func NewMysqlWorkflowRepository(db *mysql.Conn) WorkflowRepository {
	return &MysqlWorkflowRepository{db: db}
}

func (r *MysqlWorkflowRepository) GetTriggerByFlag(ctx context.Context, t, flag string) (*pb.Trigger, error) {
	var trigger pb.Trigger
	err := r.db.WithContext(ctx).
		Where("type = ?", t).
		Where("flag = ?", flag).
		First(&trigger).Error
	if err != nil {
		return nil, err
	}
	return &trigger, nil
}

func (r *MysqlWorkflowRepository) ListTriggersByType(ctx context.Context, t string) ([]*pb.Trigger, error) {
	var triggers []*pb.Trigger
	err := r.db.WithContext(ctx).Where("type = ?", t).Find(&triggers).Error
	if err != nil {
		return nil, err
	}
	return triggers, nil
}

func (r *MysqlWorkflowRepository) CreateTrigger(ctx context.Context, trigger *pb.Trigger) (int64, error) {
	err := r.db.WithContext(ctx).Create(&trigger).Error
	if err != nil {
		return 0, err
	}
	return trigger.Id, nil
}

func (r *MysqlWorkflowRepository) DeleteTriggerByMessageID(ctx context.Context, messageID int64) error {
	return r.db.WithContext(ctx).Where("message_id = ?", messageID).Delete(&pb.Trigger{}).Error
}

type RqliteWorkflowRepository struct {
	db *rqlite.Conn
}

func NewRqliteWorkflowRepository(db *rqlite.Conn) *RqliteWorkflowRepository {
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
	//trigger.CreatedAt = util.Now()
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
