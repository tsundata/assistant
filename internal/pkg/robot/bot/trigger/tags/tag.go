package tags

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type Tagger interface {
	Handle(ctx context.Context, comp component.Component, text string)
}

func Tags() map[string]Tagger {
	return map[string]Tagger{
		"issue":   NewIssue(),
		"project": NewProject(),
		"todo":    NewTodo(),
	}
}

func MapTagger(text string) Tagger {
	m := Tags()
	if t, ok := m[text]; ok {
		return t
	}
	return nil
}
