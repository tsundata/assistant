package agent

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/util"
	"time"
)

func TodoRemind(ctx rulebot.IContext) []result.Result {
	if ctx.Todo() == nil {
		return []result.Result{result.EmptyResult()}
	}
	ctxB := context.Background()
	reply, err := ctx.Todo().GetRemindTodos(ctxB, &pb.TodoRequest{})
	if err != nil {
		ctx.GetLogger().Error(err)
		return []result.Result{result.ErrorResult(err)}
	}

	if ctx.GetRedis() == nil {
		return []result.Result{result.EmptyResult()}
	}

	var res []result.Result
	for _, todo := range reply.GetTodos() {
		remindKey := fmt.Sprintf("cron:todo_remind:%d:last_remind_at", todo.Id)
		if todo.RemindAt == time.Now().Format("2006-01-02 15:04") {
			res = append(res, result.MessageResult(fmt.Sprintf("Todo #%d Remind: %s %s", todo.Id, todo.GetContent(), todo.RemindAt)))
			ctx.GetRedis().Set(ctxB, remindKey, time.Now().Format("2006-01-02 15:04"), redis.KeepTTL)
			continue
		}

		if todo.RepeatMethod != "" {
			lastRemindAt, _ := ctx.GetRedis().Get(ctxB, fmt.Sprintf("cron:todo_remind:%d:last_remind_at", todo.Id)).Result()
			nowTime := time.Now().Format("2006-01-02 15:04")

			isRemind := false
			switch todo.RepeatMethod {
			case model.RepeatDaily:
				isRemind, err = util.IsDaily(todo.RemindAt, lastRemindAt, nowTime)
			case model.RepeatWeekly:
				isRemind, err = util.IsWeekly(todo.RemindAt, lastRemindAt, nowTime)
			case model.RepeatMonthly:
				isRemind, err = util.IsMonthly(todo.RemindAt, lastRemindAt, nowTime)
			case model.RepeatAnnually:
				isRemind, err = util.IsAnnually(todo.RemindAt, lastRemindAt, nowTime)
			}
			if err != nil {
				continue
			}
			if isRemind {
				// todo RepeatEndAt
				res = append(res, result.MessageResult(fmt.Sprintf("Todo #%d Remind: %s %s", todo.Id, todo.GetContent(), todo.RemindAt)))
				ctx.GetRedis().Set(ctxB, remindKey, time.Now().Format("2006-01-02 15:04"), redis.KeepTTL)
				continue
			}
		}
	}

	return res
}
