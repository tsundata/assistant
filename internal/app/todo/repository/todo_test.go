package repository

import (
	"github.com/tsundata/assistant/internal/pkg/app"
	"testing"

	"github.com/tsundata/assistant/internal/pkg/model"
)

func TestTodoRepository_CreateTodo(t *testing.T) {
	sto, err := CreateTodoRepository(app.Todo)
	if err != nil {
		t.Fatalf("create todo Preposiory error, %+v", err)
	}
	type args struct {
		todo model.Todo
	}
	tests := []struct {
		name    string
		r       TodoRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{todo: model.Todo{Content: "test"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateTodo(tt.args.todo)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlTodoRepository.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTodoRepository_ListTodos(t *testing.T) {
	sto, err := CreateTodoRepository(app.Todo)
	if err != nil {
		t.Fatalf("create todo Preposiory error, %+v", err)
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
			_, err := tt.r.ListTodos()
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlTodoRepository.ListTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTodoRepository_GetTodo(t *testing.T) {
	sto, err := CreateTodoRepository(app.Todo)
	if err != nil {
		t.Fatalf("create todo Preposiory error, %+v", err)
	}
	type args struct {
		id int
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
			_, err := tt.r.GetTodo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlTodoRepository.GetTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTodoRepository_CompleteTodo(t *testing.T) {
	sto, err := CreateTodoRepository(app.Todo)
	if err != nil {
		t.Fatalf("create todo Preposiory error, %+v", err)
	}
	type args struct {
		id int
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
			if err := tt.r.CompleteTodo(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("MysqlTodoRepository.CompleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoRepository_UpdateTodo(t *testing.T) {
	sto, err := CreateTodoRepository(app.Todo)
	if err != nil {
		t.Fatalf("create todo Preposiory error, %+v", err)
	}
	type args struct {
		todo model.Todo
	}
	tests := []struct {
		name    string
		r       TodoRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{todo: model.Todo{ID: 1, Content: "test"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.UpdateTodo(tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("MysqlTodoRepository.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoRepository_DeleteTodo(t *testing.T) {
	sto, err := CreateTodoRepository(app.Todo)
	if err != nil {
		t.Fatalf("create todo Preposiory error, %+v", err)
	}
	type args struct {
		id int
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
			if err := tt.r.DeleteTodo(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("MysqlTodoRepository.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}