package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/util"
)

type UserRepository interface {
	GetRole(userID int64) (pb.Role, error)
	ChangeRoleExp(userID int64, exp int64) error
	ChangeRoleAttr(userID int64, attr string, val int64) error
	List() ([]pb.User, error)
	Create(user pb.User) (int64, error)
	GetByID(id int64) (pb.User, error)
	GetByName(name string) (pb.User, error)
	Update(user pb.User) error
}

type MysqlUserRepository struct {
	logger log.Logger
	db     *sqlx.DB
}

func NewMysqlUserRepository(logger log.Logger, db *sqlx.DB) UserRepository {
	return &MysqlUserRepository{logger: logger, db: db}
}

func (r *MysqlUserRepository) GetRole(userId int64) (pb.Role, error) {
	var item pb.Role
	err := r.db.Get(&item, "SELECT * FROM `roles` WHERE user_id = ?", userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return pb.Role{}, err
	}
	return item, nil
}

func (r *MysqlUserRepository) ChangeRoleExp(userID int64, exp int64) error {
	var item pb.Role
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

func (r *MysqlUserRepository) ChangeRoleAttr(userID int64, attr string, val int64) error {
	var item pb.Role
	err := r.db.Get(&item, fmt.Sprintf("SELECT `id`, `%s` FROM `roles` WHERE `user_id` = ?", attr), userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	oldVal := int64(0)
	switch enum.RoleAttr(attr) {
	case enum.StrengthAttr:
		oldVal = item.Strength
	case enum.CultureAttr:
		oldVal = item.Culture
	case enum.EnvironmentAttr:
		oldVal = item.Environment
	case enum.CharismaAttr:
		oldVal = item.Charisma
	case enum.TalentAttr:
		oldVal = item.Talent
	case enum.IntellectAttr:
		oldVal = item.Intellect
	}
	_, err = r.db.Exec(fmt.Sprintf("UPDATE `roles` SET `%s` = ? WHERE `user_id` = ?", attr), oldVal+val, userID)
	if err != nil {
		return err
	}
	r.roleRecord(userID, 0, attr, val)
	return nil
}

func (r *MysqlUserRepository) roleRecord(userId int64, exp int64, attr string, val int64) {
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

func (r *MysqlUserRepository) List() ([]pb.User, error) {
	var users []pb.User
	err := r.db.Select(&users, "SELECT `id`, `name`, `mobile`, `remark` FROM `users`")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *MysqlUserRepository) Create(user pb.User) (int64, error) {
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

func (r *MysqlUserRepository) GetByID(id int64) (pb.User, error) {
	var user pb.User
	err := r.db.Get(&user, "SELECT id, `name`, `mobile`, `remark` FROM `users` WHERE `id` = ? LIMIT 1", id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return pb.User{}, err
	}

	return user, nil
}

func (r *MysqlUserRepository) GetByName(name string) (pb.User, error) {
	var user pb.User
	err := r.db.Get(&user, "SELECT id, `name`, `mobile`, `remark` FROM `users` WHERE `name` = ? LIMIT 1", name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return pb.User{}, err
	}

	return user, nil
}

func (r *MysqlUserRepository) Update(user pb.User) error {
	_, err := r.db.Exec("UPDATE users SET `name` = ?, `mobile` = ?, `remark` = ? WHERE id = ?", user.Name, user.Mobile, user.Remark, user.Id)
	return err
}

type RqliteUserRepository struct {
	logger log.Logger
	db     *rqlite.Conn
}

func NewRqliteUserRepository(logger log.Logger, db *rqlite.Conn) UserRepository {
	return &RqliteUserRepository{logger: logger, db: db}
}

func (r *RqliteUserRepository) GetRole(userId int64) (pb.Role, error) {
	rows, err := r.db.QueryOne("SELECT * FROM `roles` WHERE user_id = '%d'", userId)
	if err != nil {
		return pb.Role{}, nil
	}

	var item pb.Role
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.Role{}, err
		}
		util.Inject(&item, m)
	}

	return item, nil
}

func (r *RqliteUserRepository) ChangeRoleExp(userID int64, exp int64) error {
	rows, err := r.db.QueryOne("SELECT `id`, `exp` FROM `roles` WHERE `user_id` = '%d'", userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	var item pb.Role
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return err
		}
		util.Inject(&item, m)
	}

	_, err = r.db.WriteOne("UPDATE `roles` SET `exp` = '%d' WHERE `user_id` = '%d'", item.Exp+exp, userID)
	if err != nil {
		return err
	}
	r.roleRecord(userID, exp, "", 0)
	return nil
}

func (r *RqliteUserRepository) ChangeRoleAttr(userID int64, attr string, val int64) error {
	rows, err := r.db.QueryOne("SELECT `id`, `%s` FROM `roles` WHERE `user_id` = '%d' Limit 1", attr, userID)
	if err != nil {
		return err
	}

	var item pb.Role
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return err
		}
		util.Inject(&item, m)
	}

	oldVal := int64(0)
	switch enum.RoleAttr(attr) {
	case enum.StrengthAttr:
		oldVal = item.Strength
	case enum.CultureAttr:
		oldVal = item.Culture
	case enum.EnvironmentAttr:
		oldVal = item.Environment
	case enum.CharismaAttr:
		oldVal = item.Charisma
	case enum.TalentAttr:
		oldVal = item.Talent
	case enum.IntellectAttr:
		oldVal = item.Intellect
	}
	_, err = r.db.WriteOne("UPDATE `roles` SET `%s` = '%d' WHERE `user_id` = '%d'", attr, oldVal+val, userID)
	if err != nil {
		return err
	}
	r.roleRecord(userID, 0, attr, val)
	return nil
}

