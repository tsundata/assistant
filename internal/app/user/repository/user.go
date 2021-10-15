package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetRole(ctx context.Context, userID int) (*pb.Role, error)
	ChangeRoleExp(ctx context.Context, userID int64, exp int64) error
	ChangeRoleAttr(ctx context.Context, userID int64, attr string, val int64) error
	List(ctx context.Context) ([]*pb.User, error)
	Create(ctx context.Context, user *pb.User) (int64, error)
	GetByID(ctx context.Context, id int64) (*pb.User, error)
	GetByName(ctx context.Context, username string) (*pb.User, error)
	Update(ctx context.Context, user *pb.User) error
	ListDevice(ctx context.Context, userID int64) ([]*pb.Device, error)
	CreateDevice(ctx context.Context, device *pb.Device) (int64, error)
	GetDevice(ctx context.Context, id int64) (*pb.Device, error)
}

type MysqlUserRepository struct {
	logger log.Logger
	id     *global.ID
	db     *mysql.Conn
}

func NewMysqlUserRepository(logger log.Logger, id *global.ID, db *mysql.Conn) UserRepository {
	return &MysqlUserRepository{logger: logger, id: id, db: db}
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
		err = r.db.WithContext(ctx).Exec(fmt.Sprintf("INSERT INTO `role_records` (`id`, `profession`, `user_id`, `exp`, `%s`) VALUES (?, '', ?, ?, ?)", attr), r.id.Generate(ctx), userId, exp, val).Error
	} else {
		err = r.db.WithContext(ctx).Exec("INSERT INTO `role_records` (`id`, `profession`, `user_id`, `exp`) VALUES (?, '', ?, ?)", r.id.Generate(ctx), userId, exp).Error
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

func (r *MysqlUserRepository) GetByName(ctx context.Context, username string) (*pb.User, error) {
	var find pb.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&find).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}

func (r *MysqlUserRepository) Update(ctx context.Context, user *pb.User) error {
	return r.db.WithContext(ctx).Model(&pb.User{}).Where("id = ?", user.Id).
		Update("nickname", user.Nickname).
		Update("mobile", user.Mobile).
		Update("remark", user.Remark).
		Error
}

func (r *MysqlUserRepository) ListDevice(ctx context.Context, userID int64) ([]*pb.Device, error) {
	var devices []*pb.Device
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}

func (r *MysqlUserRepository) CreateDevice(ctx context.Context, device *pb.Device) (int64, error) {
	device.Id = r.id.Generate(ctx)
	err := r.db.WithContext(ctx).Create(&device).Error
	if err != nil {
		return 0, err
	}
	return device.Id, nil
}

func (r *MysqlUserRepository) GetDevice(ctx context.Context, id int64) (*pb.Device, error) {
	var find pb.Device
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&find).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}
