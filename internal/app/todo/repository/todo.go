package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/gorqlite"
)

type TodoRepository interface {
	CreateTodo(todo pb.Todo) (int64, error)
	ListTodos() ([]pb.Todo, error)
	ListRemindTodos() ([]pb.Todo, error)
	GetTodo(id int64) (pb.Todo, error)
	CompleteTodo(id int64) error
	UpdateTodo(todo pb.Todo) error
	DeleteTodo(id int64) error
}

type MysqlTodoRepository struct {
	logger log.Logger
	db     *sqlx.DB
}

func NewMysqlTodoRepository(logger log.Logger, db *sqlx.DB) TodoRepository {
	return &MysqlTodoRepository{logger: logger, db: db}
}

func (r *MysqlTodoRepository) CreateTodo(todo pb.Todo) (int64, error) {
	todo.CreatedAt = util.Now()
	todo.UpdatedAt = util.Now()
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

func (r *MysqlTodoRepository) ListTodos() ([]pb.Todo, error) {
	var items []pb.Todo
	err := r.db.Select(&items, "SELECT * FROM `todos` WHERE `complete` <> 1 ORDER BY `priority` DESC, `created_at` DESC")
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return items, nil
}

func (r *MysqlTodoRepository) ListRemindTodos() ([]pb.Todo, error) {
	var items []pb.Todo
	err := r.db.Select(&items, "SELECT * FROM `todos` WHERE `complete` <> 1 AND `is_remind_at_time` = 1 ORDER BY `priority` DESC")
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return items, nil
}

func (r *MysqlTodoRepository) GetTodo(id int64) (pb.Todo, error) {
	var item pb.Todo
	err := r.db.Get(&item, "SELECT id FROM `todos` WHERE id = ?", id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return pb.Todo{}, err
	}
	return item, nil
}

func (r *MysqlTodoRepository) CompleteTodo(id int64) error {
	_, err := r.db.Exec("UPDATE `todos` SET `complete` = 1 WHERE id = ?", id)
	return err
}

func (r *MysqlTodoRepository) UpdateTodo(todo pb.Todo) error {
	_, err := r.db.Exec("UPDATE `todos` SET `content` = ? WHERE id = ?", todo.Content, todo.Id)
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

func (r *RqliteTodoRepository) CreateTodo(todo pb.Todo) (int64, error) {
	now := util.Now()
	res, err := r.db.WriteOne(fmt.Sprintf("INSERT INTO `todos` (`content`, `priority`, `is_remind_at_time`, `remind_at`, `repeat_method`, `repeat_rule`, `repeat_end_at`, `category`, `remark`, `complete`, `created_at`, `updated_at`) VALUES ('%s', %d, %d, '%s', '%s', '%s', '%s', '%s', '%s', %d, '%s', '%s')",
		todo.Content, todo.Priority, util.BoolInt(todo.IsRemindAtTime), todo.RemindAt, todo.RepeatMethod, todo.RepeatRule, todo.RepeatEndAt, todo.Category, todo.Remark, util.BoolInt(todo.Complete), now, now))
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
		util.Inject(item, m)
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
		util.Inject(item, m)
		items = append(items, item)
	}

	return items, nil
}

func (r *RqliteTodoRepository) GetTodo(id int64) (pb.Todo, error) {
	rows, err := r.db.QueryOne(fmt.Sprintf("SELECT * FROM `todos` WHERE id = %d", id))
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
	_, err := r.db.WriteOne(fmt.Sprintf("UPDATE `todos` SET `complete` = 1 WHERE id = %d", id))
	return err
}

func (r *RqliteTodoRepository) UpdateTodo(todo pb.Todo) error {
	_, err := r.db.WriteOne(fmt.Sprintf("UPDATE `todos` SET `content` = '%s' WHERE id = %d", todo.Content, todo.Id))
	return err
}

func (r *RqliteTodoRepository) DeleteTodo(id int64) error {
	_, err := r.db.WriteOne(fmt.Sprintf("DELETE FROM `todos` WHERE `id` = %d", id))
	return err
}
