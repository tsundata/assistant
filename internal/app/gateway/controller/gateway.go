package controller

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/version"
	"net/http"
	"strings"
)

type GatewayController struct {
	opt    *config.AppConfig
	rdb    *redis.Client
	logger log.Logger
	nr     *newrelic.App
	bus    event.Bus

	messageSvc  pb.MessageSvcClient
	middleSvc   pb.MiddleSvcClient
	workflowSvc pb.WorkflowSvcClient
	userSvc     pb.UserSvcClient
	chatbotSvc  pb.ChatbotSvcClient
}

func NewGatewayController(
	opt *config.AppConfig,
	rdb *redis.Client,
	logger log.Logger,
	nr *newrelic.App,
	bus event.Bus,
	messageSvc pb.MessageSvcClient,
	middleSvc pb.MiddleSvcClient,
	workflowSvc pb.WorkflowSvcClient,
	chatbotSvc pb.ChatbotSvcClient,
	userSvc pb.UserSvcClient) *GatewayController {
	return &GatewayController{
		opt:         opt,
		rdb:         rdb,
		logger:      logger,
		nr:          nr,
		bus:         bus,
		messageSvc:  messageSvc,
		middleSvc:   middleSvc,
		workflowSvc: workflowSvc,
		userSvc:     userSvc,
		chatbotSvc:  chatbotSvc,
	}
}

func (gc *GatewayController) Index(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("Gateway %s", version.Version))
}

func (gc *GatewayController) DebugEvent(c *fiber.Ctx) error {
	// chatbot handle
	text := util.ByteToString(c.Body())
	reply, err := gc.chatbotSvc.Handle(context.Background(), &pb.ChatbotRequest{Text: text})
	if err != nil {
		gc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	if len(reply.GetText()) > 0 {
		return c.Send(util.StringToByte(strings.Join(reply.GetText(), "\n")))
	}

	// or create message
	uuid := util.UUID()
	messageReply, err := gc.messageSvc.Create(context.Background(), &pb.MessageRequest{
		Message: &pb.Message{
			Uuid: uuid,
			Text: text,
		},
	})
	if err != nil {
		return err
	}
	return c.Send(util.StringToByte(fmt.Sprintf("ID: %d", messageReply.Message.GetId())))
}

func (gc *GatewayController) GetChart(c *fiber.Ctx) error {
	var in pb.ChartData
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.GetChartData(md.Outgoing(c), &pb.ChartDataRequest{ChartData: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) WebhookTrigger(c *fiber.Ctx) error {
	var in pb.TriggerRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.workflowSvc.WebhookTrigger(md.Outgoing(c), &in)
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

func (gc *GatewayController) StoreAppOAuth(c *fiber.Ctx) error {
	var in pb.AppRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.StoreAppOAuth(md.Outgoing(c), &in)
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
	var in pb.Message
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.List(md.Outgoing(c), &pb.MessageRequest{Message: &in})
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

func (gc *GatewayController) CreateCredential(c *fiber.Ctx) error {
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

func (gc *GatewayController) GetSettings(c *fiber.Ctx) error {
	var in pb.TextRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.GetSettings(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) CreateSetting(c *fiber.Ctx) error {
	var in pb.KVRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.CreateSetting(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetActionMessages(c *fiber.Ctx) error {
	var in pb.TextRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.GetActionMessages(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) CreateActionMessage(c *fiber.Ctx) error {
	var in pb.TextRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.CreateActionMessage(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) DeleteWorkflowMessage(c *fiber.Ctx) error {
	var in pb.Message
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.DeleteWorkflowMessage(md.Outgoing(c), &pb.MessageRequest{Message: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) RunMessage(c *fiber.Ctx) error {
	var in pb.Message
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.Run(md.Outgoing(c), &pb.MessageRequest{Message: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) SendMessage(c *fiber.Ctx) error {
	var in pb.Message
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.Send(md.Outgoing(c), &pb.MessageRequest{Message: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetRoleImage(c *fiber.Ctx) error {
	var in pb.RoleRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.userSvc.GetRoleImage(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}
