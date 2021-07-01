package model

import (
	"regexp"
	"strings"
	"time"
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
)

const (
	PlatformSlack    = "slack"
	PlatformTelegram = "telegram"
	PlatformDiscord  = "discord"
)

type Message struct {
	ID        int       `db:"id"`
	UUID      string    `db:"uuid"`
	Type      string    `db:"type"`
	Channel   string    `db:"channel"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}

func (m *Message) IsMessageOfAction() bool {
	lines := strings.Split(m.Text, "\n")
	if len(lines) >= 1 {
		re := regexp.MustCompile(`^#!action\s*$`)
		return re.MatchString(strings.TrimSpace(lines[0]))
	}
	return false
}

func (m *Message) RemoveActionFlag() string {
	re := regexp.MustCompile(`^#!action\s*`)
	return re.ReplaceAllString(m.Text, "")
}
