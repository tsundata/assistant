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
	GetTodo(id int) (model.Todo, error)
	CompleteTodo(id int) error
	UpdateTodo(todo model.Todo) error
	DeleteTodo(id int) error
}

type MysqlTodoRepository struct {
	logger *logger.Logger
	db     *sqlx.DB
}

func NewMysqlTodoRepository(logger *logger.Logger, db *sqlx.DB) TodoRepository {
	return &MysqlTodoRepository{logger: logger, db: db}
}

func (r *MysqlTodoRepository) CreateTodo(todo model.Todo) (int64, error) {
	todo.Time = time.Now()
	res, err := r.db.NamedExec("INSERT INTO `todos` (`content`, `priority`, `is_remind_at_time`, `remind_at`, `repeat_method`, `repeat_rule`, `category`, `remark`, `complete`, `time`) VALUES (:content, :priority, :is_remind_at_time, :remind_at, :repeat_method, :repeat_rule, :category, :remark, :complete, :time)", todo)
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
	err := r.db.Select(&items, "SELECT * FROM `todos` ORDER BY `id` DESC")
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return items, nil
}

func (r *MysqlTodoRepository) GetTodo(id int) (model.Todo, error) {
	var item model.Todo
	err := r.db.Get(&item, "SELECT id FROM `todos` WHERE id = ?", id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.Todo{}, err
	}
	return item, nil
}

func (r *MysqlTodoRepository) CompleteTodo(id int) error {
	_, err := r.db.Exec("UPDATE `todos` SET `complete` = 1 WHERE id = ?", id)
	return err
}

func (r *MysqlTodoRepository) UpdateTodo(todo model.Todo) error {
	_, err := r.db.Exec("UPDATE `todos` SET `content` = ? WHERE id = ?", todo.Content, todo.ID)
	return err
}

func (r *MysqlTodoRepository) DeleteTodo(id int) error {
	_, err := r.db.Exec("DELETE FROM `todos` WHERE `id` = ?", id)
	return err
}
