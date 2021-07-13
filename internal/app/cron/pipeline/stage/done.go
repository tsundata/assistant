package stage

import (
	"context"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func Done(_ context.Context,_ rulebot.IComponent, _ result.Result) result.Result {
	return result.EmptyResult()
}
