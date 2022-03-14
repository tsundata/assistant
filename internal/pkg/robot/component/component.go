package component

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	repositoryTodo "github.com/tsundata/assistant/internal/app/bot/todo/repository"
	"github.com/tsundata/assistant/internal/app/bot/todo/service"
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

	MessageClient pb.MessageSvcClient
	ChatbotClient pb.ChatbotSvcClient
	MiddleClient  pb.MiddleSvcClient
	StorageClient pb.StorageSvcClient
	UserClient    pb.UserSvcClient
	NLPClient     pb.NLPSvcClient

	repoTodo repositoryTodo.TodoRepository
}

func (c Comp) Message() pb.MessageSvcClient {
	return c.MessageClient
}

func (c Comp) Middle() pb.MiddleSvcClient {
	return c.MiddleClient
}

func (c Comp) Chatbot() pb.ChatbotSvcClient {
	return c.ChatbotClient
}

func (c Comp) Storage() pb.StorageSvcClient {
	return c.StorageClient
}

func (c Comp) User() pb.UserSvcClient {
	return c.UserClient
}

func (c Comp) NLP() pb.NLPSvcClient {
	return c.NLPClient
}

func (c Comp) Todo() pb.TodoSvcServer {
	return service.NewTodo(c.Bus, c.Logger, c.repoTodo)
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
	Message() pb.MessageSvcClient
	Chatbot() pb.ChatbotSvcClient
	Middle() pb.MiddleSvcClient
	Storage() pb.StorageSvcClient
	User() pb.UserSvcClient
	NLP() pb.NLPSvcClient
	Todo() pb.TodoSvcServer
}

func NewComponent(
	conf *config.AppConfig,
	bus event.Bus,
	rdb *redis.Client,
	logger log.Logger,

	messageClient pb.MessageSvcClient,
	middleClient pb.MiddleSvcClient,
	storageClient pb.StorageSvcClient,
	userClient pb.UserSvcClient,
	nlpClient pb.NLPSvcClient,

	repoTodo repositoryTodo.TodoRepository,
) Component {
	return Comp{
		Conf:          conf,
		Bus:           bus,
		RDB:           rdb,
		Logger:        logger,
		MessageClient: messageClient,
		MiddleClient:  middleClient,
		StorageClient: storageClient,
		UserClient:    userClient,
		NLPClient:     nlpClient,
		repoTodo:      repoTodo,
	}
}

func MockComponent(deps ...interface{}) Component {
	var (
		conf   *config.AppConfig
		bus    event.Bus
		rdb    *redis.Client
		logger log.Logger

		messageClient pb.MessageSvcClient
		middleClient  pb.MiddleSvcClient
		chatbotClient pb.ChatbotSvcClient
		storageClient pb.StorageSvcClient
		userClient    pb.UserSvcClient
		nlpClient     pb.NLPSvcClient

		repoTodo repositoryTodo.TodoRepository
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

		case *mock.MockMessageSvcClient:
			messageClient = v
		case *mock.MockMiddleSvcClient:
			middleClient = v
		case *mock.MockChatbotSvcClient:
			chatbotClient = v
		case *mock.MockStorageSvcClient:
			storageClient = v
		case *mock.MockUserSvcClient:
			userClient = v
		case *mock.MockNLPSvcClient:
			nlpClient = v

		case repositoryTodo.TodoRepository:
			repoTodo = v
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
		NLPClient:     nlpClient,
		repoTodo:      repoTodo,
	}
}

var ProviderSet = wire.NewSet(NewComponent)
