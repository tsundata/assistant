package model

import (
	"crypto/rand"
	"fmt"
	"io"
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

// GenerateMessageID generates a random ID for a message
func GenerateMessageUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

type Message struct {
	UUID string
	Type string      `json:"type"`
	Text string      `json:"text"`
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
		return strings.Replace(strings.TrimSpace(lines[0]), "#!script:", "", -1)
	}
	return MessageScriptOfUndefined
}
