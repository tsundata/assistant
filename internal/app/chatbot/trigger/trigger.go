package trigger

import (
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/ctx"
	"sync"
)

type Trigger interface {
	Cond(text string) bool
	Handle(ctx *ctx.Context)
}

func triggers() []Trigger {
	return []Trigger{
		NewUrl(),
		NewEmail(),
		NewTag(),
		NewUser(),
	}
}

func Run(c *ctx.Context, message string) {
	triggers := triggers()
	wg := sync.WaitGroup{}
	for _, item := range triggers {
		wg.Add(1)
		go func(t Trigger) {
			defer wg.Done()
			if t.Cond(message) {
				t.Handle(c)
			}
		}(item)
	}
	wg.Wait()
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
