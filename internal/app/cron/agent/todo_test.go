package agent

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"github.com/tsundata/assistant/mock"
	"math/rand"
	"testing"
	"time"

	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func TestTodoRemind1(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	id := rand.Int63()
	clear(t, id)

	todo := mock.NewMockTodoClient(ctl)
	gomock.InOrder(
		todo.EXPECT().
			GetRemindTodos(gomock.Any(), gomock.Any()).
			Return(&pb.TodosReply{Todos: []*pb.TodoItem{
				{
					Id:             id,
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

func TestTodoRemind2(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	id := rand.Int63()
	clear(t, id)

	todo := mock.NewMockTodoClient(ctl)
	gomock.InOrder(
		todo.EXPECT().
			GetRemindTodos(gomock.Any(), gomock.Any()).
			Return(&pb.TodosReply{Todos: []*pb.TodoItem{
				{
					Id:             id,
					Content:        "test",
					Priority:       1,
					IsRemindAtTime: true,
					RemindAt:       time.Now().Add(-time.Minute).Format("2006-01-02 15:04"),
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
			0,
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

func TestTodoRemind3(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	id := rand.Int63()
	clear(t, id)

	todo := mock.NewMockTodoClient(ctl)
	gomock.InOrder(
		todo.EXPECT().
			GetRemindTodos(gomock.Any(), gomock.Any()).
			Return(&pb.TodosReply{Todos: []*pb.TodoItem{
				{
					Id:             id,
					Content:        "test",
					Priority:       1,
					IsRemindAtTime: true,
					RemindAt:       time.Now().Add(- 3 * 24 * time.Hour).Format("2006-01-02 15:04"),
					RepeatMethod:   "daily",
					RepeatRule:     "",
					RepeatEndAt:    time.Now().Add(- 24 * time.Hour).Format("2006-01-02 15:04"),
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
			0,
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

func TestTodoRemind4(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	id := rand.Int63()
	clear(t, id)

	todo := mock.NewMockTodoClient(ctl)
	gomock.InOrder(
		todo.EXPECT().
			GetRemindTodos(gomock.Any(), gomock.Any()).
			Return(&pb.TodosReply{Todos: []*pb.TodoItem{
				{
					Id:             id,
					Content:        "test",
					Priority:       1,
					IsRemindAtTime: true,
					RemindAt:       time.Now().Add(- 3 * 24 * time.Hour).Format("2006-01-02 15:04"),
					RepeatMethod:   "daily",
					RepeatRule:     "",
					RepeatEndAt:    time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04"),
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

func TestTodoRemind5(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	id := rand.Int63()
	clear(t, id)
	now := time.Now()

	todo := mock.NewMockTodoClient(ctl)
	gomock.InOrder(
		todo.EXPECT().
			GetRemindTodos(gomock.Any(), gomock.Any()).
			Return(&pb.TodosReply{Todos: []*pb.TodoItem{
				{
					Id:             id,
					Content:        "test",
					Priority:       1,
					IsRemindAtTime: true,
					RemindAt:       now.Add(- 3 * 24 * time.Hour).Format("2006-01-02 15:04"),
					RepeatMethod:   "daily",
					RepeatRule:     "",
					RepeatEndAt:    now.Add(24 * time.Hour).Format("2006-01-02 15:04"),
					Remark:         "",
					Complete:       false,
				},
			}}, nil),
		todo.EXPECT().
			GetRemindTodos(gomock.Any(), gomock.Any()).
			Return(&pb.TodosReply{Todos: []*pb.TodoItem{
				{
					Id:             id,
					Content:        "test",
					Priority:       1,
					IsRemindAtTime: true,
					RemindAt:       now.Add(- 3 * 24 * time.Hour).Format("2006-01-02 15:04"),
					RepeatMethod:   "daily",
					RepeatRule:     "",
					RepeatEndAt:    now.Add(24 * time.Hour).Format("2006-01-02 15:04"),
					Remark:         "",
					Complete:       false,
				},
			}}, nil),
	)

	rdb, err := vendors.CreateRedisClient(app.Cron)
	if err != nil {
		t.Fatal(err)
	}

	ctx1 := rulebot.NewContext(nil, rdb, nil, nil,
		nil, nil, nil, nil,
		todo, nil, nil)

	ctx2 := rulebot.NewContext(nil, rdb, nil, nil,
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
			args{ctx1},
			1,
		},
		{
			"case2",
			args{ctx2},
			0,
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

func clear(t *testing.T, id int64) {
	rdb, err := vendors.CreateRedisClient(app.Cron)
	if err != nil {
		t.Fatal(err)
	}
	remindKey := fmt.Sprintf("cron:todo_remind:%d:last_remind_at", id)
	rdb.Del(context.Background(), remindKey)
}
