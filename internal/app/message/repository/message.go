package repository

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
)

type MessageRepository interface {
	GetByID(ctx context.Context, id int64) (*pb.Message, error)
	GetByUUID(ctx context.Context, uuid string) (*pb.Message, error)
	ListByType(ctx context.Context, t string) ([]*pb.Message, error)
	List(ctx context.Context, ) ([]*pb.Message, error)
	Create(ctx context.Context, message *pb.Message) (int64, error)
	Delete(ctx context.Context, id int64) error
	GetGroup(ctx context.Context, id int64) (*pb.Group, error)
	GetGroupByUUID(ctx context.Context, uuid string) (*pb.Group, error)
	ListGroup(ctx context.Context, userId int64) ([]*pb.Group, error)
	CreateGroup(ctx context.Context, group *pb.Group) (int64, error)
	DeleteGroup(ctx context.Context, id int64) error
}

type MysqlMessageRepository struct {
	db *mysql.Conn
}

func NewMysqlMessageRepository(db *mysql.Conn) MessageRepository {
	return &MysqlMessageRepository{db: db}
}

func (r *MysqlMessageRepository) GetByID(ctx context.Context, id int64) (*pb.Message, error) {
	var message pb.Message
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&message).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *MysqlMessageRepository) GetByUUID(ctx context.Context, uuid string) (*pb.Message, error) {
	var message pb.Message
	err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&message).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *MysqlMessageRepository) ListByType(ctx context.Context, t string) ([]*pb.Message, error) {
	var messages []*pb.Message
	err := r.db.WithContext(ctx).Where("type = ?", t).Order("id DESC").Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *MysqlMessageRepository) List(ctx context.Context) ([]*pb.Message, error) {
	var messages []*pb.Message
	err := r.db.WithContext(ctx).Where("type <> ?", enum.MessageTypeAction).Order("id DESC").Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *MysqlMessageRepository) Create(ctx context.Context, message *pb.Message) (int64, error) {
	err := r.db.WithContext(ctx).Create(&message).Error
	if err != nil {
		return 0, err
	}
	return message.Id, nil
}

func (r *MysqlMessageRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.Message{}).Error
}

func (r *MysqlMessageRepository) GetGroup(ctx context.Context, id int64) (*pb.Group, error) {
	var find pb.Group
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&find).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}

func (r *MysqlMessageRepository) GetGroupByUUID(ctx context.Context, uuid string) (*pb.Group, error) {
	var find pb.Group
	err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&find).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}

func (r *MysqlMessageRepository) ListGroup(ctx context.Context, userId int64) ([]*pb.Group, error) {
	var list []*pb.Group
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Order("id DESC").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *MysqlMessageRepository) CreateGroup(ctx context.Context, group *pb.Group) (int64, error) {
	err := r.db.WithContext(ctx).Create(&group).Error
	if err != nil {
		return 0, err
	}
	return group.Id, nil
}

func (r *MysqlMessageRepository) DeleteGroup(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.Group{}).Error
}
