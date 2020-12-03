package model

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"
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

type Event struct {
	ID      int
	UUID    string
	Type    string
	Time    time.Time
	Context EventContext
	Data    EventData
}

type EventContext struct {
	Platform   string
	Via        string
	Type       string
	UserID     string
	UserTID    string
	GroupID    string
	GroupTID   string
	DiscussID  string
	DiscussTID string
	Extra      interface{}
}

type EventData struct {
	Type              string
	Message           Message
	SenderID          string
	SenderTID         string
	SenderName        string
	SenderRemarkName  string
	Sender            string
	GroupID           string
	GroupTID          string
	GroupName         string
	GroupRemarkName   string
	Group             string
	DiscussID         string
	DiscussTID        string
	DiscussName       string
	DiscussRemarkName string
	Discuss           string
	SenderRole        string
	Extra             interface{}
}

type Message struct {
	Type string
	Text string
	Data interface{}
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
