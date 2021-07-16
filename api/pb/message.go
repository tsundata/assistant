package pb

import (
	"regexp"
	"strings"
)

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
