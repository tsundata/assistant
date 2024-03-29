package agent

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"github.com/tsundata/assistant/mock"
	"testing"
	"time"
)

func TestTodoRemind1(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	randId, _ := rand.Read(nil)
	id := int64(randId)
	clear(t, id)

	todo := mock.NewMockTodoSvcClient(ctl)
	gomock.InOrder(
		todo.EXPECT().
			GetRemindTodos(gomock.Any(), gomock.Any()).
			Return(&pb.TodosReply{Todos: []*pb.Todo{
				{
					Id:             id,
					Content:        "test",
					Priority:       1,
					IsRemindAtTime: true,
					RemindAt:       time.Now().Unix(),
					RepeatMethod:   "daily",
					RepeatRule:     "",
					RepeatEndAt:    time.Now().Add(7 * 24 * time.Hour).Unix(),
					Remark:         "",
					Complete:       false,
				},
			}}, nil),
	)

	rdb, err := vendors.CreateRedisClient(enum.Cron)
	if err != nil {
		t.Fatal(err)
	}

	comp := component.MockComponent(rdb)

	type args struct {
		comp component.Component
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{comp},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TodoRemind(context.Background(), tt.args.comp); len(got) != tt.want {
				t.Errorf("TodoRemind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRemind2(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	randId, _ := rand.Read(nil)
	id := int64(randId)
	clear(t, id)

	todo := mock.NewMockTodoSvcClient(ctl)
	gomock.InOrder(
		todo.EXPECT().
			GetRemindTodos(gomock.Any(), gomock.Any()).
			Return(&pb.TodosReply{Todos: []*pb.Todo{
				{
					Id:             id,
					Content:        "test",
					Priority:       1,
					IsRemindAtTime: true,
					RemindAt:       time.Now().Add(-time.Minute).Unix(),
					RepeatMethod:   "daily",
					RepeatRule:     "",
					RepeatEndAt:    time.Now().Add(7 * 24 * time.Hour).Unix(),
					Remark:         "",
					Complete:       false,
				},
			}}, nil),
	)

	rdb, err := vendors.CreateRedisClient(enum.Cron)
	if err != nil {
		t.Fatal(err)
	}

	comp := component.MockComponent(rdb)

	type args struct {
		comp component.Component
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{comp},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TodoRemind(context.Background(), tt.args.comp); len(got) != tt.want {
				t.Errorf("TodoRemind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRemind3(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	randId, _ := rand.Read(nil)
	id := int64(randId)
	clear(t, id)

	todo := mock.NewMockTodoSvcClient(ctl)
	gomock.InOrder(
		todo.EXPECT().
			GetRemindTodos(gomock.Any(), gomock.Any()).
			Return(&pb.TodosReply{Todos: []*pb.Todo{
				{
					Id:             id,
					Content:        "test",
					Priority:       1,
					IsRemindAtTime: true,
					RemindAt:       time.Now().Add(-3 * 24 * time.Hour).Unix(),
					RepeatMethod:   "daily",
					RepeatRule:     "",
					RepeatEndAt:    time.Now().Add(-24 * time.Hour).Unix(),
					Remark:         "",
					Complete:       false,
				},
			}}, nil),
	)

	rdb, err := vendors.CreateRedisClient(enum.Cron)
	if err != nil {
		t.Fatal(err)
	}

	comp := component.MockComponent(rdb)

	type args struct {
		comp component.Component
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{comp},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TodoRemind(context.Background(), tt.args.comp); len(got) != tt.want {
				t.Errorf("TodoRemind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRemind4(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	randId, _ := rand.Read(nil)
	id := int64(randId)
	clear(t, id)

	todo := mock.NewMockTodoSvcClient(ctl)
	gomock.InOrder(
		todo.EXPECT().
			GetRemindTodos(gomock.Any(), gomock.Any()).
			Return(&pb.TodosReply{Todos: []*pb.Todo{
				{
					Id:             id,
					Content:        "test",
					Priority:       1,
					IsRemindAtTime: true,
					RemindAt:       time.Now().Add(-3 * 24 * time.Hour).Unix(),
					RepeatMethod:   "daily",
					RepeatRule:     "",
					RepeatEndAt:    time.Now().Add(24 * time.Hour).Unix(),
					Remark:         "",
					Complete:       false,
				},
			}}, nil),
	)

	rdb, err := vendors.CreateRedisClient(enum.Cron)
	if err != nil {
		t.Fatal(err)
	}

	comp := component.MockComponent(rdb)

	type args struct {
		comp component.Component
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{comp},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TodoRemind(context.Background(), tt.args.comp); len(got) != tt.want {
				t.Errorf("TodoRemind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRemind5(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	randId, _ := rand.Read(nil)
	id := int64(randId)
	clear(t, id)
	now := time.Now()

	todo := mock.NewMockTodoSvcClient(ctl)
	gomock.InOrder(
		todo.EXPECT().
			GetRemindTodos(gomock.Any(), gomock.Any()).
			Return(&pb.TodosReply{Todos: []*pb.Todo{
				{
					Id:             id,
					Content:        "test",
					Priority:       1,
					IsRemindAtTime: true,
					RemindAt:       now.Add(-3 * 24 * time.Hour).Unix(),
					RepeatMethod:   "daily",
					RepeatRule:     "",
					RepeatEndAt:    now.Add(24 * time.Hour).Unix(),
					Remark:         "",
					Complete:       false,
				},
			}}, nil),
		todo.EXPECT().
			GetRemindTodos(gomock.Any(), gomock.Any()).
			Return(&pb.TodosReply{Todos: []*pb.Todo{
				{
					Id:             id,
					Content:        "test",
					Priority:       1,
					IsRemindAtTime: true,
					RemindAt:       now.Add(-3 * 24 * time.Hour).Unix(),
					RepeatMethod:   "daily",
					RepeatRule:     "",
					RepeatEndAt:    now.Add(24 * time.Hour).Unix(),
					Remark:         "",
					Complete:       false,
				},
			}}, nil),
	)

	rdb, err := vendors.CreateRedisClient(enum.Cron)
	if err != nil {
		t.Fatal(err)
	}

	comp1 := component.MockComponent(rdb)
	comp2 := component.MockComponent(rdb)

	type args struct {
		comp component.Component
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{comp1},
			1,
		},
		{
			"case2",
			args{comp2},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TodoRemind(context.Background(), tt.args.comp); len(got) != tt.want {
				t.Errorf("TodoRemind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func clear(t *testing.T, id int64) { //nolint
	rdb, err := vendors.CreateRedisClient(enum.Cron)
	if err != nil {
		t.Fatal(err)
	}
	remindKey := fmt.Sprintf("cron:todo_remind:%d:last_remind_at", id)
	rdb.Del(context.Background(), remindKey)
}
