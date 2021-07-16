package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/api/model"
	"github.com/tsundata/assistant/internal/pkg/log"
)

type MessageRepository interface {
	GetByID(id int64) (model.Message, error)
	GetByUUID(uuid string) (model.Message, error)
	ListByType(t string) ([]model.Message, error)
	List() ([]model.Message, error)
	Create(message model.Message) (int64, error)
	Delete(id int64) error
}

type MysqlMessageRepository struct {
	logger log.Logger
	db     *sqlx.DB
}

func NewMysqlMessageRepository(logger log.Logger, db *sqlx.DB) MessageRepository {
	return &MysqlMessageRepository{logger: logger, db: db}
}

func (r *MysqlMessageRepository) GetByID(id int64) (model.Message, error) {
	var message model.Message
	err := r.db.Get(&message, "SELECT id, uuid, text, `type`, `created_at` FROM `messages` WHERE `id` = ? LIMIT 1", id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.Message{}, err
	}

	return message, nil
}

func (r *MysqlMessageRepository) GetByUUID(uuid string) (model.Message, error) {
	var message model.Message
	err := r.db.Get(&message, "SELECT id, uuid, text, `type`, `created_at` FROM `messages` WHERE `uuid` = ? LIMIT 1", uuid)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.Message{}, err
	}

	return message, nil
}

func (r *MysqlMessageRepository) ListByType(t string) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Select(&messages, "SELECT id, uuid, text, `type`, `created_at` FROM `messages` WHERE `type` = ? ORDER BY `id` DESC", t)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return []model.Message{}, err
	}

	return messages, nil
}

func (r *MysqlMessageRepository) List() ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Select(&messages, "SELECT uuid, text, `type`, `created_at` FROM `messages` WHERE `type` <> ? ORDER BY `id` DESC",
		model.MessageTypeAction)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MysqlMessageRepository) Create(message model.Message) (int64, error) {
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
