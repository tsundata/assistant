package classifier

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/model"
	"strings"
)

type Classifier struct {
	conf  *config.AppConfig
	rules []Rule
}

func NewClassifier(conf *config.AppConfig) *Classifier {
	return &Classifier{conf: conf}
}

func (c *Classifier) LoadRule() error {
	s, err := c.conf.GetConfig("classifier")
	if err != nil {
		return err
	}
	rules := strings.Split(s, "\n")

	c.rules = []Rule{}
	for _, rule := range rules {
		c.rules = append(c.rules, Rule{Format: rule})
	}
	return nil
}

func (c *Classifier) Do(check string) (model.RoleAttr, error) {
	for _, rule := range c.rules {
		res, err := rule.Do(check)
		if err != nil && !errors.Is(err, ErrEmpty) {
			return "", err
		}
		if res != "" {
			return res, nil
		}
	}
	return "", ErrEmpty
}
