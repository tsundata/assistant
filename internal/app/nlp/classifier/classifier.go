package classifier

import (
	"github.com/tsundata/assistant/internal/pkg/config"
	"strings"
)

type Classifier struct {
	conf  *config.AppConfig
	rules []Rule
}

func NewClassifier(conf *config.AppConfig) *Classifier {
	return &Classifier{conf: conf}
}

func (c Classifier) LoadRule() error {
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
