package repository

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
)

type ChatbotRepository interface {
	GetByID(ctx context.Context, id int64) (*pb.Bot, error)
	GetByUUID(ctx context.Context, uuid string) (*pb.Bot, error)
	GetByIdentifier(ctx context.Context, uuid string) (*pb.Bot, error)
	List(ctx context.Context, ) ([]*pb.Bot, error)
	Create(ctx context.Context, message *pb.Bot) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type MysqlChatbotRepository struct {
	id *global.ID
	db *mysql.Conn
}

func NewMysqlChatbotRepository(id *global.ID, db *mysql.Conn) ChatbotRepository {
	return &MysqlChatbotRepository{id: id, db: db}
}

func (r *MysqlChatbotRepository) GetByID(ctx context.Context, id int64) (*pb.Bot, error) {
	var bot pb.Bot
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&bot).Error
	if err != nil {
		return nil, err
	}
	return &bot, nil
}

func (r *MysqlChatbotRepository) GetByUUID(ctx context.Context, uuid string) (*pb.Bot, error) {
	var bot pb.Bot
	err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&bot).Error
	if err != nil {
		return nil, err
	}
	return &bot, nil
}

func (r *MysqlChatbotRepository) GetByIdentifier(ctx context.Context, identifier string) (*pb.Bot, error) {
	var bot pb.Bot
	err := r.db.WithContext(ctx).Where("identifier = ?", identifier).First(&bot).Error
	if err != nil {
		return nil, err
	}
	return &bot, nil
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
