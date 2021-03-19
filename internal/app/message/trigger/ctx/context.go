package ctx

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/logger"
)

type Context struct {
	Logger    *logger.Logger
	MsgClient pb.MessageClient
	MidClient pb.MiddleClient
}

func NewContext() *Context {
	return &Context{}
}
