package utils

import (
	"github.com/tsundata/assistant/internal/pkg/model"
	"net"
	"regexp"
	"strings"
)

func IsUrl(text string) bool {
	re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
	return re.MatchString(text)
}

func IsIPv4(host string) bool {
	return net.ParseIP(host) != nil
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
		return model.MessageScriptOfUndefined
	}

	lines := strings.Split(text, "\n")
	if len(lines) >= 1 {
		return strings.Replace(strings.TrimSpace(lines[0]), "#!script:", "", -1)
	}
	return model.MessageScriptOfUndefined
}
