package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/model"
)

type UserRepository interface {
	GetRole(userID int) (model.Role, error)
	ChangeRoleExp(userID int, exp int) error
	ChangeRoleAttr(userID int, attr string, val int) error
	List() ([]model.User, error)
	Create(user model.User) (int64, error)
	GetByID(id int64) (model.User, error)
	GetByName(name string) (model.User, error)
	Update(user model.User) error
}

type MysqlUserRepository struct {
	logger *logger.Logger
	db     *sqlx.DB
}

func NewMysqlUserRepository(logger *logger.Logger, db *sqlx.DB) UserRepository {
	return &MysqlUserRepository{logger: logger, db: db}
}

func (r *MysqlUserRepository) GetRole(userId int) (model.Role, error) {
	var item model.Role
	err := r.db.Get(&item, "SELECT * FROM `roles` WHERE user_id = ?", userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.Role{}, err
	}
	return item, nil
}

func (r *MysqlUserRepository) ChangeRoleExp(userID int, exp int) error {
	var item model.Role
	err := r.db.Get(&item, "SELECT `id`, `exp` FROM `roles` WHERE `user_id` = ?", userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	_, err = r.db.Exec("UPDATE `roles` SET `exp` = ? WHERE `user_id` = ?", item.Exp+exp, userID)
	if err != nil {
		return err
	}
	r.roleRecord(userID, exp, "", 0)
	return nil
}

func (r *MysqlUserRepository) ChangeRoleAttr(userID int, attr string, val int) error {
	var item model.Role
	err := r.db.Get(&item, fmt.Sprintf("SELECT `id`, `%s` FROM `roles` WHERE `user_id` = ?", attr), userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	oldVal := 0
	switch model.RoleAttr(attr) {
	case model.StrengthAttr:
		oldVal = item.Strength
	case model.CultureAttr:
		oldVal = item.Culture
	case model.EnvironmentAttr:
		oldVal = item.Environment
	case model.CharismaAttr:
		oldVal = item.Charisma
	case model.TalentAttr:
		oldVal = item.Talent
	case model.IntellectAttr:
		oldVal = item.Intellect
	}
	_, err = r.db.Exec(fmt.Sprintf("UPDATE `roles` SET `%s` = ? WHERE `user_id` = ?", attr), oldVal+val, userID)
	if err != nil {
		return err
	}
	r.roleRecord(userID, 0, attr, val)
	return nil
}

func (r *MysqlUserRepository) roleRecord(userId int, exp int, attr string, val int) {
	var err error
	if attr != "" {
		_, err = r.db.Exec(fmt.Sprintf("INSERT INTO `role_records` (`profession`, `user_id`, `exp`, `%s`) VALUES ('', ?, ?, ?, ?)", attr), userId, exp, val)
	} else {
		_, err = r.db.Exec("INSERT INTO `role_records` (`profession`, `user_id`, `exp`) VALUES ('', ?, ?, ?)", userId, exp)
	}
	if err != nil {
		r.logger.Error(err)
	}
}

func (r *MysqlUserRepository) List() ([]model.User, error) {
	var users []model.User
	err := r.db.Select(&users, "SELECT `id`, `name`, `mobile`, `remark` FROM `users`")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *MysqlUserRepository) Create(user model.User) (int64, error) {
	res, err := r.db.NamedExec("INSERT INTO `users` (`name`, `mobile`, `remark`) VALUES (:name, :mobile, :remark)", user)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MysqlUserRepository) GetByID(id int64) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT id, `name`, `mobile`, `remark` FROM `users` WHERE `id` = ? LIMIT 1", id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.User{}, err
	}

	return user, nil
}

func (r *MysqlUserRepository) GetByName(name string) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT id, `name`, `mobile`, `remark` FROM `users` WHERE `name` = ? LIMIT 1", name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.User{}, err
	}

	return user, nil
}

func (r *MysqlUserRepository) Update(user model.User) error {
	_, err := r.db.Exec("UPDATE users SET `name` = ?, `mobile` = ?, `remark` = ? WHERE id = ?", user.Name, user.Mobile, user.Remark, user.ID)
	return err
}
