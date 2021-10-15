package repository

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *pb.Todo) (int64, error)
	ListTodos(ctx context.Context) ([]*pb.Todo, error)
	ListRemindTodos(ctx context.Context) ([]*pb.Todo, error)
	GetTodo(ctx context.Context, id int64) (*pb.Todo, error)
	CompleteTodo(ctx context.Context, id int64) error
	UpdateTodo(ctx context.Context, todo *pb.Todo) error
	DeleteTodo(ctx context.Context, id int64) error
}

type MysqlTodoRepository struct {
	id *global.ID
	db *mysql.Conn
}

func NewMysqlTodoRepository(id *global.ID, db *mysql.Conn) TodoRepository {
	return &MysqlTodoRepository{id: id, db: db}
}

func (r *MysqlTodoRepository) CreateTodo(ctx context.Context, todo *pb.Todo) (int64, error) {
	todo.Id = r.id.Generate(ctx)
	err := r.db.WithContext(ctx).Create(&todo)
	if err != nil {
		return 0, nil
	}
	return todo.Id, nil
}

func (r *MysqlTodoRepository) ListTodos(ctx context.Context) ([]*pb.Todo, error) {
	var items []*pb.Todo
	err := r.db.WithContext(ctx).Where("complete <> ?", 1).
		Order("priority DESC").
		Order("created_at DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlTodoRepository) ListRemindTodos(ctx context.Context) ([]*pb.Todo, error) {
	var items []*pb.Todo
	err := r.db.WithContext(ctx).Where("complete <> ?", 1).
		Where("is_remind_at_time = ?", 1).
		Order("priority DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlTodoRepository) GetTodo(ctx context.Context, id int64) (*pb.Todo, error) {
	var find pb.Todo
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&find).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}

func (r *MysqlTodoRepository) CompleteTodo(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Model(&pb.Todo{}).Where("id = ?", id).Update("complete", true).Error
}

func (r *MysqlTodoRepository) UpdateTodo(ctx context.Context, todo *pb.Todo) error {
	return r.db.WithContext(ctx).Save(&todo).Error
}

func (r *MysqlTodoRepository) DeleteTodo(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.Todo{}).Error
}