func (r *RqliteUserRepository) roleRecord(userId int64, exp int64, attr string, val int64) {
	var err error
	if attr != "" {
		_, err = r.db.WriteOne("INSERT INTO `role_records` (`profession`, `user_id`, `exp`, `%s`) VALUES ('', '%d', '%d', '%d')", attr, userId, exp, val)
	} else {
		_, err = r.db.WriteOne("INSERT INTO `role_records` (`profession`, `user_id`, `exp`) VALUES ('', '%d', '%d')", userId, exp)
	}
	if err != nil {
		r.logger.Error(err)
	}
}

func (r *RqliteUserRepository) List() ([]pb.User, error) {
	rows, err := r.db.QueryOne("SELECT `id`, `name`, `mobile`, `remark` FROM `users`")
	if err != nil {
		return nil, err
	}

	var users []pb.User
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return nil, err
		}
		var item pb.User
		util.Inject(&item, m)
		users = append(users, item)
	}

	return users, nil
}

func (r *RqliteUserRepository) Create(user pb.User) (int64, error) {
	user.CreatedAt = util.Now()
	user.UpdatedAt = util.Now()
	res, err := r.db.WriteOne("INSERT INTO `users` (`name`, `mobile`, `remark`, `created_at`, `updated_at`) VALUES ('%s', '%s', '%s', '%s', '%s')",
		user.Name, user.Mobile, user.Remark, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return 0, err
	}

	return res.LastInsertID, nil
}

func (r *RqliteUserRepository) GetByID(id int64) (pb.User, error) {
	rows, err := r.db.QueryOne("SELECT id, `name`, `mobile`, `remark` FROM `users` WHERE `id` = '%d' LIMIT 1", id)
	if err != nil {
		return pb.User{}, err
	}

	var user pb.User
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.User{}, err
		}
		util.Inject(&user, m)
	}

	return user, nil
}

func (r *RqliteUserRepository) GetByName(name string) (pb.User, error) {
	rows, err := r.db.QueryOne("SELECT id, `name`, `mobile`, `remark` FROM `users` WHERE `name` = '%s' LIMIT 1", name)
	if err != nil {
		return pb.User{}, err
	}

	var user pb.User
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.User{}, err
		}
		util.Inject(&user, m)
	}

	return user, nil
}

func (r *RqliteUserRepository) Update(user pb.User) error {
	_, err := r.db.WriteOne("UPDATE users SET `name` = '%s', `mobile` = '%s', `remark` = '%s' WHERE id = '%d'", user.Name, user.Mobile, user.Remark, user.Id)
	return err
}
