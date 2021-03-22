package result

import (
	"github.com/tsundata/assistant/internal/pkg/utils"
)

type Kind int

const (
	Done Kind = iota
	Error
	Message
	Url
	Repos
)

type Result struct {
	ID      string
	Kind    Kind
	Content interface{}
}

func id() (i string) {
	i, _ = utils.GenerateUUID()
	return
}

func ErrorResult(err error) Result {
	return Result{
		ID:      id(),
		Kind:    Error,
		Content: err,
	}
}

func MessageResult(text string) Result {
	return Result{
		ID:      id(),
		Kind:    Message,
		Content: text,
	}
}

func EmptyResult() Result {
	return Result{}
}

func DoneResult() Result {
	return Result{Kind: Done}
}
