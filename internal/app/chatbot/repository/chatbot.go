package repository

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ChatbotRepository interface {
	GetByID(ctx context.Context, id int64) (pb.Bot, error)
	GetByUUID(ctx context.Context, uuid string) (pb.Bot, error)
	GetByIdentifier(ctx context.Context, uuid string) (pb.Bot, error)
	List(ctx context.Context, ) ([]*pb.Bot, error)
	Create(ctx context.Context, bot *pb.Bot) (int64, error)
	Delete(ctx context.Context, id int64) error
	ListGroupBot(ctx context.Context, groupId int64) ([]*pb.Bot, error)
	CreateGroupBot(ctx context.Context, groupId int64, bot *pb.Bot) error
	DeleteGroupBot(ctx context.Context, groupId, botId int64) error
	GetGroup(ctx context.Context, id int64) (pb.Group, error)
	GetGroupByUUID(ctx context.Context, uuid string) (pb.Group, error)
	GetGroupBySequence(ctx context.Context, userId, sequence int64) (pb.Group, error)
	ListGroup(ctx context.Context, userId int64) ([]*pb.Group, error)
	CreateGroup(ctx context.Context, group *pb.Group) (int64, error)
	DeleteGroup(ctx context.Context, id int64) error
	UpdateGroup(ctx context.Context, group *pb.Group) error
	UpdateGroupSetting(ctx context.Context, groupId int64, kvs []*pb.KV) error
	UpdateGroupBotSetting(ctx context.Context, groupId, botId int64, kvs []*pb.KV) error
	ListGroupTag(ctx context.Context, groupId int64) ([]*pb.GroupTag, error)
	CreateGroupTag(ctx context.Context, tag *pb.GroupTag) (int64, error)
	DeleteGroupTag(ctx context.Context, id int64) error
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

func (r *MysqlChatbotRepository) ListGroupBot(ctx context.Context, groupId int64) ([]*pb.Bot, error) {
	var bots []*pb.Bot
	var groupBots []*pb.GroupBot
	err := r.db.WithContext(ctx).Where("group_id = ?", groupId).Find(&groupBots).Error
	if err != nil {
		return nil, err
	}
	var botId []int64
	for _, item := range groupBots {
		botId = append(botId, item.BotId)
	}
	if len(botId) > 0 {
		err = r.db.WithContext(ctx).Where("id IN ?", botId).Find(&bots).Error
		if err != nil {
			return nil, err
		}
	}

	return bots, nil
}

func (r *MysqlChatbotRepository) CreateGroupBot(ctx context.Context, groupId int64, bot *pb.Bot) error {
	groupBot := pb.GroupBot{}
	groupBot.Id = r.id.Generate(ctx)
	groupBot.GroupId = groupId
	groupBot.BotId = bot.Id
	return r.db.WithContext(ctx).Create(&groupBot).Error
}

func (r *MysqlChatbotRepository) DeleteGroupBot(ctx context.Context, groupId, botId int64) error {
	return r.db.WithContext(ctx).Where("group_id = ? AND bot_id = ?", groupId, botId).Delete(&pb.GroupBot{}).Error
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

func (r *MysqlChatbotRepository) UpdateGroup(ctx context.Context, group *pb.Group) error {
	return r.db.WithContext(ctx).Where("id = ?", group.Id).Update("name", group.Name).Error
}

func (r *MysqlChatbotRepository) UpdateGroupSetting(ctx context.Context, groupId int64, kvs []*pb.KV) error {
	for _, item := range kvs {
		groupSetting := pb.GroupSetting{
			GroupId: groupId,
			Key:     item.Key,
			Value:   item.Value,
		}
		err := r.db.WithContext(ctx).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "group_id"}, {Name: "key"}},
			DoUpdates: clause.AssignmentColumns([]string{"value"}),
		}).Create(&groupSetting).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *MysqlChatbotRepository) UpdateGroupBotSetting(ctx context.Context, groupId, botId int64, kvs []*pb.KV) error {
	for _, item := range kvs {
		groupBotSetting := pb.GroupBotSetting{
			GroupId: groupId,
			BotId:   botId,
			Key:     item.Key,
			Value:   item.Value,
		}
		err := r.db.WithContext(ctx).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "group_id"}, {Name: "bot_id"}, {Name: "key"}},
			DoUpdates: clause.AssignmentColumns([]string{"value"}),
		}).Create(&groupBotSetting).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *MysqlChatbotRepository) ListGroupTag(ctx context.Context, groupId int64) ([]*pb.GroupTag, error) {
	var list []*pb.GroupTag
	err := r.db.WithContext(ctx).Where("group_id = ?", groupId).Order("id DESC").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *MysqlChatbotRepository) CreateGroupTag(ctx context.Context, tag *pb.GroupTag) (int64, error) {
	tag.Id = r.id.Generate(ctx)
	err := r.db.WithContext(ctx).Create(&tag).Error
	if err != nil {
		return 0, err
	}
	return tag.Id, nil
}

func (r *MysqlChatbotRepository) DeleteGroupTag(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.GroupTag{}).Error
}
