package models

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

type Message struct {
	ID                string
	Type              MessageType
	Service           MessageService
	ChannelID         string
	ChannelName       string
	Input             string
	Output            string
	Error             string
	Timestamp         string
	ThreadTimestamp   string
	BotMentioned      bool
	DirectMessageOnly bool
	Debug             bool
	IsEphemeral       bool
	StartTime         int64
	EndTime           int64
	Attributes        map[string]string
	Vars              map[string]string
	OutputToRooms     []string
	OutputToUsers     []string
	Remotes           RemoteType
	SourceLink        string
}

type MessageType int
type MessageService int

const (
	MsgTypeUnknown MessageType = iota
	MsgTypeDirect
	MsgTypeChannel
	MsgTypePrivateChannel
)

// GenerateMessageID generates a random ID for a message
func GenerateMessageID() (string, error) {
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

type RemoteType int

const (
	RemoteSlack RemoteType = iota
	RemoteDiscord
)

func NewMessage() Message {
	uuid, _ := GenerateMessageID()
	return Message{
		ID:            uuid,
		StartTime:     MessageTimestamp(),
		Attributes:    make(map[string]string),
		Vars:          make(map[string]string),
		OutputToRooms: []string{},
		OutputToUsers: []string{},
		Debug:         false,
		IsEphemeral:   false,
	}
}
