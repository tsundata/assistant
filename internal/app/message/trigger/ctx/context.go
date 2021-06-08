package ctx

import (
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

type Context struct {
	Logger *logger.Logger
	Client *rpc.Client
}

func NewContext() *Context {
	return &Context{}
}
