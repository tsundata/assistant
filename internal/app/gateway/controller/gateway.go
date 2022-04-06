package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/skip2/go-qrcode"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/gateway/health"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/version"
	"go.uber.org/zap"
	"google.golang.org/grpc/health/grpc_health_v1"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type GatewayController struct {
	config *config.AppConfig
	rdb    *redis.Client
	logger log.Logger
	nr     *newrelic.App
	bus    event.Bus

	messageSvc   pb.MessageSvcClient
	middleSvc    pb.MiddleSvcClient
	userSvc      pb.UserSvcClient
	chatbotSvc   pb.ChatbotSvcClient
	storageSvc   pb.StorageSvcClient
	healthClient *health.Client
}

func NewGatewayController(
	config *config.AppConfig,
	rdb *redis.Client,
	logger log.Logger,
	nr *newrelic.App,
	bus event.Bus,
	messageSvc pb.MessageSvcClient,
	middleSvc pb.MiddleSvcClient,
	chatbotSvc pb.ChatbotSvcClient,
	userSvc pb.UserSvcClient,
	storageSvc pb.StorageSvcClient,
	healthClient *health.Client) *GatewayController {
	return &GatewayController{
		config:       config,
		rdb:          rdb,
		logger:       logger,
		nr:           nr,
		bus:          bus,
		messageSvc:   messageSvc,
		middleSvc:    middleSvc,
		userSvc:      userSvc,
		chatbotSvc:   chatbotSvc,
		storageSvc:   storageSvc,
		healthClient: healthClient,
	}
}

func (gc *GatewayController) Index(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("Gateway %s", version.Version))
}

func (gc *GatewayController) Robots(c *fiber.Ctx) error {
	txt := `User-agent: *
Disallow: /`

	return c.SendString(txt)
}

