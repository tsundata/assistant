package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/influxdata/cron"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot"
	"github.com/tsundata/assistant/internal/pkg/robot/action"
	"github.com/tsundata/assistant/internal/pkg/robot/action/opcode"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/exception"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type Chatbot struct {
	logger  log.Logger
	bus     event.Bus
	rdb     *redis.Client
	bot     *rulebot.RuleBot
	repo    repository.ChatbotRepository
	message pb.MessageSvcClient
	middle  pb.MiddleSvcClient
	comp    component.Component
}

func NewChatbot(
	logger log.Logger,
	bus event.Bus,
	rdb *redis.Client,
	repo repository.ChatbotRepository,
	message pb.MessageSvcClient,
	middle pb.MiddleSvcClient,
	bot *rulebot.RuleBot,
	comp component.Component) *Chatbot {
	return &Chatbot{
		logger:  logger,
		bus:     bus,
		rdb:     rdb,
		bot:     bot,
		repo:    repo,
		message: message,
		middle:  middle,
		comp:    comp,
	}
}

func (s *Chatbot) Handle(ctx context.Context, payload *pb.ChatbotRequest) (*pb.ChatbotReply, error) {
	id, _ := md.FromIncoming(ctx)
	reply, err := s.message.GetById(ctx, &pb.MessageRequest{Message: &pb.Message{Id: payload.MessageId}})
	if err != nil {
		return nil, err
	}

	r := robot.NewRobot()

	// touch updated
	err = s.repo.TouchGroupUpdatedAt(ctx, reply.Message.GroupId)
	if err != nil {
		return nil, err
	}

	// group bots
	groupBots, err := s.repo.ListGroupBot(ctx, reply.Message.GetGroupId())
	if err != nil {
		return nil, err
	}
	groupBotsMap := make(map[string]struct{})
	for _, item := range groupBots {
		groupBotsMap[item.Identifier] = struct{}{}
	}

	// help
	outMessages, err := r.Help(groupBots, reply.Message.GetText())
	if err != nil {
		return nil, err
	}
	if len(outMessages) > 0 {
		docReply, err := s.ActionDoc(ctx, &pb.WorkflowRequest{})
		if err != nil {
			return nil, err
		}
		outMessages[0] = []pb.MsgPayload{
			pb.TextMsg{Text: docReply.Text},
		}
	}

	if len(outMessages) == 0 {
		// bot settings
		groupSetting, err := s.repo.GetGroupSetting(ctx, reply.Message.GetGroupId())
		if err != nil {
			return nil, err
		}
		groupBotSetting, err := s.repo.GetGroupBotSettingByGroup(ctx, reply.Message.GetGroupId())
		if err != nil {
			return nil, err
		}

		// merge setting
		setting := make(map[int64][]*pb.KV)
		setting[reply.Message.GetGroupId()] = groupSetting
		for i, kvs := range groupBotSetting {
			setting[i] = kvs
		}

		botCtx := bot.Context{Setting: setting}

		// trigger
		err = r.ProcessTrigger(ctx, botCtx, s.comp, reply.Message)
		if err != nil {
			return nil, err
		}

		// lexer
		tokens, objects, tags, messages, commands, err := r.ParseText(reply.Message)
		if err != nil {
			return nil, err
		}

		// filter bots
		objectBots, err := s.repo.GetBotsByText(ctx, objects)
		if err != nil {
			return nil, err
		}
		for k, item := range objectBots {
			if _, ok := groupBotsMap[item.Identifier]; !ok {
				delete(objectBots, k)
			}
		}

		if len(commands) > 0 {
			var commandsBots []*pb.Bot
			if len(objects) == 0 {
				commandsBots = groupBots
			} else {
				for k := range objectBots {
					commandsBots = append(commandsBots, objectBots[k])
				}
			}
			for _, item := range commandsBots {
				for _, commandText := range commands {
					commandMessages, err := r.ProcessCommand(ctx, botCtx, s.comp, item.Identifier, commandText)
					if err != nil {
						return nil, err
					}
					outMessages[item.Id] = commandMessages
				}
			}
		}

		// messages
		if len(messages) > 0 {
			for _, mid := range messages {
				sequence, _ := strconv.ParseInt(mid, 10, 64)
				messageReply, _ := s.message.GetBySequence(ctx, &pb.MessageRequest{Message: &pb.Message{UserId: id, Sequence: sequence}})
				if messageReply != nil && messageReply.Message.GetId() > 0 {
					messageReply.Message.Direction = enum.MessageIncomingDirection
					_ = s.bus.Publish(ctx, enum.Message, event.MessageChannelSubject, messageReply.Message)
				}
			}
		}

		// workflow
		if len(outMessages) == 0 {
			outMessages, err = r.ProcessWorkflow(ctx, botCtx, s.comp, tokens, objectBots)
			if err != nil {
				return nil, err
			}
		}

		// tags
		if len(tags) > 0 {
			var tagBots []*pb.Bot
			if len(objects) == 0 {
				tagBots = groupBots
			} else {
				for k := range objectBots {
					tagBots = append(tagBots, objectBots[k])
				}
			}
			for _, tag := range tags {
				// save tag
				_, err = s.middle.SaveModelTag(ctx, &pb.ModelTagRequest{
					Model: &pb.ModelTag{
						Service: enum.Message,
						Model:   util.ModelName(pb.Message{}),
						ModelId: reply.Message.GetId(),
					},
					Tag: tag,
				})
				if err != nil {
					s.logger.Error(err)
				}
				// trigger tag
				for _, item := range tagBots {
					result, err := r.ProcessTag(ctx, botCtx, s.comp, item.Identifier, tag)
					if err != nil {
						return nil, err
					}
					outMessages[item.Id] = append(outMessages[item.Id], result...)
				}
			}
		}
	}

	// send message
	for botId, messages := range outMessages {
		for _, item := range messages {
			text := ""
			if item.Type() == enum.MessageTypeText {
				if v, ok := item.(pb.TextMsg); ok {
					text = v.Text
				}
			}
			j, err := json.Marshal(item)
			if err != nil {
				return nil, err
			}
			outMessage := &pb.Message{
				GroupId:      reply.Message.GetGroupId(),
				UserId:       id,
				Sender:       botId,
				SenderType:   enum.MessageBotType,
				Receiver:     id,
				ReceiverType: enum.MessageUserType,
				Type:         string(item.Type()),
				Text:         text,
				Payload:      string(j),
				Status:       0,
				Direction:    enum.MessageIncomingDirection,
				SendTime:     util.Format(time.Now().Unix()),
			}
			messageReply, err := s.message.Save(ctx, &pb.MessageRequest{Message: outMessage})
			if err != nil {
				return nil, err
			}
			err = s.bus.Publish(ctx, enum.Message, event.MessageChannelSubject, messageReply.Message)
			if err != nil {
				return nil, err
			}
		}
	}

	return &pb.ChatbotReply{
		State: true,
	}, nil
}

