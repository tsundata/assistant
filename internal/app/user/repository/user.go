package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/util"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetRole(ctx context.Context, userID int) (*pb.Role, error)
	ChangeRoleExp(ctx context.Context, userID int64, exp int64) error
	ChangeRoleAttr(ctx context.Context, userID int64, attr string, val int64) error
	List(ctx context.Context) ([]*pb.User, error)
	Create(ctx context.Context, user *pb.User) (int64, error)
	GetByID(ctx context.Context, id int64) (*pb.User, error)
	GetByName(ctx context.Context, name string) (*pb.User, error)
	Update(ctx context.Context, user *pb.User) error
}

type MysqlUserRepository struct {
	logger log.Logger
	db     *mysql.Conn
}

func NewMysqlUserRepository(logger log.Logger, db *mysql.Conn) UserRepository {
	return &MysqlUserRepository{logger: logger, db: db}
}

func (r *MysqlUserRepository) GetRole(ctx context.Context, userID int) (*pb.Role, error) {
	var find pb.Role
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&find).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}

func (r *MysqlUserRepository) ChangeRoleExp(ctx context.Context, userID int64, exp int64) error {
	var item pb.Role
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&item).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	item.Exp = item.Exp + exp
	err = r.db.WithContext(ctx).Save(&item).Error
	if err != nil {
		return err
	}
	r.roleRecord(ctx, userID, exp, "", 0)
	return nil
}

func (r *MysqlUserRepository) ChangeRoleAttr(ctx context.Context, userID int64, attr string, val int64) error {
	var item pb.Role
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&item).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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

	err = r.db.WithContext(ctx).Exec(fmt.Sprintf("UPDATE `roles` SET `%s` = ? WHERE `user_id` = ?", attr), oldVal+val, userID).Error
	if err != nil {
		return err
	}
	r.roleRecord(ctx, userID, 0, attr, val)
	return nil
}

func (r *MysqlUserRepository) roleRecord(ctx context.Context, userId int64, exp int64, attr string, val int64) {
	var err error
	if attr != "" {
		err = r.db.WithContext(ctx).Exec(fmt.Sprintf("INSERT INTO `role_records` (`profession`, `user_id`, `exp`, `%s`) VALUES ('', ?, ?, ?)", attr), userId, exp, val).Error
	} else {
		err = r.db.WithContext(ctx).Exec("INSERT INTO `role_records` (`profession`, `user_id`, `exp`) VALUES ('', ?, ?)", userId, exp).Error
	}
	if err != nil {
		r.logger.Error(err)
	}
}

func (r *MysqlUserRepository) List(ctx context.Context) ([]*pb.User, error) {
	var users []*pb.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *MysqlUserRepository) Create(ctx context.Context, user *pb.User) (int64, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (r *MysqlUserRepository) GetByID(ctx context.Context, id int64) (*pb.User, error) {
	var find pb.User
	err := r.db.WithContext(ctx).First(&find, id).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}

func (r *MysqlUserRepository) GetByName(ctx context.Context, name string) (*pb.User, error) {
	var find pb.User
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&find).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}

func (r *MysqlUserRepository) Update(ctx context.Context, user *pb.User) error {
	return r.db.WithContext(ctx).Model(&pb.User{}).Where("id = ?", user.Id).
		Update("name", user.Name).
		Update("mobile", user.Mobile).
		Update("remark", user.Remark).
		Error
}

type RqliteUserRepository struct {
	logger log.Logger
	db     *rqlite.Conn
}

func NewRqliteUserRepository(logger log.Logger, db *rqlite.Conn) *RqliteUserRepository {
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
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
	//user.CreatedAt = util.Now()
	//user.UpdatedAt = util.Now()
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
