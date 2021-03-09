package worker

import (
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"go.uber.org/zap"
)

func NewApp(name string, logger *zap.Logger, rdb *redis.Client, msgClick pb.MessageClient) (*app.Application, error) {
	a, err := app.New(name, logger)
	if err != nil {
		return nil, err
	}

	return a, nil
}
