// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package repository

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

// Injectors from wire.go:

func CreateWorkflowRepository(id string) (WorkflowRepository, error) {
	client, err := consul.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(id, client)
	rollbarRollbar := rollbar.New(appConfig)
	logger := log.NewZapLogger(rollbarRollbar)
	logLogger := log.NewAppLogger(logger)
	db, err := mysql.New(appConfig)
	if err != nil {
		return nil, err
	}
	workflowRepository := NewMysqlWorkflowRepository(logLogger, db)
	return workflowRepository, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, mysql.ProviderSet, config.ProviderSet, consul.ProviderSet, ProviderSet, rollbar.ProviderSet)
