package repository

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
)

type WorkflowRepository interface {
	GetTriggerByFlag(ctx context.Context, t, flag string) (pb.Trigger, error)
	ListTriggersByType(ctx context.Context, t string) ([]*pb.Trigger, error)
	CreateTrigger(ctx context.Context, trigger *pb.Trigger) (int64, error)
	DeleteTriggerByMessageID(ctx context.Context, messageID int64) error
}

type MysqlWorkflowRepository struct {
	id *global.ID
	db *mysql.Conn
}

func NewMysqlWorkflowRepository(id *global.ID, db *mysql.Conn) WorkflowRepository {
	return &MysqlWorkflowRepository{id: id, db: db}
}

func (r *MysqlWorkflowRepository) GetTriggerByFlag(ctx context.Context, t, flag string) (pb.Trigger, error) {
	var trigger pb.Trigger
	err := r.db.WithContext(ctx).
		Where("type = ?", t).
		Where("flag = ?", flag).
		First(&trigger).Error
	if err != nil {
		return pb.Trigger{}, err
	}
	return trigger, nil
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
	trigger.Id = r.id.Generate(ctx)
	err := r.db.WithContext(ctx).Create(&trigger).Error
	if err != nil {
		return 0, err
	}
	return trigger.Id, nil
}

func (r *MysqlWorkflowRepository) DeleteTriggerByMessageID(ctx context.Context, messageID int64) error {
	return r.db.WithContext(ctx).Where("message_id = ?", messageID).Delete(&pb.Trigger{}).Error
}
