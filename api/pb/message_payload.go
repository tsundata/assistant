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

type FileMsg struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

func (i FileMsg) Type() enum.MessageType {
	return enum.MessageTypeFile
}

type VideoMsg struct {
	Src      string  `json:"src"`
	Width    int     `json:"width"`
	Height   int     `json:"height"`
	Alt      string  `json:"alt"`
	Duration float64 `json:"duration"`
}

func (i VideoMsg) Type() enum.MessageType {
	return enum.MessageTypeVideo
}

type AudioMsg struct {
	Src      string  `json:"src"`
	Alt      string  `json:"alt"`
	Duration float64 `json:"duration"`
}

func (i AudioMsg) Type() enum.MessageType {
	return enum.MessageTypeAudio
}

type ScriptMsg struct {
	Kind string `json:"kind"`
	Code string `json:"code"`
}

func (a ScriptMsg) Type() enum.MessageType {
	return enum.MessageTypeScript
}

type ActionMsg struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Option []string `json:"option"`
	Value  string   `json:"value"`
}

func (a ActionMsg) Type() enum.MessageType {
	return enum.MessageTypeAction
}

type FormMsg struct {
	ID    string      `json:"id"`
	Title string      `json:"title"`
	Field []FormField `json:"field"`
}

func (a FormMsg) Type() enum.MessageType {
	return enum.MessageTypeForm
}

type FormField struct {
	Key      string      `json:"key"`
	Type     string      `json:"type"`
	Required bool        `json:"required"`
	Value    interface{} `json:"value"`
}

type LinkMsg struct {
	Title string `json:"title"`
	Cover string `json:"cover"`
	Url   string `json:"url"`
}

func (a LinkMsg) Type() enum.MessageType {
	return enum.MessageTypeLink
}

type LocationMsg struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Address   string  `json:"address"`
}

func (a LocationMsg) Type() enum.MessageType {
	return enum.MessageTypeLocation
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
			ID:    "demo",
			Title: "demo?",
			Option: []string{
				"true",
				"false",
			},
		},
		FormMsg{
			ID:    "demo",
			Title: "demo?",
			Field: []FormField{
				{
					Key:      "title",
					Type:     "string",
					Required: true,
					Value:    nil,
				},
			},
		},
		LocationMsg{
			Longitude: 112.5,
			Latitude:  52.1,
		},
		FileMsg{
			Src: "test.go",
			Alt: "Test",
		},
		VideoMsg{
			Src:      "test.mp4",
			Width:    1080,
			Height:   720,
			Alt:      "Test",
			Duration: 120,
		},
		AudioMsg{
			Src:      "test.mp3",
			Alt:      "Test",
			Duration: 50,
		},
	}
}
