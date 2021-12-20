package repository

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"gorm.io/gorm"
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
	GetGroupBySequence(ctx context.Context, userId, sequence int64) (*pb.Group, error)
	ListGroup(ctx context.Context, userId int64) ([]*pb.Group, error)
	CreateGroup(ctx context.Context, group *pb.Group) (int64, error)
	DeleteGroup(ctx context.Context, id int64) error
}

type MysqlMessageRepository struct {
	id     *global.ID
	locker *global.Locker
	db     *mysql.Conn
}

func NewMysqlMessageRepository(id *global.ID, locker *global.Locker, db *mysql.Conn) MessageRepository {
	return &MysqlMessageRepository{id: id, locker: locker, db: db}
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
	l, err := r.locker.Acquire(fmt.Sprintf("message:message:create:%d", message.UserId))
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = l.Release()
	}()

	var max pb.Message
	err = r.db.Where("user_id = ?", message.UserId).Order("sequence DESC").First(&max).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	// sequence
	sequence := int64(0)
	if max.Sequence > 0 {
		sequence = max.Sequence
	}
	sequence += 1

	message.Id = r.id.Generate(ctx)
	message.Sequence = sequence
	err = r.db.WithContext(ctx).Create(&message).Error
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

func (r *MysqlMessageRepository) GetGroupBySequence(ctx context.Context, userId, sequence int64) (*pb.Group, error) {
	var find pb.Group
	err := r.db.WithContext(ctx).Where("user_id = ? AND sequence = ?", userId, sequence).First(&find).Error
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
	l, err := r.locker.Acquire(fmt.Sprintf("message:group:create:%d", group.UserId))
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = l.Release()
	}()

	var max pb.Group
	err = r.db.Where("user_id = ?", group.UserId).Order("sequence DESC").First(&max).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	// sequence
	sequence := int64(0)
	if max.Sequence > 0 {
		sequence = max.Sequence
	}
	sequence += 1

	group.Id = r.id.Generate(ctx)
	group.Sequence = sequence
	err = r.db.WithContext(ctx).Create(&group).Error
	if err != nil {
		return 0, err
	}
	return group.Id, nil
}

func (r *MysqlMessageRepository) DeleteGroup(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.Group{}).Error
}
