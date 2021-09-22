package repository

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/util"
)

type OrgRepository interface {
	GetObjectiveByID(ctx context.Context, id int64) (*pb.Objective, error)
	ListObjectives(ctx context.Context, ) ([]*pb.Objective, error)
	CreateObjective(ctx context.Context, objective *pb.Objective) (int64, error)
	DeleteObjective(ctx context.Context, id int64) error
	GetKeyResultByID(ctx context.Context, id int64) (*pb.KeyResult, error)
	ListKeyResults(ctx context.Context, ) ([]*pb.KeyResult, error)
	CreateKeyResult(ctx context.Context, keyResult *pb.KeyResult) (int64, error)
	DeleteKeyResult(ctx context.Context, id int64) error
}

type MysqlOrgRepository struct {
	db *mysql.Conn
}

func NewMysqlOrgRepository(db *mysql.Conn) OrgRepository {
	return &MysqlOrgRepository{db: db}
}

func (r *MysqlOrgRepository) GetObjectiveByID(ctx context.Context, id int64) (*pb.Objective, error) {
	var objective pb.Objective
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&objective).Error
	if err != nil {
		return nil, err
	}
	return &objective, nil
}

func (r *MysqlOrgRepository) ListObjectives(ctx context.Context) ([]*pb.Objective, error) {
	var objectives []*pb.Objective
	err := r.db.WithContext(ctx).Order("id DESC").Find(&objectives).Error
	if err != nil {
		return nil, err
	}
	return objectives, nil
}

func (r *MysqlOrgRepository) CreateObjective(ctx context.Context, objective *pb.Objective) (int64, error) {
	err := r.db.WithContext(ctx).Create(&objective).Error
	if err != nil {
		return 0, err
	}
	return objective.Id, nil
}

func (r *MysqlOrgRepository) DeleteObjective(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.Objective{}).Error
}

func (r *MysqlOrgRepository) GetKeyResultByID(ctx context.Context, id int64) (*pb.KeyResult, error) {
	var keyResult pb.KeyResult
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&keyResult).Error
	if err != nil {
		return nil, err
	}
	return &keyResult, nil
}

func (r *MysqlOrgRepository) ListKeyResults(ctx context.Context) ([]*pb.KeyResult, error) {
	var keyResult []*pb.KeyResult
	err := r.db.WithContext(ctx).Order("id DESC").Find(&keyResult).Error
	if err != nil {
		return nil, err
	}
	return keyResult, nil
}

func (r *MysqlOrgRepository) CreateKeyResult(ctx context.Context, keyResult *pb.KeyResult) (int64, error) {
	err := r.db.WithContext(ctx).Create(&keyResult).Error
	if err != nil {
		return 0, err
	}
	return keyResult.Id, nil
}

func (r *MysqlOrgRepository) DeleteKeyResult(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.KeyResult{}).Error
}

type RqliteOrgRepository struct {
	db *rqlite.Conn
}

func NewRqliteOrgRepository(db *rqlite.Conn) *RqliteOrgRepository {
	return &RqliteOrgRepository{db: db}
}

func (r *RqliteOrgRepository) GetObjectiveByID(id int64) (pb.Objective, error) {
	rows, err := r.db.QueryOne("SELECT * FROM `objectives` WHERE `id` = '%d' LIMIT 1", id)
	if err != nil {
		return pb.Objective{}, nil
	}

	var objective pb.Objective
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.Objective{}, err
		}
		util.Inject(&objective, m)
	}

	return objective, nil
}

func (r *RqliteOrgRepository) ListObjectives() ([]pb.Objective, error) {
	rows, err := r.db.QueryOne("SELECT * FROM `objectives` ORDER BY `id` DESC")
	if err != nil {
		return nil, err
	}

	var objectives []pb.Objective
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return nil, err
		}
		var item pb.Objective
		util.Inject(&item, m)
		objectives = append(objectives, item)
	}

	return objectives, nil
}

func (r *RqliteOrgRepository) CreateObjective(objective pb.Objective) (int64, error) {
	res, err := r.db.WriteOne("INSERT INTO `objectives` (`name`, `tag_id`, `created_at`) VALUES ('%s', %d, '%s')", objective.Name, objective.TagId, util.Now())
	if err != nil {
		return 0, err
	}

	return res.LastInsertID, nil
}

func (r *RqliteOrgRepository) DeleteObjective(id int64) error {
	_, err := r.db.WriteOne("DELETE FROM `objectives` WHERE `id` = '%d'", id)
	return err
}

func (r *RqliteOrgRepository) GetKeyResultByID(id int64) (pb.KeyResult, error) {
	rows, err := r.db.QueryOne("SELECT * FROM `key_results` WHERE `id` = '%d' LIMIT 1", id)
	if err != nil {
		return pb.KeyResult{}, nil
	}

	var keyResult pb.KeyResult
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.KeyResult{}, err
		}
		util.Inject(&keyResult, m)
	}

	return keyResult, nil
}

func (r *RqliteOrgRepository) ListKeyResults() ([]pb.KeyResult, error) {
	rows, err := r.db.QueryOne("SELECT * FROM `key_results` ORDER BY `id` DESC")
	if err != nil {
		return nil, err
	}

	var keyResults []pb.KeyResult
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return nil, err
		}
		var item pb.KeyResult
		util.Inject(&item, m)
		keyResults = append(keyResults, item)
	}

	return keyResults, nil
}

func (r *RqliteOrgRepository) CreateKeyResult(keyResult pb.KeyResult) (int64, error) {
	res, err := r.db.WriteOne("INSERT INTO `key_results` (`objective_id`, `name`, `tag_id`, `complete`, `created_at`, `updated_at`) VALUES (%d, '%s', %d, 0, '%s', '%s')", keyResult.ObjectiveId, keyResult.Name, keyResult.TagId, util.Now(), util.Now())
	if err != nil {
		return 0, err
	}

	return res.LastInsertID, nil
}

func (r *RqliteOrgRepository) DeleteKeyResult(id int64) error {
	_, err := r.db.WriteOne("DELETE FROM `key_results` WHERE `id` = '%d'", id)
	return err
}