func (gc *GatewayController) WebhookTrigger(c *fiber.Ctx) error {
	var in pb.TriggerRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.WebhookTrigger(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetPage(c *fiber.Ctx) error {
	var in pb.Page
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.GetPage(md.Outgoing(c), &pb.PageRequest{Page: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) Authorization(c *fiber.Ctx) error {
	var in pb.LoginRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.userSvc.Login(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

// GetApps godoc
// @Summary Get Apps
// @Description get apps
// @ID get-apps
// @Accept json
// @Produce json
// @Success 200 {object} pb.TextRequest
// @Router /apps [get]
func (gc *GatewayController) GetApps(c *fiber.Ctx) error {
	var in pb.TextRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.GetApps(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetMessages(c *fiber.Ctx) error {
	var in pb.GetMessagesRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.ListByGroup(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetMaskingCredentials(c *fiber.Ctx) error {
	var in pb.TextRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.GetMaskingCredentials(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetCredential(c *fiber.Ctx) error {
	var in pb.CredentialRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.GetCredential(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) UpdateSetting(c *fiber.Ctx) error {
	var in pb.KVsRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.CreateCredential(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) Health(c *fiber.Ctx) error {
	str := strings.Builder{}
	gc.healthClient.Status.Range(func(key, value interface{}) bool {
		if service, ok := key.(string); ok {
			str.WriteString(service)
			str.WriteString(": ")
		}
		if status, ok := value.(grpc_health_v1.HealthCheckResponse_ServingStatus); ok {
			str.WriteString(strings.ToLower(grpc_health_v1.HealthCheckResponse_ServingStatus_name[int32(status)]))
			str.WriteString("\n")
		}
		return true
	})
	return c.SendString(str.String())
}

func (gc *GatewayController) GetGroups(c *fiber.Ctx) error {
	var in pb.Group
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.GetGroups(md.Outgoing(c), &pb.GroupRequest{Group: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetGroup(c *fiber.Ctx) error {
	var in pb.Group
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.GetGroup(md.Outgoing(c), &pb.GroupRequest{Group: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetMessage(c *fiber.Ctx) error {
	var in pb.Message
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.GetById(md.Outgoing(c), &pb.MessageRequest{Message: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetUser(c *fiber.Ctx) error {
	var in pb.User
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.userSvc.GetUser(md.Outgoing(c), &pb.UserRequest{User: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetGroupSetting(c *fiber.Ctx) error {
	var in pb.GroupSettingRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.GetGroupSetting(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) UpdateGroupSetting(c *fiber.Ctx) error {
	var in pb.GroupSettingRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.UpdateGroupSetting(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetGroupBots(c *fiber.Ctx) error {
	var in pb.BotsRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.GetBots(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetGroupBot(c *fiber.Ctx) error {
	var in pb.BotRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.GetBot(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) AddGroupBot(c *fiber.Ctx) error {
	var in pb.GroupBotRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.CreateGroupBot(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetGroupBotSetting(c *fiber.Ctx) error {
	var in pb.BotSettingRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.GetGroupBotSetting(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) UpdateGroupBotSetting(c *fiber.Ctx) error {
	var in pb.BotSettingRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.UpdateGroupBotSetting(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetInboxes(c *fiber.Ctx) error {
	var in pb.InboxRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.ListInbox(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetUserSetting(c *fiber.Ctx) error {
	options := vendors.UserCredentialOptions
	return c.JSON(options)
}

func (gc *GatewayController) GetSystemSetting(c *fiber.Ctx) error {
	if c.Locals(enum.AuthKey).(int64) != enum.SuperUserID {
		return c.SendStatus(http.StatusForbidden)
	}
	options := vendors.SystemCredentialOptions
	return c.JSON(options)
}

func (gc *GatewayController) Notify(conn *websocket.Conn, userId int64) {
	gc.logger.Info("[Notify] listening...", zap.Any("user", userId))

	go func() {
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				gc.logger.Warn(err.Error())
				break
			}
			var inbox pb.Inbox
			err = json.Unmarshal(data, &inbox)
			if err != nil {
				gc.logger.Error(err)
				continue
			}
			_, err = gc.messageSvc.MarkReadInbox(md.BuildAuthContext(userId), &pb.InboxRequest{InboxId: inbox.Id})
			if err != nil {
				gc.logger.Error(err)
				continue
			}
		}
	}()

	t := time.NewTicker(10 * time.Second)
	for {
		<-t.C
		err := conn.WriteMessage(websocket.PingMessage, []byte{})
		if err != nil {
			t.Stop()
			gc.logger.Warn(err.Error())
			break
		}

		inbox, err := gc.messageSvc.LastInbox(md.BuildAuthContext(userId), &pb.InboxRequest{})
		if err != nil {
			gc.logger.Error(err)
			continue
		}

		if len(inbox.Inbox) > 0 && inbox.Inbox[0].Id > 0 {
			data, err := json.Marshal(inbox.Inbox[0])
			if err != nil {
				gc.logger.Error(err)
				continue
			}
			err = conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				t.Stop()
				gc.logger.Warn(err.Error())
				break
			}
			_, err = gc.messageSvc.MarkSendInbox(md.BuildAuthContext(userId), &pb.InboxRequest{InboxId: inbox.Inbox[0].Id})
			if err != nil {
				gc.logger.Error(err)
				continue
			}
		}
	}

	gc.logger.Info("[Notify] end", zap.Any("user", userId))
}

func (gc *GatewayController) App(c *fiber.Ctx) error {
	category := c.Params("category")
	provider := vendors.NewOAuthProvider(gc.rdb, category, gc.config.Gateway.Url)
	return provider.Redirect(c, gc.middleSvc)
}

func (gc *GatewayController) OAuth(c *fiber.Ctx) error {
	category := c.Params("category")
	provider := vendors.NewOAuthProvider(gc.rdb, category, gc.config.Gateway.Url)
	err := provider.StoreAccessToken(c, gc.middleSvc)
	if err != nil {
		gc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	return c.SendString("ok")
}

func (gc *GatewayController) QR(c *fiber.Ctx) error {
	text := c.Params("text", "")
	if text == "" {
		return c.SendStatus(http.StatusNotFound)
	}

	txt, err := url.QueryUnescape(text)
	if err != nil {
		gc.logger.Error(err)
		return c.Status(http.StatusNotFound).SendString("error text")
	}

	png, err := qrcode.Encode(txt, qrcode.Medium, 512)
	if err != nil {
		gc.logger.Error(err)
		return c.Status(http.StatusNotFound).SendString("error qr")
	}

	c.Response().Header.Set("Content-Type", "image/png")
	return c.Send(png)
}

func (gc *GatewayController) Webhook(c *fiber.Ctx) error {
	flag := c.Params("flag", "")

	// Headers(Authorization: Base ?) -> query(secret)
	secret := c.Get("Authorization", "")
	secret = strings.ReplaceAll(secret, "Base ", "")
	if secret == "" {
		secret = c.Query("secret", "")
	}

	_, err := gc.chatbotSvc.WebhookTrigger(md.BuildAuthContext(enum.SuperUserID), &pb.TriggerRequest{
		Trigger: &pb.Trigger{
			Type:   "webhook",
			Flag:   flag,
			Secret: secret,
		},
		Info: &pb.TriggerInfo{
			Header: c.Request().Header.String(),
			Body:   util.ByteToString(c.Request().Body()),
		},
	})
	if err != nil {
		gc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.SendString("ok")
}

func (gc *GatewayController) GetFile(c *fiber.Ctx) error {
	path := c.Params("*")
	reply, err := gc.storageSvc.AbsolutePath(md.BuildAuthContext(enum.SuperUserID), &pb.TextRequest{Text: path})
	if err != nil {
		return err
	}
	return c.SendFile(reply.Text, false)
}

func (gc *GatewayController) UploadFile(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	userId, err := getUser(gc.userSvc, token)
	if err != nil {
		return err
	}

	// file
	fh, err := c.FormFile("file")
	if err != nil {
		return err
	}
	f, err := fh.Open()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	// group id
	groupIdStr := c.FormValue("group_id")
	groupId, err := strconv.ParseInt(groupIdStr, 10, 64)
	if err != nil {
		return err
	}
	// save
	reply, err := gc.messageSvc.Create(md.Outgoing(c), &pb.MessageRequest{Message: &pb.Message{
		UserId:  userId,
		Type:    c.FormValue("type"),
		Data:    data,
		GroupId: groupId,
	}})
	if err != nil {
		return err
	}

	return c.JSON(reply)
}

func (gc *GatewayController) MessageAction(c *fiber.Ctx) error {
	var in pb.ActionRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.Action(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) MessageForm(c *fiber.Ctx) error {
	var in pb.FormRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.Form(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) Debug(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	cmd := gc.rdb.Get(context.Background(), fmt.Sprintf("debug:%s", uuid))
	return c.SendString(cmd.Val())
}
