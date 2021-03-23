package inside

import (
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
	"sync"
)

type Context struct {
	mu sync.Mutex

	Debug      bool
	Continue   bool
	Value      interface{}
	Credential map[string]string

	RDB        *redis.Client
	MidClient  pb.MiddleClient
	MsgClient  pb.MessageClient
	WfClient   pb.WorkflowClient
	TaskClient pb.TaskClient
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
