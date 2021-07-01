package stage

import (
	"encoding/json"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/util"
)

func URL(_ rulebot.IContext, in result.Result) result.Result {
	if in.Kind == result.Url {
		j, err := json.Marshal(in.Content)
		if err != nil {
			return result.ErrorResult(err)
		}
		return result.MessageResult(util.ByteToString(j))
	}
	return result.EmptyResult()
}
