package automate

import (
	"context"
	"sync"
)

type Workflow struct {
	done        chan struct{}
	doneOnce    *sync.Once
	alreadyDone bool
	root        *Node
	End         *Node
	edges       []*Edge
}

func NewWorkflow() *Workflow {
	wf := &Workflow{
		done:     make(chan struct{}, 1),
		doneOnce: &sync.Once{},
		root:     &Node{Task: nil},
	}

	endNode := &EndWorkFlowAction{
		done: wf.done,
		s:    wf.doneOnce,
	}
	wf.End = NewNode(endNode)

	return wf
}

func (w *Workflow) AddStartNode(node *Node) {
	w.edges = append(w.edges, AddEdge(w.root, node))
}

func (w *Workflow) AddEdge(from, to *Node) {
	w.edges = append(w.edges, AddEdge(from, to))
}

func (w *Workflow) ConnectToEnd(node *Node) {
	w.edges = append(w.edges, AddEdge(node, w.End))
}

func (w *Workflow) StartWithContext(ctx context.Context, i interface{}) {
	w.root.ExecuteWithContext(ctx, w, i)
}

func (w *Workflow) WaitDone() {
	<-w.done
	close(w.done)
}

func (w *Workflow) interruptDone() {
	w.alreadyDone = true
	w.doneOnce.Do(func() {
		w.done <- struct{}{}
	})
}

type Runnable interface {
	Run(i interface{})
}
