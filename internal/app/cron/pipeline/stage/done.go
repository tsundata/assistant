package stage

import (
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func Done(_ *rulebot.RuleBot, _ result.Result) result.Result {
	return result.EmptyResult()
}
