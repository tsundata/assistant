package classifier

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/model"
	"strings"
)

var ErrEmpty = errors.New("empty check")

type Rule struct {
	Format string
}

func NewRule() *Rule {
	return &Rule{}
}

func (r *Rule) Do(check string) (model.RoleAttr, error) {
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
	return "", ErrEmpty
}

func toRoleAttr(shortAttr string) (model.RoleAttr, error) {
	switch model.AttrShort(shortAttr) {
	case model.StrengthShort:
		return model.StrengthAttr, nil
	case model.CultureShort:
		return model.CultureAttr, nil
	case model.EnvironmentShort:
		return model.EnvironmentAttr, nil
	case model.CharismaShort:
		return model.CharismaAttr, nil
	case model.TalentShort:
		return model.TalentAttr, nil
	case model.IntellectShort:
		return model.IntellectAttr, nil
	default:
		return "", errors.New("error role attr")
	}
}
