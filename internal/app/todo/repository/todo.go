package repository

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/util"
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
	db *mysql.Conn
}

func NewMysqlTodoRepository(db *mysql.Conn) TodoRepository {
	return &MysqlTodoRepository{db: db}
}

func (r *MysqlTodoRepository) CreateTodo(ctx context.Context, todo *pb.Todo) (int64, error) {
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

func NewRqliteTodoRepository(db *rqlite.Conn) *RqliteTodoRepository {
	return &RqliteTodoRepository{db: db}
}

type RqliteTodoRepository struct {
	db *rqlite.Conn
}

func (r *RqliteTodoRepository) CreateTodo(todo pb.Todo) (int64, error) {
	now := util.Now()
	res, err := r.db.WriteOne("INSERT INTO `todos` (`content`, `priority`, `is_remind_at_time`, `remind_at`, `repeat_method`, `repeat_rule`, `repeat_end_at`, `category`, `remark`, `complete`, `created_at`, `updated_at`) VALUES ('%s', %d, %d, '%s', '%s', '%s', '%s', '%s', '%s', %d, '%s', '%s')",
		todo.Content, todo.Priority, util.BoolInt(todo.IsRemindAtTime), todo.RemindAt, todo.RepeatMethod, todo.RepeatRule, todo.RepeatEndAt, todo.Category, todo.Remark, util.BoolInt(todo.Complete), now, now)
	if err != nil {
		return 0, err
	}
	return res.LastInsertID, nil
}

func (r *RqliteTodoRepository) ListTodos() ([]pb.Todo, error) {
	rows, err := r.db.QueryOne("SELECT * FROM `todos` WHERE `complete` <> 1 ORDER BY `priority` DESC, `created_at` DESC")
	if err != nil {
		return nil, err
	}

	var items []pb.Todo
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return nil, err
		}
		var item pb.Todo
		util.Inject(&item, m)
		items = append(items, item)
	}

	return items, nil
}

func (r *RqliteTodoRepository) ListRemindTodos() ([]pb.Todo, error) {
	rows, err := r.db.QueryOne("SELECT * FROM `todos` WHERE `complete` <> 1 AND `is_remind_at_time` = 1 ORDER BY `priority` DESC")
	if err != nil {
		return nil, err
	}

	var items []pb.Todo
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return nil, err
		}
		var item pb.Todo
		util.Inject(&item, m)
		items = append(items, item)
	}

	return items, nil
}

func (r *RqliteTodoRepository) GetTodo(id int64) (pb.Todo, error) {
	rows, err := r.db.QueryOne("SELECT * FROM `todos` WHERE id = %d", id)
	if err != nil {
		return pb.Todo{}, err
	}

	var find pb.Todo
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.Todo{}, err
		}
		util.Inject(&find, m)
	}

	return find, nil
}

func (r *RqliteTodoRepository) CompleteTodo(id int64) error {
	_, err := r.db.WriteOne("UPDATE `todos` SET `complete` = 1 WHERE id = %d", id)
	return err
}

func (r *RqliteTodoRepository) UpdateTodo(todo pb.Todo) error {
	_, err := r.db.WriteOne("UPDATE `todos` SET `content` = '%s' WHERE id = %d", todo.Content, todo.Id)
	return err
}

func (r *RqliteTodoRepository) DeleteTodo(id int64) error {
	_, err := r.db.WriteOne("DELETE FROM `todos` WHERE `id` = %d", id)
	return err
}
