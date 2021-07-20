package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/util"
)

type MessageRepository interface {
	GetByID(id int64) (pb.Message, error)
	GetByUUID(uuid string) (pb.Message, error)
	ListByType(t string) ([]pb.Message, error)
	List() ([]pb.Message, error)
	Create(message pb.Message) (int64, error)
	Delete(id int64) error
}

type MysqlMessageRepository struct {
	logger log.Logger
	db     *sqlx.DB
}

func NewMysqlMessageRepository(logger log.Logger, db *sqlx.DB) MessageRepository {
	return &MysqlMessageRepository{logger: logger, db: db}
}

func (r *MysqlMessageRepository) GetByID(id int64) (pb.Message, error) {
	var message pb.Message
	err := r.db.Get(&message, "SELECT id, uuid, text, `type`, `created_at` FROM `messages` WHERE `id` = ? LIMIT 1", id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return pb.Message{}, err
	}

	return message, nil
}

func (r *MysqlMessageRepository) GetByUUID(uuid string) (pb.Message, error) {
	var message pb.Message
	err := r.db.Get(&message, "SELECT id, uuid, text, `type`, `created_at` FROM `messages` WHERE `uuid` = ? LIMIT 1", uuid)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return pb.Message{}, err
	}

	return message, nil
}

func (r *MysqlMessageRepository) ListByType(t string) ([]pb.Message, error) {
	var messages []pb.Message
	err := r.db.Select(&messages, "SELECT id, uuid, text, `type`, `created_at` FROM `messages` WHERE `type` = ? ORDER BY `id` DESC", t)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return []pb.Message{}, err
	}

	return messages, nil
}

func (r *MysqlMessageRepository) List() ([]pb.Message, error) {
	var messages []pb.Message
	err := r.db.Select(&messages, "SELECT uuid, text, `type`, `created_at` FROM `messages` WHERE `type` <> ? ORDER BY `id` DESC",
		enum.MessageTypeAction)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MysqlMessageRepository) Create(message pb.Message) (int64, error) {
	res, err := r.db.NamedExec("INSERT INTO `messages` (`uuid`, `type`, `text`) VALUES (:uuid, :type, :text)", message)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MysqlMessageRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM `messages` WHERE `id` = ?", id)
	return err
}

type RqliteMessageRepository struct {
	logger log.Logger
	db     *rqlite.Conn
}

func NewRqliteMessageRepository(logger log.Logger, db *rqlite.Conn) MessageRepository {
	return &RqliteMessageRepository{logger: logger, db: db}
}

func (r *RqliteMessageRepository) GetByID(id int64) (pb.Message, error) {
	rows, err := r.db.QueryOne("SELECT id, uuid, text, `type`, `created_at` FROM `messages` WHERE `id` = '%d' LIMIT 1", id)
	if err != nil {
		return pb.Message{}, nil
	}

	var message pb.Message
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.Message{}, err
		}
		util.Inject(&message, m)
	}

	return message, nil
}

func (r *RqliteMessageRepository) GetByUUID(uuid string) (pb.Message, error) {
	rows, err := r.db.QueryOne("SELECT id, uuid, text, `type`, `created_at` FROM `messages` WHERE `uuid` = '%s' LIMIT 1", uuid)
	if err != nil {
		return pb.Message{}, err
	}

	var message pb.Message
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.Message{}, err
		}
		util.Inject(&message, m)
	}

	return message, nil
}

func (r *RqliteMessageRepository) ListByType(t string) ([]pb.Message, error) {
	rows, err := r.db.QueryOne("SELECT id, uuid, text, `type`, `created_at` FROM `messages` WHERE `type` = '%s' ORDER BY `id` DESC", t)
	if err != nil {
		return nil, nil
	}

	var messages []pb.Message
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return nil, err
		}
		var item pb.Message
		util.Inject(item, m)
		messages = append(messages, item)
	}

	return messages, nil
}

func (r *RqliteMessageRepository) List() ([]pb.Message, error) {
	rows, err := r.db.QueryOne("SELECT uuid, text, `type`, `created_at` FROM `messages` WHERE `type` <> '%s' ORDER BY `id` DESC", enum.MessageTypeAction)
	if err != nil {
		return nil, err
	}

	var messages []pb.Message
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return nil, err
		}
		var item pb.Message
		util.Inject(item, m)
		messages = append(messages, item)
	}

	return messages, nil
}

func (r *RqliteMessageRepository) Create(message pb.Message) (int64, error) {
	res, err := r.db.WriteOne("INSERT INTO `messages` (`uuid`, `type`, `text`) VALUES ('%s', '%s', '%s')", message.Uuid, message.Type, message.Text)
	if err != nil {
		return 0, err
	}

	return res.LastInsertID, nil
}

func (r *RqliteMessageRepository) Delete(id int64) error {
	_, err := r.db.WriteOne("DELETE FROM `messages` WHERE `id` = '%d'", id)
	return err
}
