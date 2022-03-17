package repository

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/util"
	"gorm.io/gorm"
	"time"
)

type MessageRepository interface {
	GetByID(ctx context.Context, id int64) (pb.Message, error)
	GetByUUID(ctx context.Context, uuid string) (pb.Message, error)
	GetBySequence(ctx context.Context, userId, sequence int64) (pb.Message, error)
	GetLastByGroup(ctx context.Context, groupId int64) (pb.Message, error)
	ListByType(ctx context.Context, t string) ([]*pb.Message, error)
	List(ctx context.Context) ([]*pb.Message, error)
	ListByGroup(ctx context.Context, groupId int64, page, limit int) (int64, []*pb.Message, error)
	Create(ctx context.Context, message *pb.Message) (int64, error)
	Delete(ctx context.Context, id int64) error
	GetInbox(ctx context.Context, id int64) (pb.Inbox, error)
	ListInbox(ctx context.Context, userId int64, page, limit int) (int64, []*pb.Inbox, error)
	LastInbox(ctx context.Context, userId int64) (pb.Inbox, error)
	CreateInbox(ctx context.Context, inbox pb.Inbox) (int64, error)
	UpdateInboxStatus(ctx context.Context, id int64, status int) error
}

type MysqlMessageRepository struct {
	id     *global.ID
	locker *global.Locker
	db     *mysql.Conn
}

func NewMysqlMessageRepository(id *global.ID, locker *global.Locker, db *mysql.Conn) MessageRepository {
	return &MysqlMessageRepository{id: id, locker: locker, db: db}
}

func (r *MysqlMessageRepository) GetByID(ctx context.Context, id int64) (pb.Message, error) {
	var message pb.Message
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&message).Error
	if err != nil {
		return pb.Message{}, err
	}
	return message, nil
}

func (r *MysqlMessageRepository) GetByUUID(ctx context.Context, uuid string) (pb.Message, error) {
	var message pb.Message
	err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&message).Error
	if err != nil {
		return pb.Message{}, err
	}
	return message, nil
}

func (r *MysqlMessageRepository) GetBySequence(ctx context.Context, userId, sequence int64) (pb.Message, error) {
	var message pb.Message
	err := r.db.WithContext(ctx).Where("user_id = ? AND sequence = ?", userId, sequence).First(&message).Error
	if err != nil {
		return pb.Message{}, err
	}
	return message, nil
}

func (r *MysqlMessageRepository) GetLastByGroup(ctx context.Context, groupId int64) (pb.Message, error) {
	var message pb.Message
	err := r.db.WithContext(ctx).Where("group_id = ?", groupId).Order("created_at DESC, id DESC").Take(&message).Error
	if err != nil {
		return pb.Message{}, err
	}
	return message, nil
}

func (r *MysqlMessageRepository) ListByGroup(ctx context.Context, groupId int64, page, limit int) (int64, []*pb.Message, error) {
	var messages []*pb.Message
	var total int64
	err := r.db.WithContext(ctx).Model(&pb.Message{}).Where("group_id = ?", groupId).Count(&total).Error
	if err != nil {
		return 0, nil, err
	}
	err = r.db.WithContext(ctx).Where("group_id = ?", groupId).Order("created_at DESC, id DESC").
		Limit(limit).Offset((page - 1) * limit).
		Find(&messages).Error
	if err != nil {
		return 0, nil, err
	}
	return total, messages, nil
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
	err := r.db.WithContext(ctx).Where("type <> ?", enum.MessageTypeScript).Order("id DESC").Find(&messages).Error
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
	err = r.db.Where("user_id = ?", message.UserId).Order("sequence DESC").Take(&max).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	// sequence
	sequence := int64(0)
	if max.Sequence > 0 {
		sequence = max.Sequence
	}
	sequence += 1

	if message.Uuid == "" {
		message.Uuid = util.UUID()
	}
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

func (r *MysqlMessageRepository) GetInbox(ctx context.Context, id int64) (pb.Inbox, error) {
	var inbox pb.Inbox
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&inbox).Error
	if err != nil {
		return pb.Inbox{}, err
	}
	return inbox, nil
}

func (r *MysqlMessageRepository) ListInbox(ctx context.Context, userId int64, page, limit int) (int64, []*pb.Inbox, error) {
	var inbox []*pb.Inbox
	var total int64
	err := r.db.WithContext(ctx).Model(&pb.Inbox{}).Where("user_id = ?", userId).Count(&total).Error
	if err != nil {
		return 0, nil, err
	}
	err = r.db.WithContext(ctx).Where("user_id = ?", userId).Order("created_at DESC, id DESC").
		Limit(limit).Offset((page - 1) * limit).
		Find(&inbox).Error
	if err != nil {
		return 0, nil, err
	}
	return total, inbox, nil
}

func (r *MysqlMessageRepository) LastInbox(ctx context.Context, userId int64) (pb.Inbox, error) {
	var inbox pb.Inbox
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND status = ?", userId, enum.InboxCreate).
		Order("id ASC").Take(&inbox).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return pb.Inbox{}, err
	}
	return inbox, nil
}

func (r *MysqlMessageRepository) CreateInbox(ctx context.Context, inbox pb.Inbox) (int64, error) {
	l, err := r.locker.Acquire(fmt.Sprintf("message:index:create:%d", inbox.UserId))
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = l.Release()
	}()

	var max pb.Inbox
	err = r.db.Where("user_id = ?", inbox.UserId).Order("sequence DESC").Take(&max).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	// sequence
	sequence := int64(0)
	if max.Sequence > 0 {
		sequence = max.Sequence
	}
	sequence += 1

	if inbox.Uuid == "" {
		inbox.Uuid = util.UUID()
	}
	inbox.Id = r.id.Generate(ctx)
	inbox.Sequence = sequence
	inbox.CreatedAt = time.Now().Unix()
	inbox.UpdatedAt = time.Now().Unix()
	err = r.db.WithContext(ctx).Create(&inbox).Error
	if err != nil {
		return 0, err
	}
	return inbox.Id, nil
}

func (r *MysqlMessageRepository) UpdateInboxStatus(ctx context.Context, id int64, status int) error {
	return r.db.WithContext(ctx).Model(&pb.Inbox{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":     status,
		"updated_at": time.Now().Unix(),
	}).Error
}
