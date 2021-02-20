package rules

import (
	"bytes"
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

type Rule struct {
	Regex        string
	HelpMessage  string
	ParseMessage func(*rulebot.RuleBot, string, []string) []string
}

type regexRuleset struct {
	regexes map[string]*template.Template
	rules   []Rule
}

func (r regexRuleset) Name() string {
	return "Regex Ruleset"
}

func (r regexRuleset) Boot(_ *rulebot.RuleBot) {}

func (r regexRuleset) HelpMessage(b *rulebot.RuleBot, _ model.Message) string {
	botName := b.Name()
	var helpMsg string
	for _, rule := range r.rules {
		var finalRegex bytes.Buffer
		_ = r.regexes[rule.Regex].Execute(&finalRegex, struct{ RobotName string }{botName})

		helpMsg = fmt.Sprintln(helpMsg, finalRegex.String(), "-", rule.HelpMessage)
	}
	return strings.TrimLeftFunc(helpMsg, unicode.IsSpace)
}

func (r regexRuleset) ParseMessage(b *rulebot.RuleBot, in model.Message) []model.Message {
	for _, rule := range r.rules {
		botName := b.Name()
		var finalRegex bytes.Buffer
		if _, ok := r.regexes[rule.Regex]; !ok {
			r.regexes[rule.Regex] = template.Must(template.New(rule.Regex).Parse(rule.Regex))
		}
		_ = r.regexes[rule.Regex].Execute(&finalRegex, struct{ RobotName string }{botName})
		sanitizedRegex := strings.TrimSpace(finalRegex.String())
		re := regexp.MustCompile(sanitizedRegex)
		matched := re.MatchString(in.Text)
		if !matched {
			continue
		}

		args := re.FindStringSubmatch(in.Text)
		if ret := rule.ParseMessage(b, in.Text, args); len(ret) > 0 {
			var retMsgs []model.Message
			for _, m := range ret {
				retMsgs = append(retMsgs, model.Message{
					Text: m,
				})
			}
			return retMsgs
		}
	}

	return []model.Message{}
}

func New(rules []Rule) *regexRuleset {
	r := &regexRuleset{
		regexes: make(map[string]*template.Template),
		rules:   rules,
	}
	for _, rule := range rules {
		r.regexes[rule.Regex] = template.Must(template.New(rule.Regex).Parse(rule.Regex))
	}
	return r
}
