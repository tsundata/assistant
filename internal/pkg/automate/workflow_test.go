package automate

import (
	"context"
	"log"
	"testing"
)

type WearSocksAction struct {
}

func (w *WearSocksAction) Run(_ interface{}) {
	log.Println("WearSocksAction")
}

type WearUnderpantsAction struct {
}

func (w *WearUnderpantsAction) Run(_ interface{}) {
	log.Println("WearUnderpantsAction")
}

type ShirtNodeAction struct {
}

func (w *ShirtNodeAction) Run(_ interface{}) {
	log.Println("ShirtNodeAction")
}

type WatchNodeAction struct {
}

func (w *WatchNodeAction) Run(_ interface{}) {
	log.Println("WatchNodeAction")
}

type WearTrouserNodeAction struct {
}

func (w *WearTrouserNodeAction) Run(_ interface{}) {
	log.Println("WearTrouserNodeAction")
}

type WearShoesNodeAction struct {
}

func (w *WearShoesNodeAction) Run(_ interface{}) {
	log.Println("WearShoesNodeAction")
}

type WearCoatNodeAction struct {
}

func (w *WearCoatNodeAction) Run(_ interface{}) {
	log.Println("WearCoatNodeAction")
}

func TestWorkflow(t *testing.T) {
	wf := NewWorkflow()
	// nodes
	UnderpantsNode := NewNode(&WearUnderpantsAction{})
	SocksNode := NewNode(&WearSocksAction{})
	ShirtNode := NewNode(&ShirtNodeAction{})
	WatchNode := NewNode(&WatchNodeAction{})
	TrousersNode := NewNode(&WearTrouserNodeAction{})
	ShoesNode := NewNode(&WearShoesNodeAction{})
	CoatNode := NewNode(&WearCoatNodeAction{})
	// edges
	wf.AddStartNode(UnderpantsNode)
	wf.AddStartNode(SocksNode)
	wf.AddStartNode(ShirtNode)
	wf.AddStartNode(WatchNode)
	wf.AddEdge(UnderpantsNode, TrousersNode)
	wf.AddEdge(TrousersNode, ShoesNode)
	wf.AddEdge(SocksNode, ShoesNode)
	wf.AddEdge(ShirtNode, CoatNode)
	wf.AddEdge(WatchNode, CoatNode)
	wf.ConnectToEnd(ShoesNode)
	wf.ConnectToEnd(CoatNode)
	var completedAction []string
	wf.StartWithContext(context.Background(), completedAction)
	wf.WaitDone()
}
