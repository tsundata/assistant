package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewId, NewChatbot, NewMessage, NewMiddle, NewTask, NewUser)
