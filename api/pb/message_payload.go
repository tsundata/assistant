package pb

import (
	"github.com/tsundata/assistant/api/enum"
)

type MsgPayload interface {
	Type() enum.MessageType
}

type TextMsg struct {
	Text string `json:"text"`
}

func (t TextMsg) Type() enum.MessageType {
	return enum.MessageTypeText
}

type ImageMsg struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Alt    string `json:"alt"`
}

func (i ImageMsg) Type() enum.MessageType {
	return enum.MessageTypeImage
}

type ScriptMsg struct {
	Kind string `json:"kind"`
	Code string `json:"code"`
}

func (a ScriptMsg) Type() enum.MessageType {
	return enum.MessageTypeScript
}

type ActionMsg struct {
	Id     int64         `json:"id"`
	Option []interface{} `json:"option"`
}

func (a ActionMsg) Type() enum.MessageType {
	return enum.MessageTypeAction
}

type LinkMsg struct {
	Title string `json:"title"`
	Cover string `json:"cover"`
	Url   string `json:"url"`
}

func (a LinkMsg) Type() enum.MessageType {
	return enum.MessageTypeLink
}

func MockMsgPayload() []MsgPayload {
	return []MsgPayload{
		TextMsg{Text: "test"},
		ImageMsg{
			Src:    "https://chatscope.io/storybook/react/static/media/zoe.e31a4ff8.svg",
			Width:  100,
			Height: 100,
			Alt:    "Avatar",
		},
		ScriptMsg{
			Kind: enum.ActionScript,
			Code: "#!action\necho 1",
		},
		LinkMsg{
			Title: "test",
			Cover: "https://chatscope.io/storybook/react/static/media/zoe.e31a4ff8.svg",
			Url:   "https://test.dev",
		},
		ActionMsg{
			Id: 1,
			Option: []interface{}{
				"true", "false",
			},
		},
	}
}