func (s *Chatbot) Action(ctx context.Context, payload *pb.BotRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)
	b, err := s.repo.GetByID(ctx, payload.BotId)
	if err != nil {
		return nil, err
	}

	// process
	botCtx := bot.Context{}
	r := robot.NewRobot()
	msg, err := r.ProcessAction(ctx, botCtx, s.comp, b.Identifier, payload.ActionId, payload.Value)
	if err != nil {
		return nil, err
	}

	// send msg
	for _, item := range msg {
		text := ""
		if item.Type() == enum.MessageTypeText {
			if v, ok := item.(pb.TextMsg); ok {
				text = v.Text
			}
		}
		j, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		outMessage := &pb.Message{
			GroupId:      payload.GroupId,
			UserId:       id,
			Sender:       payload.BotId,
			SenderType:   enum.MessageBotType,
			Receiver:     id,
			ReceiverType: enum.MessageUserType,
			Type:         string(item.Type()),
			Text:         text,
			Payload:      string(j),
			Status:       0,
			Direction:    enum.MessageIncomingDirection,
			SendTime:     util.Format(time.Now().Unix()),
		}
		messageReply, err := s.message.Save(ctx, &pb.MessageRequest{Message: outMessage})
		if err != nil {
			return nil, err
		}
		err = s.bus.Publish(ctx, enum.Message, event.MessageChannelSubject, messageReply.Message)
		if err != nil {
			return nil, err
		}
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) Form(ctx context.Context, payload *pb.BotRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)
	b, err := s.repo.GetByID(ctx, payload.BotId)
	if err != nil {
		return nil, err
	}

	// process
	var field []bot.FieldItem
	for _, item := range payload.Form {
		field = append(field, bot.FieldItem{
			Key:   item.Key,
			Value: item.Value,
		})
	}
	botCtx := bot.Context{FieldItem: field}
	r := robot.NewRobot()
	msg, err := r.ProcessForm(ctx, botCtx, s.comp, b.Identifier, payload.FormId)
	if err != nil {
		return nil, err
	}

	// send msg
	for _, item := range msg {
		text := ""
		if item.Type() == enum.MessageTypeText {
			if v, ok := item.(pb.TextMsg); ok {
				text = v.Text
			}
		}
		j, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		outMessage := &pb.Message{
			GroupId:      payload.GroupId,
			UserId:       id,
			Sender:       payload.BotId,
			SenderType:   enum.MessageBotType,
			Receiver:     id,
			ReceiverType: enum.MessageUserType,
			Type:         string(item.Type()),
			Text:         text,
			Payload:      string(j),
			Status:       0,
			Direction:    enum.MessageIncomingDirection,
			SendTime:     util.Format(time.Now().Unix()),
		}
		messageReply, err := s.message.Save(ctx, &pb.MessageRequest{Message: outMessage})
		if err != nil {
			return nil, err
		}
		err = s.bus.Publish(ctx, enum.Message, event.MessageChannelSubject, messageReply.Message)
		if err != nil {
			return nil, err
		}
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) Register(ctx context.Context, payload *pb.BotRequest) (*pb.StateReply, error) {
	b, err := s.repo.GetByIdentifier(ctx, payload.Bot.Identifier)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if b.Id > 0 {
		payload.Bot.Id = b.Id
		err = s.repo.Update(ctx, payload.Bot)
		if err != nil {
			return nil, err
		}
		return &pb.StateReply{State: true}, nil
	}

	avatarReply, err := s.middle.CreateAvatar(ctx, &pb.TextRequest{Text: payload.Bot.Identifier})
	if err != nil {
		return nil, err
	}
	payload.Bot.Avatar = avatarReply.Text
	_, err = s.repo.Create(ctx, payload.Bot)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) GetBot(ctx context.Context, payload *pb.BotRequest) (*pb.BotReply, error) {
	b, err := s.repo.GetGroupBot(ctx, payload.GroupId, payload.BotId)
	if err != nil {
		return nil, err
	}
	return &pb.BotReply{Bot: &b}, nil
}

