package context

import (
	"github.com/tsundata/assistant/api/pb"
	"sync"
)

type Context struct {
	mu sync.Mutex

	Value     interface{}
	MidClient pb.MiddleClient
	MsgClient pb.MessageClient
}

func NewContext() *Context {
	return &Context{mu: sync.Mutex{}}
}

func (c *Context) SetValue(v interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Value = v
}
