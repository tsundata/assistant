package classifier

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/app"
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

func (c *Classifier) Do(check string) (enum.RoleAttr, error) {
	for _, rule := range c.rules {
		res, err := rule.Do(check)
		if err != nil && !errors.Is(err, app.ErrInvalidParameter) {
			return "", err
		}
		if res != "" {
			return res, nil
		}
	}
	return "", app.ErrInvalidParameter
}

func ReadRulesConfig(conf *config.AppConfig) (string, error) {
	return conf.GetConfig(context.Background(), "classifier")
}
