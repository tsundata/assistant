package classifier

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/model"
	"github.com/tsundata/assistant/internal/pkg/config"
	"strings"
)

type Classifier struct {
	rules []Rule
}

func NewClassifier() *Classifier {
	return &Classifier{}
}

func (c *Classifier) SetRules(data string) error {

	rules := strings.Split(data, "\n")

	c.rules = []Rule{}
	for _, rule := range rules {
		if rule == "" {
			continue
		}
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

func ReadRulesConfig(conf *config.AppConfig) (string, error) {
	return conf.GetConfig("classifier")
}
