package model

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

type Message struct {
	ID          int `gorm:"primaryKey"`
	UUID        string
	Remotes     RemoteType
	Type        MessageType
	ChannelID   string `gorm:"channel_id"`
	ChannelName string `gorm:"channel_name"`
	Content     string
	Attributes  map[string]string `gorm:"-"`
	Vars        map[string]string `gorm:"-"`
	SourceLink  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MessageType int

const (
	MessageTypeText MessageType = iota
	MessageTypeVoice
	MessageTypeImage
	MessageTypeFile
	MessageTypeLocation
	MessageTypeVideo
	MessageTypeUrl
	MessageTypeAction
	MessageTypeScript
)

const (
	MessageScriptOfJavascript = "javascript"
	MessageScriptOfDSL        = "dsl"
	MessageScriptOfUndefined  = "undefined"
)

type RemoteType int

const (
	RemoteSlack RemoteType = iota
	RemoteDiscord
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

// MessageTimestamp timestamps the message
func MessageTimestamp() int64 {
	return time.Now().Unix()
}

func NewMessage() Message {
	uuid, _ := GenerateMessageUUID()
	return Message{
		UUID:       uuid,
		Attributes: make(map[string]string),
		Vars:       make(map[string]string),
	}
}
