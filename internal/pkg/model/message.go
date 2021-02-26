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
	MessageScriptOfFlowscript = "flowscript"
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

func IsMessageOfScript(text string) bool {
	lines := strings.Split(text, "\n")
	if len(lines) >= 1 {
		re := regexp.MustCompile(`^#!/usr/bin/env\s+\w+\s*$`)
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
		return strings.TrimSpace(strings.ReplaceAll(lines[0], "#!/usr/bin/env", ""))
	}
	return MessageScriptOfUndefined
}
