package ctx

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
)

type Component struct {
	Bus    event.Bus
	Logger log.Logger

	Middle pb.MiddleSvcClient
	Todo   pb.TodoSvcClient
	User   pb.UserSvcClient
}

func NewComponent() *Component {
	return &Component{}
}
