package context

import (
	"github.com/tsundata/assistant/api/pb"
	"sync"
)

type Context struct {
	mu *sync.RWMutex

	Value     interface{}
	MidClient pb.MiddleClient
	MsgClient pb.MessageClient
}

func NewContext() *Context {
	return &Context{mu: &sync.RWMutex{}}
}

func (c *Context) SetValue(v interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Value = v
}
