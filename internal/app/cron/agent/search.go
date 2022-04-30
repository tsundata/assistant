package agent

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

func SearchMetadata(ctx context.Context, comp component.Component) []result.Result {
	if comp.Middle() == nil {
		return []result.Result{result.EmptyResult()}
	}
	_, err := comp.Middle().CollectMetadata(ctx, &pb.TextRequest{})
	if err != nil {
		comp.GetLogger().Error(err)
		return []result.Result{result.ErrorResult(err)}
	}
	return []result.Result{result.EmptyResult()}
}
