package stage

import (
	"context"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

func Done(_ context.Context, _ component.Component, _ result.Result) result.Result {
	return result.EmptyResult()
}
