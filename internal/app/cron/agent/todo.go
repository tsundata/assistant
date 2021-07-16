package agent

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/util"
	"time"
)

const HourLayout = "2006-01-02 15:04"

func TodoRemind(ctx context.Context, comp rulebot.IComponent) []result.Result {
	if comp.Todo() == nil {
		return []result.Result{result.EmptyResult()}
	}
	reply, err := comp.Todo().GetRemindTodos(ctx, &pb.TodoRequest{})
	if err != nil {
		comp.GetLogger().Error(err)
		return []result.Result{result.ErrorResult(err)}
	}

	if comp.GetRedis() == nil {
		return []result.Result{result.EmptyResult()}
	}

	var res []result.Result
	for _, todo := range reply.GetTodos() {
		remindKey := fmt.Sprintf("cron:todo_remind:%d:last_remind_at", todo.Id)
		if todo.RemindAt == time.Now().Format(HourLayout) {
			res = append(res, result.MessageResult(fmt.Sprintf("Todo #%d Remind: %s %s", todo.Id, todo.GetContent(), todo.RemindAt)))
			comp.GetRedis().Set(ctx, remindKey, time.Now().Format(HourLayout), redis.KeepTTL)
			continue
		}

		if todo.RepeatMethod != "" {
			// RepeatEndAt
			if todo.RepeatEndAt != "" {
				endAt, _ := time.ParseInLocation(HourLayout, todo.RepeatEndAt, time.Local)
				if endAt.Before(time.Now()) {
					continue
				}
			}

			lastRemindAt, _ := comp.GetRedis().Get(ctx, remindKey).Result()
			nowTime := time.Now().Format(HourLayout)

			isRemind := false
			switch todo.RepeatMethod {
			case enum.RepeatDaily:
				isRemind, err = util.IsDaily(todo.RemindAt, lastRemindAt, nowTime)
			case enum.RepeatWeekly:
				isRemind, err = util.IsWeekly(todo.RemindAt, lastRemindAt, nowTime)
			case enum.RepeatMonthly:
				isRemind, err = util.IsMonthly(todo.RemindAt, lastRemindAt, nowTime)
			case enum.RepeatAnnually:
				isRemind, err = util.IsAnnually(todo.RemindAt, lastRemindAt, nowTime)
			}
			if err != nil {
				continue
			}
			if isRemind {
				res = append(res, result.MessageResult(fmt.Sprintf("Todo #%d Remind: %s %s", todo.Id, todo.GetContent(), todo.RemindAt)))
				comp.GetRedis().Set(ctx, remindKey, time.Now().Format(HourLayout), redis.KeepTTL)
				continue
			}
		}
	}

	return res
}
