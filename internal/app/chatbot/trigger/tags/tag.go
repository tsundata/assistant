package tags

import (
	"context"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/ctx"
)

type Tagger interface {
	Handle(ctx context.Context, comp *ctx.Component, text string)
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
