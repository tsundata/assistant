package agent

import (
	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"github.com/tsundata/assistant/mock"
	"testing"
	"time"

	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func TestTodoRemind(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	todo := mock.NewMockTodoClient(ctl)
	gomock.InOrder(
		todo.EXPECT().
			GetRemindTodos(gomock.Any(), gomock.Any()).
			Return(&pb.TodosReply{Todos: []*pb.TodoItem{
				{
					Id:             1,
					Content:        "test",
					Priority:       1,
					IsRemindAtTime: true,
					RemindAt:       time.Now().Format("2006-01-02 15:04"),
					RepeatMethod:   "daily",
					RepeatRule:     "",
					RepeatEndAt:    time.Now().Add(7 * 24 * time.Hour).Format("2006-01-02 15:04"),
					Remark:         "",
					Complete:       false,
				},
			}}, nil),
	)

	rdb, err := vendors.CreateRedisClient(app.Cron)
	if err != nil {
		t.Fatal(err)
	}

	ctx := rulebot.NewContext(nil, rdb, nil, nil,
		nil, nil, nil, nil,
		todo, nil, nil)

	type args struct {
		ctx rulebot.IContext
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{ctx},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TodoRemind(tt.args.ctx); len(got) != tt.want {
				t.Errorf("TodoRemind() = %v, want %v", got, tt.want)
			}
		})
	}
}
