package command

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
)

type Comp struct {
	Conf   *config.AppConfig
	RDB    *redis.Client
	Logger log.Logger

	MessageClient     pb.MessageSvcClient
	MiddleClient      pb.MiddleSvcClient
	WorkflowSvcClient pb.WorkflowSvcClient
	StorageClient     pb.StorageSvcClient
	UserClient        pb.UserSvcClient
	NLPClient         pb.NLPSvcClient
}

func (c Comp) Message() pb.MessageSvcClient {
	return c.MessageClient
}

func (c Comp) Middle() pb.MiddleSvcClient {
	return c.MiddleClient
}

func (c Comp) Workflow() pb.WorkflowSvcClient {
	return c.WorkflowSvcClient
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

func (c Comp) GetConfig() *config.AppConfig {
	return c.Conf
}

func (c Comp) GetRedis() *redis.Client {
	return c.RDB
}

func (c Comp) GetLogger() log.Logger {
	return c.Logger
}

type Component interface {
	GetConfig() *config.AppConfig
	GetRedis() *redis.Client
	GetLogger() log.Logger
	Message() pb.MessageSvcClient
	Middle() pb.MiddleSvcClient
	Workflow() pb.WorkflowSvcClient
	Storage() pb.StorageSvcClient
	User() pb.UserSvcClient
	NLP() pb.NLPSvcClient
}

func NewComponent(
	conf *config.AppConfig,
	rdb *redis.Client,
	logger log.Logger,

	messageClient pb.MessageSvcClient,
	middleClient pb.MiddleSvcClient,
	workflowClient pb.WorkflowSvcClient,
	storageClient pb.StorageSvcClient,
	userClient pb.UserSvcClient,
	nlpClient pb.NLPSvcClient,
) Component {
	return Comp{
		Conf:              conf,
		RDB:               rdb,
		Logger:            logger,
		MessageClient:     messageClient,
		MiddleClient:      middleClient,
		WorkflowSvcClient: workflowClient,
		StorageClient:     storageClient,
		UserClient:        userClient,
		NLPClient:         nlpClient,
	}
}

var ProviderSet = wire.NewSet(NewComponent)
