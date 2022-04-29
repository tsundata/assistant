package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"gorm.io/gorm"
	"time"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *pb.Todo) (int64, error)
	ListTodos(ctx context.Context, userId int64) ([]*pb.Todo, error)
	ListRemindTodos(ctx context.Context, userId int64) ([]*pb.Todo, error)
	GetTodo(ctx context.Context, id int64) (*pb.Todo, error)
	GetTodoBySequence(ctx context.Context, userId, sequence int64) (*pb.Todo, error)
	CompleteTodo(ctx context.Context, id int64) error
	CompleteTodoBySequence(ctx context.Context, userId, sequence int64) error
	UpdateTodo(ctx context.Context, todo *pb.Todo) error
	DeleteTodo(ctx context.Context, id int64) error
	DeleteTodoBySequence(ctx context.Context, userId, sequence int64) error
}

type MysqlTodoRepository struct {
	locker *global.Locker
	id     *global.ID
	db     *mysql.Conn
}

func NewMysqlTodoRepository(locker *global.Locker, id *global.ID, db *mysql.Conn) TodoRepository {
	return &MysqlTodoRepository{locker: locker, id: id, db: db}
}

func (r *MysqlTodoRepository) CreateTodo(ctx context.Context, todo *pb.Todo) (int64, error) {
	l, err := r.locker.Acquire(fmt.Sprintf("chatbot:todo:create:%d", todo.UserId))
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = l.Release()
	}()

	// sequence
	sequence := int64(0)
	var max pb.Todo
	err = r.db.WithContext(ctx).Where("user_id = ?", todo.UserId).Order("sequence DESC").Take(&max).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	if max.Sequence > 0 {
		sequence = max.Sequence
	}
	sequence += 1

	todo.Id = r.id.Generate(ctx)
	todo.Sequence = sequence
	todo.CreatedAt = time.Now().Unix()
	todo.UpdatedAt = time.Now().Unix()
	err = r.db.WithContext(ctx).Create(&todo).Error
	if err != nil {
		return 0, nil
	}
	return todo.Id, nil
}

func (r *MysqlTodoRepository) ListTodos(ctx context.Context, userId int64) ([]*pb.Todo, error) {
	var items []*pb.Todo
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Order("priority DESC").
		Order("created_at DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlTodoRepository) ListRemindTodos(ctx context.Context, userId int64) ([]*pb.Todo, error) {
	var items []*pb.Todo
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Where("complete <> ?", 1).
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

func (r *MysqlTodoRepository) GetTodoBySequence(ctx context.Context, userId, sequence int64) (*pb.Todo, error) {
	var find pb.Todo
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND sequence = ?", userId, sequence).
		First(&find).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}

func (r *MysqlTodoRepository) CompleteTodo(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Model(&pb.Todo{}).
		Where("id = ?", id).
		Update("complete", true).Error
}

func (r *MysqlTodoRepository) CompleteTodoBySequence(ctx context.Context, userId, sequence int64) error {
	return r.db.WithContext(ctx).Model(&pb.Todo{}).
		Where("user_id = ? AND sequence = ?", userId, sequence).
		Update("complete", true).Error
}

func (r *MysqlTodoRepository) UpdateTodo(ctx context.Context, todo *pb.Todo) error {
	return r.db.WithContext(ctx).Model(&pb.Todo{}).
		Where("user_id = ? AND sequence = ?", todo.UserId, todo.Sequence).
		UpdateColumns(map[string]interface{}{
			"content":    todo.Content,
			"category":   todo.Category,
			"remark":     todo.Remark,
			"priority":   todo.Priority,
			"updated_at": time.Now().Unix(),
		}).Error
}

func (r *MysqlTodoRepository) DeleteTodo(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.Todo{}).Error
}

func (r *MysqlTodoRepository) DeleteTodoBySequence(ctx context.Context, userId, sequence int64) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND sequence = ?", userId, sequence).
		Delete(&pb.Todo{}).Error
}
