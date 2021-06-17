package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/model"
	"time"
)

type UserRepository interface {
	GetRole(userID int) (model.Role, error)
	ChangeRoleExp(userID int, exp int) error
	ChangeRoleAttr(userID int, attr string, val int) error
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
	switch attr {
	case model.RoleStrength:
		oldVal = item.Strength
	case model.RoleCulture:
		oldVal = item.Culture
	case model.RoleEnvironment:
		oldVal = item.Environment
	case model.RoleCharisma:
		oldVal = item.Charisma
	case model.RoleTalent:
		oldVal = item.Talent
	case model.RoleIntellect:
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
		_, err = r.db.Exec(fmt.Sprintf("INSERT INTO `role_records` (`profession`, `user_id`, `exp`, `%s`, `time`) VALUES ('', ?, ?, ?, ?)", attr), userId, exp, val, time.Now())
	} else {
		_, err = r.db.Exec("INSERT INTO `role_records` (`profession`, `user_id`, `exp`, `time`) VALUES ('', ?, ?, ?)", userId, exp, time.Now())
	}
	if err != nil {
		r.logger.Error(err)
	}
}
