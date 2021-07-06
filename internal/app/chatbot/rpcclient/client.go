package rpcclient

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewMessageClient, NewMiddleClient, NewUserClient, NewWorkflowClient, NewSubscribe, NewNLPClient, NewTodoClient, NewStorageClient)
