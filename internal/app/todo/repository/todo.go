package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/model"
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
	logger *logger.Logger
	db     *sqlx.DB
}

func NewMysqlTodoRepository(logger *logger.Logger, db *sqlx.DB) TodoRepository {
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
