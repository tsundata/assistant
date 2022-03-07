package trigger

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"sync"
)

type Trigger interface {
	Cond(message *pb.Message) bool
	Handle(ctx context.Context, comp component.Component)
}

func triggers() []Trigger {
	return []Trigger{
		NewUrl(),
		NewEmail(),
		//NewTag(), todo
	}
}

func Process(ctx context.Context, comp component.Component, message *pb.Message) error {
	triggers := triggers()
	wg := sync.WaitGroup{}
	for _, item := range triggers {
		wg.Add(1)
		go func(t Trigger) {
			defer wg.Done()
			if t.Cond(message) {
				t.Handle(ctx, comp)
			}
		}(item)
	}
	wg.Wait()
	return nil
}

func clear(arr []string) []string {
	keys := make(map[string]struct{})
	var result []string

	for _, item := range arr {
		if item == "" {
			continue
		}
		if _, ok := keys[item]; !ok {
			keys[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}
