package repository

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"gorm.io/gorm"
)

type ChatbotRepository interface {
	GetByID(ctx context.Context, id int64) (pb.Bot, error)
	GetByUUID(ctx context.Context, uuid string) (pb.Bot, error)
	GetByIdentifier(ctx context.Context, uuid string) (pb.Bot, error)
	List(ctx context.Context, ) ([]*pb.Bot, error)
	Create(ctx context.Context, message *pb.Bot) (int64, error)
	Delete(ctx context.Context, id int64) error
	GetGroup(ctx context.Context, id int64) (pb.Group, error)
	GetGroupByUUID(ctx context.Context, uuid string) (pb.Group, error)
	GetGroupBySequence(ctx context.Context, userId, sequence int64) (pb.Group, error)
	ListGroup(ctx context.Context, userId int64) ([]*pb.Group, error)
	CreateGroup(ctx context.Context, group *pb.Group) (int64, error)
	DeleteGroup(ctx context.Context, id int64) error
}

type MysqlChatbotRepository struct {
	id     *global.ID
	locker *global.Locker
	db     *mysql.Conn
}

func NewMysqlChatbotRepository(id *global.ID, locker *global.Locker, db *mysql.Conn) ChatbotRepository {
	return &MysqlChatbotRepository{id: id, locker: locker, db: db}
}

func (r *MysqlChatbotRepository) GetByID(ctx context.Context, id int64) (pb.Bot, error) {
	var bot pb.Bot
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&bot).Error
	if err != nil {
		return pb.Bot{}, err
	}
	return bot, nil
}

func (r *MysqlChatbotRepository) GetByUUID(ctx context.Context, uuid string) (pb.Bot, error) {
	var bot pb.Bot
	err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&bot).Error
	if err != nil {
		return pb.Bot{}, err
	}
	return bot, nil
}

func (r *MysqlChatbotRepository) GetByIdentifier(ctx context.Context, identifier string) (pb.Bot, error) {
	var bot pb.Bot
	err := r.db.WithContext(ctx).Where("identifier = ?", identifier).First(&bot).Error
	if err != nil {
		return pb.Bot{}, err
	}
	return bot, nil
}

func (r *MysqlChatbotRepository) List(ctx context.Context) ([]*pb.Bot, error) {
	var bots []*pb.Bot
	err := r.db.WithContext(ctx).Order("id DESC").Find(&bots).Error
	if err != nil {
		return nil, err
	}
	return bots, nil
}

func (r *MysqlChatbotRepository) Create(ctx context.Context, bot *pb.Bot) (int64, error) {
	bot.Id = r.id.Generate(ctx)
	err := r.db.WithContext(ctx).Create(&bot).Error
	if err != nil {
		return 0, err
	}
	return bot.Id, nil
}

func (r *MysqlChatbotRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.Bot{}).Error
}

func (r *MysqlChatbotRepository) GetGroup(ctx context.Context, id int64) (pb.Group, error) {
	var find pb.Group
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&find).Error
	if err != nil {
		return pb.Group{}, err
	}
	return find, nil
}

func (r *MysqlChatbotRepository) GetGroupByUUID(ctx context.Context, uuid string) (pb.Group, error) {
	var find pb.Group
	err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&find).Error
	if err != nil {
		return pb.Group{}, err
	}
	return find, nil
}

func (r *MysqlChatbotRepository) GetGroupBySequence(ctx context.Context, userId, sequence int64) (pb.Group, error) {
	var find pb.Group
	err := r.db.WithContext(ctx).Where("user_id = ? AND sequence = ?", userId, sequence).First(&find).Error
	if err != nil {
		return pb.Group{}, err
	}
	return find, nil
}

func (r *MysqlChatbotRepository) ListGroup(ctx context.Context, userId int64) ([]*pb.Group, error) {
	var list []*pb.Group
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Order("id DESC").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *MysqlChatbotRepository) CreateGroup(ctx context.Context, group *pb.Group) (int64, error) {
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

func (r *MysqlChatbotRepository) DeleteGroup(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.Group{}).Error
}
