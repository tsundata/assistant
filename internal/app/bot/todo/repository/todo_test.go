package repository

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"testing"
)

func TestTodoRepository_CreateTodo(t *testing.T) {
	sto, err := CreateTodoRepository(enum.Todo)
	if err != nil {
		t.Fatalf("create todo Repository error, %+v", err)
	}
	type args struct {
		todo *pb.Todo
	}
	tests := []struct {
		name    string
		r       TodoRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{todo: &pb.Todo{Content: "test"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateTodo(context.Background(), tt.args.todo)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTodoRepository_ListTodos(t *testing.T) {
	sto, err := CreateTodoRepository(enum.Todo)
	if err != nil {
		t.Fatalf("create todo Repository error, %+v", err)
	}
	tests := []struct {
		name    string
		r       TodoRepository
		wantErr bool
	}{
		{"case1", sto, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListTodos(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.ListTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTodoRepository_ListRemindTodos(t *testing.T) {
	sto, err := CreateTodoRepository(enum.Todo)
	if err != nil {
		t.Fatalf("create todo Repository error, %+v", err)
	}
	tests := []struct {
		name    string
		r       TodoRepository
		wantErr bool
	}{
		{"case1", sto, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListRemindTodos(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.ListTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTodoRepository_GetTodo(t *testing.T) {
	sto, err := CreateTodoRepository(enum.Todo)
	if err != nil {
		t.Fatalf("create todo Repository error, %+v", err)
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       TodoRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{id: 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetTodo(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.GetTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTodoRepository_CompleteTodo(t *testing.T) {
	sto, err := CreateTodoRepository(enum.Todo)
	if err != nil {
		t.Fatalf("create todo Repository error, %+v", err)
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       TodoRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{id: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.CompleteTodo(context.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.CompleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoRepository_UpdateTodo(t *testing.T) {
	sto, err := CreateTodoRepository(enum.Todo)
	if err != nil {
		t.Fatalf("create todo Repository error, %+v", err)
	}
	type args struct {
		todo *pb.Todo
	}
	tests := []struct {
		name    string
		r       TodoRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{todo: &pb.Todo{Id: 1448946120695771136, Content: "test update"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.UpdateTodo(context.Background(), tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoRepository_DeleteTodo(t *testing.T) {
	sto, err := CreateTodoRepository(enum.Todo)
	if err != nil {
		t.Fatalf("create todo Repository error, %+v", err)
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       TodoRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{id: 100}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.DeleteTodo(context.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
