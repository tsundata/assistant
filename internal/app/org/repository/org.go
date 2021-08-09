package repository

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/util"
)

type OrgRepository interface {
	GetObjectiveByID(id int64) (pb.Objective, error)
	ListObjectives() ([]pb.Objective, error)
	CreateObjective(objective pb.Objective) (int64, error)
	DeleteObjective(id int64) error
	GetKeyResultByID(id int64) (pb.KeyResult, error)
	ListKeyResults() ([]pb.KeyResult, error)
	CreateKeyResult(keyResult pb.KeyResult) (int64, error)
	DeleteKeyResult(id int64) error
}

type RqliteOrgRepository struct {
	db *rqlite.Conn
}

func NewRqliteOrgRepository(db *rqlite.Conn) OrgRepository {
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
	res, err := r.db.WriteOne("INSERT INTO `objectives` (`name`, `tag`, `created_at`) VALUES ('%s', '%s', '%s')", objective.Name, objective.Tag, util.Now())
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
	res, err := r.db.WriteOne("INSERT INTO `key_results` (`objective_id`, `name`, `tag`, `complete`, `created_at`, `updated_at`) VALUES (%d, '%s', '%s', 0, '%s', '%s')", keyResult.ObjectiveId, keyResult.Name, keyResult.Tag, util.Now(), util.Now())
	if err != nil {
		return 0, err
	}

	return res.LastInsertID, nil
}

func (r *RqliteOrgRepository) DeleteKeyResult(id int64) error {
	_, err := r.db.WriteOne("DELETE FROM `key_results` WHERE `id` = '%d'", id)
	return err
}
