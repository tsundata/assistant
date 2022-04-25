package repository

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
)

type OkrRepository interface {
	GetObjectiveByID(ctx context.Context, id int64) (*pb.Objective, error)
	ListObjectives(ctx context.Context) ([]*pb.Objective, error)
	CreateObjective(ctx context.Context, objective *pb.Objective) (int64, error)
	DeleteObjective(ctx context.Context, id int64) error
	GetKeyResultByID(ctx context.Context, id int64) (*pb.KeyResult, error)
	ListKeyResults(ctx context.Context) ([]*pb.KeyResult, error)
	CreateKeyResult(ctx context.Context, keyResult *pb.KeyResult) (int64, error)
	DeleteKeyResult(ctx context.Context, id int64) error
}

type MysqlOkrRepository struct {
	id *global.ID
	db *mysql.Conn
}

func NewMysqlOkrRepository(id *global.ID, db *mysql.Conn) OkrRepository {
	return &MysqlOkrRepository{id: id, db: db}
}

func (r *MysqlOkrRepository) GetObjectiveByID(ctx context.Context, id int64) (*pb.Objective, error) {
	var objective pb.Objective
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&objective).Error
	if err != nil {
		return nil, err
	}
	return &objective, nil
}

func (r *MysqlOkrRepository) ListObjectives(ctx context.Context) ([]*pb.Objective, error) {
	var objectives []*pb.Objective
	err := r.db.WithContext(ctx).Order("id DESC").Find(&objectives).Error
	if err != nil {
		return nil, err
	}
	return objectives, nil
}

func (r *MysqlOkrRepository) CreateObjective(ctx context.Context, objective *pb.Objective) (int64, error) {
	objective.Id = r.id.Generate(ctx)
	err := r.db.WithContext(ctx).Create(&objective).Error
	if err != nil {
		return 0, err
	}
	return objective.Id, nil
}

func (r *MysqlOkrRepository) DeleteObjective(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.Objective{}).Error
}

func (r *MysqlOkrRepository) GetKeyResultByID(ctx context.Context, id int64) (*pb.KeyResult, error) {
	var keyResult pb.KeyResult
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&keyResult).Error
	if err != nil {
		return nil, err
	}
	return &keyResult, nil
}

func (r *MysqlOkrRepository) ListKeyResults(ctx context.Context) ([]*pb.KeyResult, error) {
	var keyResult []*pb.KeyResult
	err := r.db.WithContext(ctx).Order("id DESC").Find(&keyResult).Error
	if err != nil {
		return nil, err
	}
	return keyResult, nil
}

func (r *MysqlOkrRepository) CreateKeyResult(ctx context.Context, keyResult *pb.KeyResult) (int64, error) {
	keyResult.Id = r.id.Generate(ctx)
	err := r.db.WithContext(ctx).Create(&keyResult).Error
	if err != nil {
		return 0, err
	}
	return keyResult.Id, nil
}

func (r *MysqlOkrRepository) DeleteKeyResult(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.KeyResult{}).Error
}
