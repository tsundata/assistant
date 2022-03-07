package agent

import (
	"context"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

const HourLayout = "2006-01-02 15:04"

func TodoRemind(ctx context.Context, comp component.Component) []result.Result {
	//if comp.Todo() == nil {
	//	return []result.Result{result.EmptyResult()}
	//}
	//reply, err := comp.Todo().GetRemindTodos(ctx, &pb.TodoRequest{})
	//if err != nil {
	//	comp.GetLogger().Error(err)
	//	return []result.Result{result.ErrorResult(err)}
	//}
	//
	//if comp.GetRedis() == nil {
	//	return []result.Result{result.EmptyResult()}
	//}
	//
	//var res []result.Result
	//for _, todo := range reply.GetTodos() {
	//	remindKey := fmt.Sprintf("cron:todo_remind:%d:last_remind_at", todo.Id)
	//	remindAt := time.Unix(todo.RemindAt, 0).Format(HourLayout)
	//	if remindAt == time.Now().Format(HourLayout) {
	//		res = append(res, result.MessageResult(fmt.Sprintf("Todo #%d Remind: %s %s", todo.Id, todo.GetContent(), remindAt)))
	//		comp.GetRedis().Set(ctx, remindKey, time.Now().Format(HourLayout), redis.KeepTTL)
	//		continue
	//	}
	//
	//	if todo.RepeatMethod != "" {
	//		// RepeatEndAt
	//		if todo.RepeatEndAt != 0 {
	//			endAt := time.Unix(todo.RepeatEndAt, 0)
	//			if endAt.Before(time.Now()) {
	//				continue
	//			}
	//		}
	//
	//		lastRemindAt, _ := comp.GetRedis().Get(ctx, remindKey).Result()
	//		nowTime := time.Now().Format(HourLayout)
	//
	//		isRemind := false
	//		switch todo.RepeatMethod {
	//		case enum.RepeatDaily:
	//			isRemind, err = util.IsDaily(remindAt, lastRemindAt, nowTime)
	//		case enum.RepeatWeekly:
	//			isRemind, err = util.IsWeekly(remindAt, lastRemindAt, nowTime)
	//		case enum.RepeatMonthly:
	//			isRemind, err = util.IsMonthly(remindAt, lastRemindAt, nowTime)
	//		case enum.RepeatAnnually:
	//			isRemind, err = util.IsAnnually(remindAt, lastRemindAt, nowTime)
	//		}
	//		if err != nil {
	//			continue
	//		}
	//		if isRemind {
	//			res = append(res, result.MessageResult(fmt.Sprintf("Todo #%d Remind: %s %s", todo.Id, todo.GetContent(), remindAt)))
	//			comp.GetRedis().Set(ctx, remindKey, time.Now().Format(HourLayout), redis.KeepTTL)
	//			continue
	//		}
	//	}
	//}
	//
	//return res
	return nil
}
