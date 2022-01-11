package automate

import (
	"context"
	"sync"
	"sync/atomic"
)

type Node struct {
	Dependency   []*Edge
	DepCompleted int32
	Task         Runnable
	Children     []*Edge
}

func NewNode(task Runnable) *Node {
	return &Node{Task: task}
}

func (n *Node) dependencyHasDone() bool {
	if n.Dependency == nil {
		return true
	}
	if len(n.Dependency) == 1 {
		return true
	}
	atomic.AddInt32(&n.DepCompleted, 1)

	return n.DepCompleted == int32(len(n.Dependency))
}

func (n *Node) ExecuteWithContext(ctx context.Context, wf *Workflow, i interface{}) {
	if !n.dependencyHasDone() {
		return
	}
	if ctx.Err() != nil {
		wf.interruptDone()
		return
	}
	if n.Task != nil {
		n.Task.Run(i)
	}
	if len(n.Children) > 0 {
		for index := 1; index < len(n.Children); index++ {
			go func(child *Edge) {
				child.ToNode.ExecuteWithContext(ctx, wf, i)
			}(n.Children[index])
		}

		n.Children[0].ToNode.ExecuteWithContext(ctx, wf, i)
	}
}

type EndWorkFlowAction struct {
	done chan struct{}
	s    *sync.Once
}

func (e *EndWorkFlowAction) Run(_ interface{}) {
	e.s.Do(func() {
		e.done <- struct{}{}
	})
}