func (s *Chatbot) GetBots(ctx context.Context, payload *pb.BotsRequest) (*pb.BotsReply, error) {
	var err error
	var bots []*pb.Bot
	if payload.GroupId > 0 {
		bots, err = s.repo.GetBotsByGroup(ctx, payload.GroupId)
	}
	if len(payload.BotId) > 0 {
		bots, err = s.repo.GetBotsByIds(ctx, payload.BotId)
	}
	if err != nil {
		return nil, err
	}
	return &pb.BotsReply{Bots: bots}, nil
}

func (s *Chatbot) UpdateBotSetting(_ context.Context, _ *pb.BotSettingRequest) (*pb.StateReply, error) {
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) GetGroups(ctx context.Context, _ *pb.GroupRequest) (*pb.GetGroupsReply, error) {
	id, _ := md.FromIncoming(ctx)

	// default group
	const defaultGroupName = "System"
	_, err := s.repo.GetGroupByName(ctx, id, defaultGroupName)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		avatarReply, err := s.middle.CreateAvatar(ctx, &pb.TextRequest{Text: defaultGroupName})
		if err != nil {
			return nil, err
		}
		_, err = s.repo.CreateGroup(ctx, &pb.Group{
			UserId:    id,
			Name:      defaultGroupName,
			Avatar:    avatarReply.Text,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		})
		if err != nil {
			return nil, err
		}
		// todo group --- system bot
	}

	// bot avatar
	bots, err := s.repo.GetBotsByUser(ctx, id)
	if err != nil {
		return nil, err
	}

	botAvatar := make(map[int64][]*pb.Avatar)
	for _, item := range bots {
		botAvatar[item.GroupId] = append(botAvatar[item.GroupId], &pb.Avatar{
			Name:       item.Name,
			Src:        item.Avatar,
			Identifier: item.Identifier,
		})
	}

	groups, err := s.repo.ListGroup(ctx, id)
	if err != nil {
		return nil, err
	}
	var res []*pb.GroupItem
	for _, group := range groups {
		// avatar
		var avatars []*pb.Avatar
		if v, ok := botAvatar[group.Id]; ok {
			avatars = v
		}

		// last message
		last, err := s.message.LastByGroup(ctx, &pb.LastByGroupRequest{GroupId: group.Id})
		if err != nil {
			return nil, err
		}

		res = append(res, &pb.GroupItem{
			Id:          group.Id,
			Sequence:    group.Sequence,
			Type:        group.Type,
			Name:        group.Name,
			Avatar:      group.Avatar,
			UnreadCount: 0, // todo
			LastMessage: &pb.LastMessage{
				LastSender: last.Message.GetSenderName(),
				Content:    last.Message.GetText(),
			},
			BotAvatar: avatars,
		})
	}
	return &pb.GetGroupsReply{Groups: res}, nil
}

