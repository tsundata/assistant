package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rqlite/gorqlite"
	"github.com/tsundata/assistant/api/model"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/util"
	"time"
)

type TodoRepository interface {
	CreateTodo(todo model.Todo) (int64, error)
	ListTodos() ([]model.Todo, error)
	ListRemindTodos() ([]model.Todo, error)
	GetTodo(id int64) (model.Todo, error)
	CompleteTodo(id int64) error
	UpdateTodo(todo model.Todo) error
	DeleteTodo(id int64) error
}

type MysqlTodoRepository struct {
	logger log.Logger
	db     *sqlx.DB
}

func NewMysqlTodoRepository(logger log.Logger, db *sqlx.DB) TodoRepository {
	return &MysqlTodoRepository{logger: logger, db: db}
}

func (r *MysqlTodoRepository) CreateTodo(todo model.Todo) (int64, error) {
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	res, err := r.db.NamedExec("INSERT INTO `todos` (`content`, `priority`, `is_remind_at_time`, `remind_at`, `repeat_method`, `repeat_rule`, `repeat_end_at`, `category`, `remark`, `complete`, `created_at`, `updated_at`) VALUES (:content, :priority, :is_remind_at_time, :remind_at, :repeat_method, :repeat_rule, :repeat_end_at, :category, :remark, :complete, :created_at, :updated_at)", todo)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MysqlTodoRepository) ListTodos() ([]model.Todo, error) {
	var items []model.Todo
	err := r.db.Select(&items, "SELECT * FROM `todos` WHERE `complete` <> 1 ORDER BY `priority` DESC, `created_at` DESC")
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return items, nil
}

func (r *MysqlTodoRepository) ListRemindTodos() ([]model.Todo, error) {
	var items []model.Todo
	err := r.db.Select(&items, "SELECT * FROM `todos` WHERE `complete` <> 1 AND `is_remind_at_time` = 1 ORDER BY `priority` DESC")
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return items, nil
}

func (r *MysqlTodoRepository) GetTodo(id int64) (model.Todo, error) {
	var item model.Todo
	err := r.db.Get(&item, "SELECT id FROM `todos` WHERE id = ?", id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.Todo{}, err
	}
	return item, nil
}

func (r *MysqlTodoRepository) CompleteTodo(id int64) error {
	_, err := r.db.Exec("UPDATE `todos` SET `complete` = 1 WHERE id = ?", id)
	return err
}

func (r *MysqlTodoRepository) UpdateTodo(todo model.Todo) error {
	_, err := r.db.Exec("UPDATE `todos` SET `content` = ? WHERE id = ?", todo.Content, todo.ID)
	return err
}

func (r *MysqlTodoRepository) DeleteTodo(id int64) error {
	_, err := r.db.Exec("DELETE FROM `todos` WHERE `id` = ?", id)
	return err
}

func NewRqliteTodoRepository(logger log.Logger, db gorqlite.Connection) TodoRepository {
	return &RqliteTodoRepository{logger: logger, db: db}
}

type RqliteTodoRepository struct {
	logger log.Logger
	db     gorqlite.Connection
}

func (r *RqliteTodoRepository) CreateTodo(todo model.Todo) (int64, error) {
	now := util.Now()
	res, err := r.db.WriteOne(fmt.Sprintf("INSERT INTO `todos` (`content`, `priority`, `is_remind_at_time`, `remind_at`, `repeat_method`, `repeat_rule`, `repeat_end_at`, `category`, `remark`, `complete`, `created_at`, `updated_at`) VALUES ('%s', %d, %d, '%s', '%s', '%s', '%s', '%s', '%s', %d, '%s', '%s')",
		todo.Content, todo.Priority, util.BoolInt(todo.IsRemindAtTime), todo.RemindAt, todo.RepeatMethod, todo.RepeatRule, todo.RepeatEndAt, todo.Category, todo.Remark, util.BoolInt(todo.Complete), now, now))
	if err != nil {
		return 0, err
	}
	return res.LastInsertID, nil
}

func (r *RqliteTodoRepository) ListTodos() ([]model.Todo, error) {
	panic("implement me")
}

func (r *RqliteTodoRepository) ListRemindTodos() ([]model.Todo, error) {
	panic("implement me")
}

func (r *RqliteTodoRepository) GetTodo(id int64) (model.Todo, error) {
	panic("implement me")
}

func (r *RqliteTodoRepository) CompleteTodo(id int64) error {
	_, err := r.db.WriteOne(fmt.Sprintf("UPDATE `todos` SET `complete` = 1 WHERE id = %d", id))
	return err
}

func (r *RqliteTodoRepository) UpdateTodo(todo model.Todo) error {
	_, err := r.db.WriteOne(fmt.Sprintf("UPDATE `todos` SET `content` = '%s' WHERE id = %d", todo.Content, todo.ID))
	return err
}

func (r *RqliteTodoRepository) DeleteTodo(id int64) error {
	_, err := r.db.WriteOne(fmt.Sprintf("DELETE FROM `todos` WHERE `id` = %d", id))
	return err
}
