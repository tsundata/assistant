package ctx

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
)

type Component struct {
	Bus    event.Bus
	Logger log.Logger

	Middle pb.MiddleClient
	Todo   pb.TodoClient
	User   pb.UserClient
}

func NewComponent() *Component {
	return &Component{}
}
