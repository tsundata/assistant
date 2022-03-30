package repository

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"gorm.io/gorm"
	"time"
)

type ChatbotRepository interface {
	GetByID(ctx context.Context, id int64) (pb.Bot, error)
	GetByUUID(ctx context.Context, uuid string) (pb.Bot, error)
	GetByIdentifier(ctx context.Context, identifier string) (pb.Bot, error)
	GetGroupBot(ctx context.Context, groupId, botId int64) (pb.Bot, error)
	List(ctx context.Context) ([]*pb.Bot, error)
	GetBotsByIds(ctx context.Context, id []int64) ([]*pb.Bot, error)
	GetBotsByGroupUuid(ctx context.Context, uuid string) ([]*pb.Bot, error)
	GetBotsByUser(ctx context.Context, userId int64) ([]*pb.Bot, error)
	GetBotsByText(ctx context.Context, text []string) (map[string]*pb.Bot, error)
	Create(ctx context.Context, bot *pb.Bot) (int64, error)
	Update(ctx context.Context, bot *pb.Bot) error
	Delete(ctx context.Context, id int64) error
	ListGroupBot(ctx context.Context, groupId int64) ([]*pb.Bot, error)
	CreateGroupBot(ctx context.Context, groupId int64, bot *pb.Bot) error
	DeleteGroupBot(ctx context.Context, groupId, botId int64) error
	GetGroup(ctx context.Context, id int64) (pb.Group, error)
	GetGroupByUUID(ctx context.Context, uuid string) (pb.Group, error)
	GetGroupBySequence(ctx context.Context, userId, sequence int64) (pb.Group, error)
	GetGroupByName(ctx context.Context, userId int64, name string) (pb.Group, error)
	TouchGroupUpdatedAt(ctx context.Context, id int64) error
	ListGroup(ctx context.Context, userId int64) ([]*pb.Group, error)
	CreateGroup(ctx context.Context, group *pb.Group) (int64, error)
	DeleteGroup(ctx context.Context, id int64) error
	UpdateGroup(ctx context.Context, group *pb.Group) error
	UpdateGroupSetting(ctx context.Context, groupId int64, kvs []*pb.KV) error
	UpdateGroupBotSetting(ctx context.Context, groupId, botId int64, kvs []*pb.KV) error
	GetGroupSetting(ctx context.Context, groupId int64) ([]*pb.KV, error)
	GetGroupSettingByUuid(ctx context.Context, groupUuid string) ([]*pb.KV, error)
	GetGroupBotSetting(ctx context.Context, groupId, botId int64) ([]*pb.KV, error)
	GetGroupBotSettingByUuid(ctx context.Context, groupUuid, botUuid string) ([]*pb.KV, error)
	GetGroupBotSettingByGroup(ctx context.Context, groupId int64) (map[int64][]*pb.KV, error)
	ListGroupTag(ctx context.Context, groupId int64) ([]*pb.GroupTag, error)
	CreateGroupTag(ctx context.Context, tag *pb.GroupTag) (int64, error)
	DeleteGroupTag(ctx context.Context, id int64) error
	GetTriggerByFlag(ctx context.Context, t, flag string) (pb.Trigger, error)
	ListTriggersByType(ctx context.Context, t string) ([]*pb.Trigger, error)
	CreateTrigger(ctx context.Context, trigger *pb.Trigger) (int64, error)
	DeleteTriggerByMessageID(ctx context.Context, messageID int64) error
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

func (r *MysqlChatbotRepository) GetGroupBot(ctx context.Context, groupId, botId int64) (pb.Bot, error) {
	var bot pb.Bot
	err := r.db.WithContext(ctx).
		Select("bots.id, bots.uuid as uuid, bots.name, bots.identifier, bots.avatar").
		Where("group_bots.group_id = ? AND group_bots.bot_id = ?", groupId, botId).
		Joins("LEFT JOIN group_bots ON group_bots.bot_id = bots.id").
		First(&bot).Error
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

func (r *MysqlChatbotRepository) GetBotsByIds(ctx context.Context, id []int64) ([]*pb.Bot, error) {
	var bots []*pb.Bot
	err := r.db.WithContext(ctx).
		Select("bots.id, bots.uuid as uuid, bots.name, bots.identifier, bots.avatar").
		Where("bots.id IN ?", id).
		Order("bots.id ASC").Find(&bots).Error
	if err != nil {
		return nil, err
	}
	return bots, nil
}

func (r *MysqlChatbotRepository) GetBotsByGroupUuid(ctx context.Context, uuid string) ([]*pb.Bot, error) {
	var bots []*pb.Bot
	err := r.db.WithContext(ctx).
		Select("bots.id, bots.uuid as uuid, bots.name, bots.identifier, bots.avatar").
		Where("groups.uuid = ?", uuid).
		Joins("LEFT JOIN group_bots ON group_bots.bot_id = bots.id").
		Joins("LEFT JOIN `groups` ON groups.id = group_bots.group_id").
		Order("group_bots.id ASC").Find(&bots).Error
	if err != nil {
		return nil, err
	}
	return bots, nil
}

func (r *MysqlChatbotRepository) GetBotsByUser(ctx context.Context, userId int64) ([]*pb.Bot, error) {
	var bots []*pb.Bot
	err := r.db.WithContext(ctx).
		Select("bots.id, groups.id as group_id, bots.name, bots.identifier, bots.avatar").
		Joins("LEFT JOIN group_bots ON group_bots.bot_id = bots.id").
		Joins("LEFT JOIN `groups` ON groups.id = group_bots.group_id").
		Where("groups.user_id = ?", userId).Find(&bots).Error
	if err != nil {
		return nil, err
	}
	return bots, nil
}

func (r *MysqlChatbotRepository) GetBotsByText(ctx context.Context, text []string) (map[string]*pb.Bot, error) {
	result := make(map[string]*pb.Bot)
	if len(text) == 0 {
		return result, nil
	}
	var bots []*pb.Bot
	err := r.db.WithContext(ctx).
		Where("(name IN ?) OR (identifier IN ?)", text, text).Find(&bots).Error
	if err != nil {
		return nil, err
	}

	botsMap := make(map[string]*pb.Bot)
	for i, item := range bots {
		botsMap[item.Name] = bots[i]
		botsMap[item.Identifier] = bots[i]
	}
	for _, s := range text {
		if v, ok := botsMap[s]; ok {
			result[s] = v
		}
	}

	return result, nil
}

func (r *MysqlChatbotRepository) Create(ctx context.Context, bot *pb.Bot) (int64, error) {
	if bot.Uuid == "" {
		return 0, app.ErrInvalidParameter
	}
	bot.Id = r.id.Generate(ctx)
	err := r.db.WithContext(ctx).Create(&bot).Error
	if err != nil {
		return 0, err
	}
	return bot.Id, nil
}

func (r *MysqlChatbotRepository) Update(ctx context.Context, bot *pb.Bot) error {
	return r.db.WithContext(ctx).Model(&bot).Updates(map[string]interface{}{
		"name":   bot.Name,
		"detail": bot.Detail,
		"avatar": bot.Avatar,
		"extend": bot.Extend,
	}).Error
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

func (r *MysqlChatbotRepository) GetGroupByName(ctx context.Context, userId int64, name string) (pb.Group, error) {
	var find pb.Group
	err := r.db.WithContext(ctx).Where("user_id = ? AND name = ?", userId, name).First(&find).Error
	if err != nil {
		return pb.Group{}, err
	}
	return find, nil
}

func (r *MysqlChatbotRepository) TouchGroupUpdatedAt(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Model(&pb.Group{}).Where("id = ?", id).
		Updates(map[string]interface{}{"updated_at": time.Now().Unix()}).Error
}

func (r *MysqlChatbotRepository) ListGroup(ctx context.Context, userId int64) ([]*pb.Group, error) {
	var list []*pb.Group
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Order("updated_at DESC").Find(&list).Error
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
	err = r.db.Where("user_id = ?", group.UserId).Order("sequence DESC").Take(&max).Error
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
	if group.Uuid == "" {
		return 0, app.ErrInvalidParameter
	}
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
	return r.db.WithContext(ctx).Model(&pb.Group{}).Where("id = ?", group.Id).Update("name", group.Name).Error
}

func (r *MysqlChatbotRepository) GetGroupSetting(ctx context.Context, groupId int64) ([]*pb.KV, error) {
	var find []*pb.GroupSetting
	err := r.db.WithContext(ctx).Where("group_id = ?", groupId).Find(&find).Error
	if err != nil {
		return nil, err
	}
	var result []*pb.KV
	for _, item := range find {
		result = append(result, &pb.KV{
			Key:   item.Key,
			Value: item.Value,
		})
	}
	return result, nil
}

func (r *MysqlChatbotRepository) GetGroupSettingByUuid(ctx context.Context, groupUuid string) ([]*pb.KV, error) {
	var find []*pb.GroupSetting
	err := r.db.WithContext(ctx).Where("groups.uuid = ?", groupUuid).
		Joins("LEFT JOIN `groups` ON groups.id = group_settings.group_id").
		Find(&find).Error
	if err != nil {
		return nil, err
	}
	var result []*pb.KV
	for _, item := range find {
		result = append(result, &pb.KV{
			Key:   item.Key,
			Value: item.Value,
		})
	}
	return result, nil
}

func (r *MysqlChatbotRepository) GetGroupBotSetting(ctx context.Context, groupId, botId int64) ([]*pb.KV, error) {
	var find []*pb.GroupBotSetting
	err := r.db.WithContext(ctx).Where("group_id = ? AND bot_id = ?", groupId, botId).Find(&find).Error
	if err != nil {
		return nil, err
	}
	var result []*pb.KV
	for _, item := range find {
		result = append(result, &pb.KV{
			Key:   item.Key,
			Value: item.Value,
		})
	}
	return result, nil
}

func (r *MysqlChatbotRepository) GetGroupBotSettingByUuid(ctx context.Context, groupUuid, botUuid string) ([]*pb.KV, error) {
	var find []*pb.GroupBotSetting
	err := r.db.WithContext(ctx).
		Where("`groups`.uuid = ? AND bots.uuid = ?", groupUuid, botUuid).
		Joins("LEFT JOIN `groups` ON groups.id = group_bot_settings.group_id").
		Joins("LEFT JOIN bots ON bots.id = group_bot_settings.bot_id").
		Find(&find).Error
	if err != nil {
		return nil, err
	}
	var result []*pb.KV
	for _, item := range find {
		result = append(result, &pb.KV{
			Key:   item.Key,
			Value: item.Value,
		})
	}
	return result, nil
}

func (r *MysqlChatbotRepository) GetGroupBotSettingByGroup(ctx context.Context, groupId int64) (map[int64][]*pb.KV, error) {
	var find []*pb.GroupBotSetting
	err := r.db.WithContext(ctx).Where("group_id = ?", groupId).Find(&find).Error
	if err != nil {
		return nil, err
	}
	result := make(map[int64][]*pb.KV)
	for _, item := range find {
		if _, ok := result[item.BotId]; ok {
			result[item.BotId] = append(result[item.BotId], &pb.KV{
				Key:   item.Key,
				Value: item.Value,
			})
		} else {
			result[item.BotId] = []*pb.KV{
				{
					Key:   item.Key,
					Value: item.Value,
				},
			}
		}
	}
	return result, nil
}

func (r *MysqlChatbotRepository) UpdateGroupSetting(ctx context.Context, groupId int64, kvs []*pb.KV) error {
	for _, item := range kvs {
		groupSetting := pb.GroupSetting{
			GroupId:   groupId,
			Key:       item.Key,
			Value:     item.Value,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		if r.db.WithContext(ctx).Model(&groupSetting).Where("group_id = ? AND `key` = ?", groupId, item.Key).
			Updates(map[string]interface{}{"value": item.Value, "updated_at": time.Now().Unix()}).RowsAffected == 0 {
			r.db.WithContext(ctx).Create(&groupSetting)
		}
	}
	return nil
}

func (r *MysqlChatbotRepository) UpdateGroupBotSetting(ctx context.Context, groupId, botId int64, kvs []*pb.KV) error {
	for _, item := range kvs {
		groupBotSetting := pb.GroupBotSetting{
			GroupId:   groupId,
			BotId:     botId,
			Key:       item.Key,
			Value:     item.Value,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		if r.db.WithContext(ctx).Model(&groupBotSetting).Where("group_id = ? AND bot_id = ? AND `key` = ?", groupId, botId, item.Key).
			Updates(map[string]interface{}{"value": item.Value, "updated_at": time.Now().Unix()}).RowsAffected == 0 {
			r.db.WithContext(ctx).Create(&groupBotSetting)
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

func (r *MysqlChatbotRepository) GetTriggerByFlag(ctx context.Context, t, flag string) (pb.Trigger, error) {
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

func (r *MysqlChatbotRepository) ListTriggersByType(ctx context.Context, t string) ([]*pb.Trigger, error) {
	var triggers []*pb.Trigger
	err := r.db.WithContext(ctx).Where("type = ?", t).Find(&triggers).Error
	if err != nil {
		return nil, err
	}
	return triggers, nil
}

func (r *MysqlChatbotRepository) CreateTrigger(ctx context.Context, trigger *pb.Trigger) (int64, error) {
	trigger.Id = r.id.Generate(ctx)
	err := r.db.WithContext(ctx).Create(&trigger).Error
	if err != nil {
		return 0, err
	}
	return trigger.Id, nil
}

func (r *MysqlChatbotRepository) DeleteTriggerByMessageID(ctx context.Context, messageID int64) error {
	return r.db.WithContext(ctx).Where("message_id = ?", messageID).Delete(&pb.Trigger{}).Error
}