func (s *Chatbot) CreateGroup(ctx context.Context, payload *pb.GroupRequest) (*pb.StateReply, error) {
	avatarReply, err := s.middle.CreateAvatar(ctx, &pb.TextRequest{Text: payload.Group.Name})
	if err != nil {
		return nil, err
	}
	payload.Group.Avatar = avatarReply.Text
	_, err = s.repo.CreateGroup(ctx, payload.Group)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) GetGroup(ctx context.Context, payload *pb.GroupRequest) (*pb.GetGroupReply, error) {
	id, _ := md.FromIncoming(ctx)
	group, err := s.repo.GetGroup(ctx, payload.Group.GetId())
	if err != nil {
		return nil, err
	}
	if group.UserId != id {
		return nil, exception.ErrGrpcUnauthenticated
	}
	return &pb.GetGroupReply{Group: &pb.GroupItem{
		Id:       group.Id,
		Sequence: group.Sequence,
		Type:     group.Type,
		Name:     group.Name,
		Avatar:   group.Avatar,
	}}, nil
}

func (s *Chatbot) CreateGroupBot(ctx context.Context, payload *pb.GroupBotRequest) (*pb.StateReply, error) {
	err := s.repo.CreateGroupBot(ctx, payload.GroupId, payload.Bot)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) DeleteGroupBot(ctx context.Context, payload *pb.GroupBotRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteGroupBot(ctx, payload.GroupId, payload.Bot.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) UpdateGroupBotSetting(ctx context.Context, payload *pb.BotSettingRequest) (*pb.StateReply, error) {
	group, err := s.repo.GetGroup(ctx, payload.GroupId)
	if err != nil {
		return nil, err
	}
	b, err := s.repo.GetByID(ctx, payload.BotId)
	if err != nil {
		return nil, err
	}
	err = s.repo.UpdateGroupBotSetting(ctx, group.Id, b.Id, payload.Kvs)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) UpdateGroupSetting(ctx context.Context, payload *pb.GroupSettingRequest) (*pb.StateReply, error) {
	group, err := s.repo.GetGroup(ctx, payload.GroupId)
	if err != nil {
		return nil, err
	}
	err = s.repo.UpdateGroupSetting(ctx, group.Id, payload.Kvs)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) DeleteGroup(ctx context.Context, payload *pb.GroupRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteGroup(ctx, payload.Group.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) UpdateGroup(ctx context.Context, payload *pb.GroupRequest) (*pb.StateReply, error) {
	err := s.repo.UpdateGroup(ctx, payload.Group)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) GetGroupBotSetting(ctx context.Context, payload *pb.BotSettingRequest) (*pb.BotSettingReply, error) {
	kv, err := s.repo.GetGroupBotSetting(ctx, payload.GroupId, payload.BotId)
	if err != nil {
		return nil, err
	}
	return &pb.BotSettingReply{Kvs: kv}, nil
}

func (s *Chatbot) GetGroupSetting(ctx context.Context, payload *pb.GroupSettingRequest) (*pb.GroupSettingReply, error) {
	kv, err := s.repo.GetGroupSetting(ctx, payload.GroupId)
	if err != nil {
		return nil, err
	}
	return &pb.GroupSettingReply{Kvs: kv}, nil
}

func (s *Chatbot) SyntaxCheck(_ context.Context, payload *pb.WorkflowRequest) (*pb.StateReply, error) {
	switch enum.MessageType(payload.Type) {
	case enum.MessageTypeScript:
		if payload.GetText() == "" {
			return nil, errors.New("empty action")
		}
		p, err := action.NewParser(action.NewLexer([]rune(payload.GetText())))
		if err != nil {
			return &pb.StateReply{State: false}, err
		}
		tree, err := p.Parse()
		if err != nil {
			return &pb.StateReply{State: false}, err
		}

		symbolTable := action.NewSemanticAnalyzer()
		err = symbolTable.Visit(tree)
		if err != nil {
			return &pb.StateReply{State: false}, err
		}

		return &pb.StateReply{State: true}, nil
	default:
		return &pb.StateReply{State: false}, nil
	}
}

func (s *Chatbot) RunActionScript(ctx context.Context, payload *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
	if payload.Message.GetText() == "" {
		return nil, errors.New("empty action")
	}
	p, err := action.NewParser(action.NewLexer([]rune(payload.Message.GetText())))
	if err != nil {
		return nil, err
	}
	tree, err := p.Parse()
	if err != nil {
		return nil, err
	}

	symbolTable := action.NewSemanticAnalyzer()
	err = symbolTable.Visit(tree)
	if err != nil {
		return nil, err
	}

	i := action.NewInterpreter(ctx, tree)
	i.SetMessage(*payload.Message)
	i.SetComponent(s.comp)
	_, err = i.Interpret()
	if err != nil {
		return nil, err
	}

	var result string
	if i.Debug {
		result = fmt.Sprintf("Tracing\n-------\n%s", i.Stdout())
	}

	return &pb.WorkflowReply{
		Text: result,
	}, nil
}

func (s *Chatbot) WebhookTrigger(ctx context.Context, payload *pb.TriggerRequest) (*pb.WorkflowReply, error) {
	trigger, err := s.repo.GetTriggerByFlag(ctx, enum.TriggerWebhookType, payload.Trigger.GetFlag())
	if err != nil {
		return nil, err
	}

	if trigger.Status == enum.TriggerDisable {
		return nil, errors.New("webhook trigger disable")
	}

	// Authorization
	if trigger.Secret != "" && payload.Trigger.GetSecret() != trigger.Secret {
		return nil, errors.New("error secret")
	}

	if trigger.MessageId <= 0 {
		return nil, errors.New("error trigger")
	}

	// get message
	message, err := s.message.GetById(ctx, &pb.MessageRequest{Message: &pb.Message{Id: trigger.MessageId}})
	if err != nil {
		return nil, err
	}

	// publish event
	err = s.bus.Publish(ctx, enum.Chatbot, event.ScriptRunSubject, message.Message)
	if err != nil {
		return nil, err
	}

	return &pb.WorkflowReply{}, nil
}

func (s *Chatbot) CronTrigger(ctx context.Context, _ *pb.TriggerRequest) (*pb.WorkflowReply, error) {
	triggers, err := s.repo.ListTriggersByType(ctx, enum.TriggerCronType)
	if err != nil {
		return nil, err
	}

	for _, trigger := range triggers {
		if trigger.Status == enum.TriggerDisable {
			continue
		}
		var lastTime time.Time
		key := fmt.Sprintf("chatbot:cron:%d:time", trigger.MessageId)
		t := s.rdb.Get(ctx, key).Val()
		if t == "" {
			lastTime = time.Time{}
		} else {
			lastTime, err = time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
			if err != nil {
				return nil, err
			}
		}

		p, err := cron.ParseUTC(trigger.When)
		if err != nil {
			return nil, err
		}
		nextTime, err := p.Next(lastTime)
		if err != nil {
			return nil, err
		}

		now := time.Now()
		if nextTime.Before(now) {
			// time
			s.rdb.Set(ctx, key, now.Format("2006-01-02 15:04:05"), 0)

			// get message
			message, err := s.message.GetById(ctx, &pb.MessageRequest{Message: &pb.Message{Id: trigger.MessageId}})
			if err != nil {
				return nil, err
			}

			// publish event
			err = s.bus.Publish(ctx, enum.Chatbot, event.ScriptRunSubject, message.Message)
			if err != nil {
				return nil, err
			}
		}
	}

	return &pb.WorkflowReply{}, nil
}

func (s *Chatbot) CreateTrigger(ctx context.Context, payload *pb.TriggerRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)
	messageReply, err := s.message.GetById(ctx, &pb.MessageRequest{Message: &pb.Message{Id: payload.Trigger.GetMessageId()}})
	if err != nil {
		return nil, err
	}
	var trigger pb.Trigger
	trigger.Type = payload.Trigger.GetType()
	trigger.Kind = payload.Trigger.GetKind()
	trigger.UserId = id
	trigger.MessageId = messageReply.Message.GetId()
	trigger.Status = enum.TriggerEnable

	switch enum.MessageType(payload.Trigger.GetKind()) {
	case enum.MessageTypeScript:
		if payload.Info.GetMessageText() == "" {
			return nil, errors.New("empty action")
		}
		p, err := action.NewParser(action.NewLexer([]rune(payload.Info.GetMessageText())))
		if err != nil {
			return nil, err
		}
		tree, err := p.Parse()
		if err != nil {
			return nil, err
		}

		symbolTable := action.NewSemanticAnalyzer()
		err = symbolTable.Visit(tree)
		if err != nil {
			return nil, err
		}

		if symbolTable.Cron == nil && symbolTable.Webhook == nil {
			return &pb.StateReply{State: false}, nil
		}

		if symbolTable.Cron != nil {
			trigger.Type = enum.TriggerCronType
			trigger.When = symbolTable.Cron.When

			// store
			_, err = s.repo.CreateTrigger(ctx, &trigger)
			if err != nil {
				return nil, err
			}
		}

		if symbolTable.Webhook != nil {
			trigger.Type = enum.TriggerWebhookType
			trigger.Flag = symbolTable.Webhook.Flag
			trigger.Secret = symbolTable.Webhook.Secret

			find, err := s.repo.GetTriggerByFlag(ctx, trigger.Type, trigger.Flag)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}

			if find.Id > 0 {
				return nil, errors.New("exist flag: " + trigger.Flag)
			}

			// store
			_, err = s.repo.CreateTrigger(ctx, &trigger)
			if err != nil {
				return nil, err
			}
		}

		return &pb.StateReply{State: true}, nil
	default:
		return &pb.StateReply{State: false}, nil
	}
}

