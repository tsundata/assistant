package ctx

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
)

type Component struct {
	Bus    *event.Bus
	Logger *logger.Logger

	Middle pb.MiddleClient
	Todo   pb.TodoClient
	User   pb.UserClient
}

func NewComponent() *Component {
	return &Component{}
}
