package trigger

import "github.com/tsundata/assistant/internal/app/message/trigger/ctx"

type Trigger interface {
	Cond(text string) bool
	Handle(ctx *ctx.Context)
}

func Triggers() []Trigger {
	return []Trigger{
		NewUrl(),
		NewEmail(),
		NewTag(),
		NewUser(),
	}
}

func clear(arr []string) []string {
	keys := make(map[string]struct{})
	var result []string

	for _, item := range arr {
		if item == "" {
			continue
		}
		if _, value := keys[item]; !value {
			keys[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}
