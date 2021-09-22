package repository

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/util"
)

type MessageRepository interface {
	GetByID(ctx context.Context, id int64) (*pb.Message, error)
	GetByUUID(ctx context.Context, uuid string) (*pb.Message, error)
	ListByType(ctx context.Context, t string) ([]*pb.Message, error)
	List(ctx context.Context, ) ([]*pb.Message, error)
	Create(ctx context.Context, message *pb.Message) (int64, error)
	Delete(ctx context.Context, id int64) error
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
	err := r.db.WithContext(ctx).Order("id DESC").Find(&messages).Error
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

type RqliteMessageRepository struct {
	db *rqlite.Conn
}

func NewRqliteMessageRepository(db *rqlite.Conn) *RqliteMessageRepository {
	return &RqliteMessageRepository{db: db}
}

func (r *RqliteMessageRepository) GetByID(id int64) (pb.Message, error) {
	rows, err := r.db.QueryOne("SELECT id, uuid, text, `type`, `created_at` FROM `messages` WHERE `id` = '%d' LIMIT 1", id)
	if err != nil {
		return pb.Message{}, nil
	}

	var message pb.Message
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.Message{}, err
		}
		util.Inject(&message, m)
	}

	return message, nil
}

func (r *RqliteMessageRepository) GetByUUID(uuid string) (pb.Message, error) {
	rows, err := r.db.QueryOne("SELECT id, uuid, text, `type`, `created_at` FROM `messages` WHERE `uuid` = '%s' LIMIT 1", uuid)
	if err != nil {
		return pb.Message{}, err
	}

	var message pb.Message
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.Message{}, err
		}
		util.Inject(&message, m)
	}

	return message, nil
}

func (r *RqliteMessageRepository) ListByType(t string) ([]pb.Message, error) {
	rows, err := r.db.QueryOne("SELECT id, uuid, text, `type`, `created_at` FROM `messages` WHERE `type` = '%s' ORDER BY `id` DESC", t)
	if err != nil {
		return nil, nil
	}

	var messages []pb.Message
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return nil, err
		}
		var item pb.Message
		util.Inject(&item, m)
		messages = append(messages, item)
	}

	return messages, nil
}

func (r *RqliteMessageRepository) List() ([]pb.Message, error) {
	rows, err := r.db.QueryOne("SELECT uuid, text, `type`, `created_at` FROM `messages` WHERE `type` <> '%s' ORDER BY `id` DESC", enum.MessageTypeAction)
	if err != nil {
		return nil, err
	}

	var messages []pb.Message
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return nil, err
		}
		var item pb.Message
		util.Inject(&item, m)
		messages = append(messages, item)
	}

	return messages, nil
}

func (r *RqliteMessageRepository) Create(message pb.Message) (int64, error) {
	res, err := r.db.WriteOne("INSERT INTO `messages` (`uuid`, `type`, `text`) VALUES ('%s', '%s', '%s')", message.Uuid, message.Type, message.Text)
	if err != nil {
		return 0, err
	}

	return res.LastInsertID, nil
}

func (r *RqliteMessageRepository) Delete(id int64) error {
	_, err := r.db.WriteOne("DELETE FROM `messages` WHERE `id` = '%d'", id)
	return err
}
