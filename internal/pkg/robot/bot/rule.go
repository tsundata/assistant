package bot

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type Metadata struct {
	Name       string
	Identifier string
	Detail     string
	Avatar     string
}

type FieldItemType string

const (
	FieldItemTypeString FieldItemType = "string"
	FieldItemTypeInt    FieldItemType = "int"
	FieldItemTypeFloat  FieldItemType = "float"
	FieldItemTypeBool   FieldItemType = "bool"
)

type FieldItem struct {
	Key      string        `json:"key"`
	Type     FieldItemType `json:"type"`
	Required bool          `json:"required"`
	Value    interface{}   `json:"value"`
}

type WorkflowRule struct {
	Plugin  []PluginRule
	RunFunc ActFunc
}

type PluginRule struct {
	Name  string
	Param []interface{}
}

type ActionRule struct {
	ID         string
	Title      string
	OptionFunc map[string]ActFunc
}

type FormRule struct {
	ID         string
	Title      string
	Field      []FieldItem
	SubmitFunc ActFunc
}

type TagRule struct {
	Tag         string
	TriggerFunc ActFunc
}

type ActFunc func(context.Context, Context, component.Component) []pb.MsgPayload
