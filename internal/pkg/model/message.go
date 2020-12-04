package model

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
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
	Context EventContext `gorm:"-"`
	Data    EventData    `gorm:"-"`
	Event   string       `json:"event"`
}

func (e *Event) BeforeCreate(tx *gorm.DB) (err error) {
	d, err := json.Marshal(e.Data)
	if err != nil {
		return
	}
	e.Event = string(d)
	return
}

func (e *Event) AfterFind(tx *gorm.DB) (err error) {
	var d EventData
	err = json.Unmarshal([]byte(e.Event), &d)
	if err != nil {
		return
	}
	e.Data = d
	return
}

type EventContext struct {
	Platform   string `json:"platform"`
	Via        string `json:"via"`
	Type       string `json:"type"`
	UserID     string `json:"user_id"`
	UserTID    string `json:"user_tid"`
	GroupID    string `json:"group_id"`
	GroupTID   string `json:"group_tid"`
	DiscussID  string `json:"discuss_id"`
	DiscussTID string `json:"discuss_tid"`
	Extra      interface{}
}

type EventData struct {
	Type              string      `json:"type"`
	Message           Message     `json:"message"`
	SenderID          string      `json:"sender_id"`
	SenderTID         string      `json:"sender_tid"`
	SenderName        string      `json:"sender_name"`
	SenderRemarkName  string      `json:"sender_remark_name"`
	Sender            string      `json:"sender"`
	GroupID           string      `json:"group_id"`
	GroupTID          string      `json:"group_tid"`
	GroupName         string      `json:"group_name"`
	GroupRemarkName   string      `json:"group_remark_name"`
	Group             string      `json:"group"`
	DiscussID         string      `json:"discuss_id"`
	DiscussTID        string      `json:"discuss_tid"`
	DiscussName       string      `json:"discuss_name"`
	DiscussRemarkName string      `json:"discuss_remark_name"`
	Discuss           string      `json:"discuss"`
	SenderRole        string      `json:"sender_role"`
	Extra             interface{} `json:"extra"`
}

type Message struct {
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
