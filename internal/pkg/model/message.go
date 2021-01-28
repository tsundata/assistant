package model

import (
	"regexp"
	"strings"
)

const (
	EventTypeMessage = "message"
	EventTypeNotice  = "notice"
	EventTypeRequest = "request"
)

const (
	MessageTypeText     = "text"
	MessageTypeAt       = "at"
	MessageTypeAudio    = "audio"
	MessageTypeImage    = "image"
	MessageTypeFile     = "file"
	MessageTypeLocation = "location"
	MessageTypeVideo    = "video"
	MessageTypeLink     = "link"
	MessageTypeContact  = "contact"
	MessageTypeGroup    = "group"
	MessageTypeRich     = "rich"
	MessageTypeAction   = "action"
	MessageTypeScript   = "script"
)

const (
	MessageScriptOfJavascript = "javascript"
	MessageScriptOfFlowscript = "flowscript"
	MessageScriptOfUndefined  = "undefined"
)

const (
	PlatformSlack    = "slack"
	PlatformTelegram = "telegram"
	PlatformDiscord  = "discord"
)

type Message struct {
	UUID string
	Type string      `json:"type"`
	Text string      `json:"text"`
	Time string      `json:"time"`
	Data interface{} `json:"data"`
}

type MessageAt struct {
	UserID         string
	UserTID        string
	UserName       string
	UserRemarkName string
	User           interface{}
}

type MessageImage struct {
	Url     string
	Path    string
	Data    string
	MediaID string
}

type MessageAudio struct {
	Url     string
	Path    string
	Data    string
	MediaID string
}

type MessageVideo struct {
	Url     string
	Path    string
	Data    string
	MediaID string
}

type MessageFile struct {
	Url     string
	Path    string
	Data    string
	MediaID string
}

type MessageLink struct {
	Url     string
	Title   string
	Content string
	Image   string
}

type MessageLocation struct {
	Latitude    float64
	Longitude   float64
	Description string
}

type MessageContact struct {
	UserID   string
	UserTID  string
	UserName string
}

type MessageGroup struct {
	GroupID   string
	GroupTID  string
	GroupName string
}

type MessageRich struct {
	Url         string
	Description string
}

type MessageAction struct {
	Action string
	Cron   string
}

type MessageScript struct {
	Type string
	Code string
}

func IsMessageOfScript(text string) bool {
	lines := strings.Split(text, "\n")
	if len(lines) >= 1 {
		re := regexp.MustCompile(`^#!script:\w+$`)
		return re.MatchString(strings.TrimSpace(lines[0]))
	}
	return false
}

func IsMessageOfAction(text string) bool {
	lines := strings.Split(text, "\n")
	if len(lines) >= 1 {
		re := regexp.MustCompile(`^#!action$`)
		return re.MatchString(strings.TrimSpace(lines[0]))
	}
	return false
}

func MessageScriptKind(text string) string {
	if !IsMessageOfScript(text) {
		return MessageScriptOfUndefined
	}

	lines := strings.Split(text, "\n")
	if len(lines) >= 1 {
		return strings.ReplaceAll(strings.TrimSpace(lines[0]), "#!script:", "")
	}
	return MessageScriptOfUndefined
}
