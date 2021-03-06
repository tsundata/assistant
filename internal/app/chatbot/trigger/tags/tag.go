package tags

import "github.com/tsundata/assistant/internal/app/chatbot/trigger/ctx"

type Tagger interface {
	Handle(ctx *ctx.Context, text string)
}

func MapTagger(text string) Tagger {
	m := map[string]Tagger{
		"issue":   NewIssue(),
		"project": NewProject(),
		"todo":    NewTodo(),
	}
	if t, ok := m[text]; ok {
		return t
	}
	return nil
}
