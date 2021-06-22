package inside

import (
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"sync"
)

type Context struct {
	mu sync.Mutex

	Debug      bool
	Continue   bool
	Value      interface{}
	Credential map[string]string

	RDB     *redis.Client
	Bus     *event.Bus
	Logger  *logger.Logger
	Middle  pb.MiddleClient
	Message pb.MessageClient
}

func NewContext() *Context {
	return &Context{mu: sync.Mutex{}, Continue: true}
}

func (c *Context) SetValue(v interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Value = v
}

func (c *Context) SetContinue(b bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Continue = b
}

func (c *Context) SetCredential(v map[string]string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Credential = v
}
