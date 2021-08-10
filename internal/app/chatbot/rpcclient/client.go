package rpcclient

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewMessageClient, NewMiddleClient, NewUserClient, NewWorkflowClient, NewNLPClient, NewTodoClient, NewStorageClient)
