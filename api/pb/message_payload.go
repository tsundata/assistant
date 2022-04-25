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
	Default  interface{} `json:"default"`
	Intro    string      `json:"intro"`
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

type TableMsg struct {
	Title  string          `json:"title"`
	Header []string        `json:"header"`
	Row    [][]interface{} `json:"row"`
}

func (t TableMsg) Type() enum.MessageType {
	return enum.MessageTypeTable
}

type DigitMsg struct {
	Title string `json:"title"`
	Digit int    `json:"digit"`
}

func (a DigitMsg) Type() enum.MessageType {
	return enum.MessageTypeDigit
}
