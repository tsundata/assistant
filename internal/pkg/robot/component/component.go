package component

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/service"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/mock"
)

type Comp struct {
	Conf   *config.AppConfig
	Bus    event.Bus
	RDB    *redis.Client
	Logger log.Logger

	MessageClient service.MessageSvcClient
	ChatbotClient service.ChatbotSvcClient
	MiddleClient  service.MiddleSvcClient
	StorageClient pb.StorageSvcClient
	UserClient    service.UserSvcClient

	serverFinance pb.FinanceSvcServer
	serverTodo    pb.TodoSvcServer
	serverOkr     pb.OkrSvcServer
}

func (c Comp) Message() service.MessageSvcClient {
	return c.MessageClient
}

func (c Comp) Middle() service.MiddleSvcClient {
	return c.MiddleClient
}

func (c Comp) Chatbot() service.ChatbotSvcClient {
	return c.ChatbotClient
}

func (c Comp) Storage() pb.StorageSvcClient {
	return c.StorageClient
}

func (c Comp) User() service.UserSvcClient {
	return c.UserClient
}

func (c Comp) Todo() pb.TodoSvcServer {
	return c.serverTodo
}

func (c Comp) Okr() pb.OkrSvcServer {
	return c.serverOkr
}

func (c Comp) Finance() pb.FinanceSvcServer {
	return c.serverFinance
}

func (c Comp) GetConfig() *config.AppConfig {
	return c.Conf
}

func (c Comp) GetRedis() *redis.Client {
	return c.RDB
}

func (c Comp) GetLogger() log.Logger {
	return c.Logger
}

func (c Comp) GetBus() event.Bus {
	return c.Bus
}

type Component interface {
	GetConfig() *config.AppConfig
	GetBus() event.Bus
	GetRedis() *redis.Client
	GetLogger() log.Logger
	Message() service.MessageSvcClient
	Chatbot() service.ChatbotSvcClient
	Middle() service.MiddleSvcClient
	Storage() pb.StorageSvcClient
	User() service.UserSvcClient
	Todo() pb.TodoSvcServer
	Okr() pb.OkrSvcServer
	Finance() pb.FinanceSvcServer
}

func NewComponent(
	conf *config.AppConfig,
	bus event.Bus,
	rdb *redis.Client,
	logger log.Logger,

	messageClient service.MessageSvcClient,
	chatbotClient service.ChatbotSvcClient,
	middleClient service.MiddleSvcClient,
	storageClient pb.StorageSvcClient,
	userClient service.UserSvcClient,

	serverFinance pb.FinanceSvcServer,
	serverTodo pb.TodoSvcServer,
	serverOkr pb.OkrSvcServer,
) Component {
	return Comp{
		Conf:          conf,
		Bus:           bus,
		RDB:           rdb,
		Logger:        logger,
		MessageClient: messageClient,
		ChatbotClient: chatbotClient,
		MiddleClient:  middleClient,
		StorageClient: storageClient,
		UserClient:    userClient,
		serverFinance: serverFinance,
		serverTodo:    serverTodo,
		serverOkr:     serverOkr,
	}
}

func MockComponent(deps ...interface{}) Component {
	var (
		conf   *config.AppConfig
		bus    event.Bus
		rdb    *redis.Client
		logger log.Logger

		messageClient service.MessageSvcClient
		middleClient  service.MiddleSvcClient
		chatbotClient service.ChatbotSvcClient
		storageClient pb.StorageSvcClient
		userClient    service.UserSvcClient

		finance pb.FinanceSvcServer
		todo    pb.TodoSvcServer
		okr     pb.OkrSvcServer
	)

	for _, dep := range deps {
		switch v := dep.(type) {
		case *config.AppConfig:
			conf = v
		case event.Bus:
			bus = v
		case *redis.Client:
			rdb = v
		case log.Logger:
			logger = v

		//case *mock.MockMessageSvcClient:
		//	messageClient = v
		//case *mock.MockMiddleSvcClient:
		//	middleClient = v
		//case *mock.MockChatbotSvcClient:
		//	chatbotClient = v
		//case *mock.MockStorageSvcClient:
		//	storageClient = v
		//case *mock.MockUserSvcClient:
		//	userClient = v

		case *mock.MockFinanceSvcServer:
			finance = v
		case *mock.MockTodoSvcServer:
			todo = v
		case *mock.MockOkrSvcServer:
			okr = v
		}
	}

	return Comp{
		Conf:          conf,
		Bus:           bus,
		RDB:           rdb,
		Logger:        logger,
		MessageClient: messageClient,
		MiddleClient:  middleClient,
		ChatbotClient: chatbotClient,
		StorageClient: storageClient,
		UserClient:    userClient,
		serverFinance: finance,
		serverTodo:    todo,
		serverOkr:     okr,
	}
}

var ProviderSet = wire.NewSet(NewComponent)