func (s *Chatbot) DeleteTrigger(ctx context.Context, payload *pb.TriggerRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteTriggerByMessageID(ctx, payload.Trigger.GetMessageId())
	if err != nil {
		return &pb.StateReply{State: false}, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) ActionDoc(_ context.Context, payload *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
	var docs string
	if payload.GetText() == "" {
		docs = fmt.Sprintf("--- Action Script ---\n%s", strings.Join(opcode.Docs(), "\n"))
	} else {
		docs = opcode.Doc(payload.GetText())
	}
	return &pb.WorkflowReply{
		Text: docs,
	}, nil
}

func (s *Chatbot) ListWebhook(ctx context.Context, _ *pb.WorkflowRequest) (*pb.WebhooksReply, error) {
	triggers, err := s.repo.ListTriggersByType(ctx, enum.TriggerWebhookType)
	if err != nil {
		return nil, err
	}
	var flags []string
	for _, item := range triggers {
		flags = append(flags, item.Flag)
	}
	return &pb.WebhooksReply{Flag: flags}, nil
}

func (s *Chatbot) GetWebhookTriggers(ctx context.Context, _ *pb.TriggerRequest) (*pb.TriggersReply, error) {
	id, _ := md.FromIncoming(ctx)
	triggers, err := s.repo.GetTriggers(ctx, id, enum.TriggerWebhookType)
	if err != nil {
		return nil, err
	}
	return &pb.TriggersReply{List: triggers}, nil
}

func (s *Chatbot) GetCronTriggers(ctx context.Context, _ *pb.TriggerRequest) (*pb.TriggersReply, error) {
	id, _ := md.FromIncoming(ctx)
	triggers, err := s.repo.GetTriggers(ctx, id, enum.TriggerCronType)
	if err != nil {
		return nil, err
	}

	// sequence
	var messageId []int64
	for _, item := range triggers {
		messageId = append(messageId, item.MessageId)
	}
	messageReply, err := s.message.GetByIds(ctx, &pb.GetMessagesRequest{Ids: messageId})
	if err != nil {
		return nil, err
	}
	messageMap := make(map[int64]int64)
	for _, item := range messageReply.Messages {
		messageMap[item.Id] = item.Sequence
	}
	for i, item := range triggers {
		if s, ok := messageMap[item.MessageId]; ok {
			triggers[i].Sequence = s
		}
	}

	return &pb.TriggersReply{List: triggers}, nil
}

func (s *Chatbot) SwitchTriggers(ctx context.Context, payload *pb.SwitchTriggersRequest) (*pb.StateReply, error) {
	for _, item := range payload.Triggers {
		messageId, err := strconv.ParseInt(item.Key, 10, 64)
		if err != nil {
			return nil, err
		}

		var status int64
		if item.Value == "1" {
			status = enum.TriggerEnable
		} else {
			status = enum.TriggerDisable
		}

		err = s.repo.SwitchTrigger(ctx, messageId, status)
		if err != nil {
			return nil, err
		}
	}
	return &pb.StateReply{State: true}, nil
}
