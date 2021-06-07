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
	MessageTypeScript   = "script"
)

const (
	MessageScriptOfJavascript = "javascript"
	MessageScriptOfUndefined  = "undefined"
)

const (
	PlatformSlack    = "slack"
	PlatformTelegram = "telegram"
	PlatformDiscord  = "discord"
)

type Message struct {
	ID   int       `db:"id"`
	UUID string    `db:"uuid"`
	Type string    `db:"type"`
	Text string    `db:"text"`
	Time time.Time `db:"time"`
}

func (m *Message) IsMessageOfScript() bool {
	lines := strings.Split(m.Text, "\n")
	if len(lines) >= 1 {
		re := regexp.MustCompile(`^#!/usr/bin/env\s+\w+\s*$`)
		return re.MatchString(strings.TrimSpace(lines[0]))
	}
	return false
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
	re := regexp.MustCompile(`^#!action\s*$`)
	return re.ReplaceAllString(m.Text, "")
}

func (m *Message) ScriptKind() string {
	if !m.IsMessageOfScript() {
		return MessageScriptOfUndefined
	}

	lines := strings.Split(m.Text, "\n")
	if len(lines) >= 1 {
		return strings.TrimSpace(strings.ReplaceAll(lines[0], "#!/usr/bin/env", ""))
	}
	return MessageScriptOfUndefined
}
