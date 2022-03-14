package inside

import (
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"sync"
)

type Component struct {
	mu sync.Mutex

	Debug      bool
	Continue   bool
	Value      interface{}
	Credential map[string]string
	Message    pb.Message

	RDB           *redis.Client
	Bus           event.Bus
	Logger        log.Logger
	Middle        pb.MiddleSvcClient
	MessageClient pb.MessageSvcClient
}

func NewComponent() *Component {
	return &Component{mu: sync.Mutex{}, Continue: true}
}

func (c *Component) SetValue(v interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Value = v
}

func (c *Component) SetContinue(b bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Continue = b
}

func (c *Component) SetCredential(v map[string]string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Credential = v
}

func (c *Component) SetMessage(message pb.Message) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Message = message
}
