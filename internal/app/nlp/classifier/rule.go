package classifier

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/app"
	"strings"
)



type Rule struct {
	Format string
}

func NewRule() *Rule {
	return &Rule{}
}

func (r *Rule) Do(check string) (enum.RoleAttr, error) {
	part := strings.Split(r.Format, ">")
	if len(part) != 2 {
		return "", errors.New("error rule part")
	}
	attr := strings.TrimSpace(part[1])
	if attr == "" {
		return "", errors.New("error rule attr")
	}
	word := strings.TrimSpace(part[0])
	words := strings.Split(word, "|")
	if len(words) <= 0 {
		return "", errors.New("error rule words")
	}
	for _, item := range words {
		if strings.Contains(check, item) {
			return toRoleAttr(attr)
		}
	}
	return "", app.ErrInvalidParameter
}

func toRoleAttr(shortAttr string) (enum.RoleAttr, error) {
	switch enum.AttrShort(shortAttr) {
	case enum.StrengthShort:
		return enum.StrengthAttr, nil
	case enum.CultureShort:
		return enum.CultureAttr, nil
	case enum.EnvironmentShort:
		return enum.EnvironmentAttr, nil
	case enum.CharismaShort:
		return enum.CharismaAttr, nil
	case enum.TalentShort:
		return enum.TalentAttr, nil
	case enum.IntellectShort:
		return enum.IntellectAttr, nil
	default:
		return "", errors.New("error role attr")
	}
}
