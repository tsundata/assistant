package repository

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"gorm.io/gorm"
	"time"
)

type SystemRepository interface {
	CreateCounter(ctx context.Context, counter *pb.Counter) (int64, error)
	IncreaseCounter(ctx context.Context, id, amount int64) error
	DecreaseCounter(ctx context.Context, id, amount int64) error
	ListCounter(ctx context.Context, userId int64) ([]*pb.Counter, error)
	GetCounter(ctx context.Context, id int64) (pb.Counter, error)
	GetCounterByFlag(ctx context.Context, userId int64, flag string) (pb.Counter, error)
}

type MysqlSystemRepository struct {
	logger log.Logger
	id     *global.ID
	db     *mysql.Conn
}

func NewMysqlSystemRepository(logger log.Logger, id *global.ID, db *mysql.Conn) SystemRepository {
	return &MysqlSystemRepository{logger: logger, id: id, db: db}
}

func (r *MysqlSystemRepository) CreateCounter(ctx context.Context, counter *pb.Counter) (int64, error) {
	counter.Id = r.id.Generate(ctx)
	counter.CreatedAt = time.Now().Unix()
	counter.UpdatedAt = time.Now().Unix()
	err := r.db.WithContext(ctx).Create(&counter)
	if err != nil {
		return 0, nil
	}
	r.record(ctx, counter.Id, counter.Digit)
	return counter.Id, nil
}

func (r *MysqlSystemRepository) IncreaseCounter(ctx context.Context, id, amount int64) error {
	err := r.db.WithContext(ctx).Model(&pb.Counter{}).
		Where("id = ?", id).
		Update("digit", gorm.Expr("digit + ?", amount)).Error
	if err != nil {
		return err
	}
	r.record(ctx, id, amount)
	return nil
}

func (r *MysqlSystemRepository) DecreaseCounter(ctx context.Context, id, amount int64) error {
	err := r.db.WithContext(ctx).Model(&pb.Counter{}).
		Where("id = ?", id).
		Update("digit", gorm.Expr("digit - ?", amount)).Error
	if err != nil {
		return err
	}
	r.record(ctx, id, -amount)
	return nil
}

func (r *MysqlSystemRepository) ListCounter(ctx context.Context, userId int64) ([]*pb.Counter, error) {
	var items []*pb.Counter
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).
		Order("updated_at DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlSystemRepository) record(ctx context.Context, id, digit int64) {
	err := r.db.WithContext(ctx).Exec("INSERT INTO `counter_records` (`id`, `counter_id`, `digit`, `created_at`) VALUES (?, ?, ?, ?)",
		r.id.Generate(ctx), id, digit, time.Now().Unix()).Error
	if err != nil {
		r.logger.Error(err)
	}
}

func (r *MysqlSystemRepository) GetCounter(ctx context.Context, id int64) (pb.Counter, error) {
	var find pb.Counter
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&find).Error
	if err != nil {
		return pb.Counter{}, err
	}
	return find, nil
}

func (r *MysqlSystemRepository) GetCounterByFlag(ctx context.Context, userId int64, flag string) (pb.Counter, error) {
	var find pb.Counter
	err := r.db.WithContext(ctx).Where("user_id = ? AND flag = ?", userId, flag).First(&find).Error
	if err != nil {
		return pb.Counter{}, err
	}
	return find, nil
}
