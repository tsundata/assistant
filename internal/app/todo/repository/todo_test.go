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
				t.Errorf("TodoRepository.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
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
				t.Errorf("TodoRepository.ListTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTodoRepository_ListRemindTodos(t *testing.T) {
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
			_, err := tt.r.ListRemindTodos()
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.ListTodos() error = %v, wantErr %v", err, tt.wantErr)
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
			_, err := tt.r.GetTodo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.GetTodo() error = %v, wantErr %v", err, tt.wantErr)
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
			if err := tt.r.CompleteTodo(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.CompleteTodo() error = %v, wantErr %v", err, tt.wantErr)
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
				t.Errorf("TodoRepository.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
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
			if err := tt.r.DeleteTodo(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
