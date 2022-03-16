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
	Width  string `json:"width"`
	Height string `json:"height"`
	Alt    string `json:"alt"`
}

func (i ImageMsg) Type() enum.MessageType {
	return enum.MessageTypeImage
}

type ActionMsg struct {
	Script string `json:"script"`
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
