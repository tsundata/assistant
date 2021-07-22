package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/mock"
	"reflect"
	"testing"
)

func TestTodo_CreateTodo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	nats, err := event.CreateNats(enum.Todo)
	if err != nil {
		t.Fatal(err)
	}
	bus := event.NewNatsBus(nats, nil)

	repo := mock.NewMockTodoRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().CreateTodo(gomock.Any()).Return(int64(1), nil),
	)

	s := NewTodo(bus, nil, repo)

	type args struct {
		in0     context.Context
		payload *pb.TodoRequest
	}
	tests := []struct {
		name    string
		s       *Todo
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{"case1", s, args{context.Background(), &pb.TodoRequest{Todo: &pb.Todo{Content: "test"}}}, &pb.StateReply{State: true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateTodo(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Todo.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Todo.CreateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodo_GetTodo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockTodoRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetTodo(gomock.Any()).Return(pb.Todo{Content: "test"}, nil),
	)

	s := NewTodo(nil, nil, repo)

	type args struct {
		in0     context.Context
		payload *pb.TodoRequest
	}
	tests := []struct {
		name    string
		s       *Todo
		args    args
		want    *pb.TodoReply
		wantErr bool
	}{
		{"case1", s, args{context.Background(),
			&pb.TodoRequest{Todo: &pb.Todo{Id: 1}}},
			&pb.TodoReply{Todo: &pb.Todo{Content: "test"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetTodo(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Todo.GetTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Todo.GetTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodo_GetTodos(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockTodoRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().ListTodos().Return([]pb.Todo{{Content: "test"}}, nil),
	)

	s := NewTodo(nil, nil, repo)

	type args struct {
		in0 context.Context
		in1 *pb.TodoRequest
	}
	tests := []struct {
		name    string
		s       *Todo
		args    args
		want    *pb.TodosReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TodoRequest{}},
			&pb.TodosReply{Todos: []*pb.Todo{
				{Content: "test"},
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetTodos(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Todo.GetTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && len(got.Todos) != len(tt.want.Todos) {
				t.Errorf("Todo.GetTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodo_DeleteTodo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockTodoRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().DeleteTodo(gomock.Any()).Return(nil),
	)

	s := NewTodo(nil, nil, repo)

	type args struct {
		in0     context.Context
		payload *pb.TodoRequest
	}
	tests := []struct {
		name    string
		s       *Todo
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TodoRequest{Todo: &pb.Todo{Id: 1}}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DeleteTodo(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Todo.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Todo.DeleteTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodo_UpdateTodo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockTodoRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().UpdateTodo(gomock.Any()).Return(nil),
	)

	s := NewTodo(nil, nil, repo)

	type args struct {
		in0     context.Context
		payload *pb.TodoRequest
	}
	tests := []struct {
		name    string
		s       *Todo
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TodoRequest{Todo: &pb.Todo{Id: 1, Content: "test update"}}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateTodo(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Todo.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Todo.UpdateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodo_CompleteTodo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockTodoRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().CompleteTodo(gomock.Any()).Return(nil),
	)

	s := NewTodo(nil, nil, repo)

	type args struct {
		in0     context.Context
		payload *pb.TodoRequest
	}
	tests := []struct {
		name    string
		s       *Todo
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TodoRequest{Todo: &pb.Todo{Id: 1}}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CompleteTodo(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Todo.CompleteTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Todo.CompleteTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodo_GetRemindTodos(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockTodoRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().ListRemindTodos().Return([]pb.Todo{{Content: "test"}}, nil),
	)

	s := NewTodo(nil, nil, repo)

	type args struct {
		in0 context.Context
		in1 *pb.TodoRequest
	}
	tests := []struct {
		name    string
		s       *Todo
		args    args
		want    int
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TodoRequest{}},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetRemindTodos(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Todo.GetRemindTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && len(got.Todos) != tt.want {
				t.Errorf("Todo.GetRemindTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}
