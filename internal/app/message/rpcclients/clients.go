package rpcclients

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewMessageClient, NewSubscribeClient, NewMiddleClient, NewWorkflowClient, NewTaskClient, NewStorageClient)
