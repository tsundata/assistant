package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"gorm.io/gorm"
	"time"
)

type OkrRepository interface {
	GetObjectiveByID(ctx context.Context, id int64) (*pb.Objective, error)
	GetObjectiveBySequence(ctx context.Context, userId, sequence int64) (*pb.Objective, error)
	ListObjectives(ctx context.Context, userId int64) ([]*pb.Objective, error)
	CreateObjective(ctx context.Context, objective *pb.Objective) (int64, error)
	UpdateObjective(ctx context.Context, objective *pb.Objective) error
	DeleteObjective(ctx context.Context, id int64) error
	DeleteObjectiveBySequence(ctx context.Context, userId, sequence int64) error
	GetKeyResultByID(ctx context.Context, id int64) (*pb.KeyResult, error)
	GetKeyResultBySequence(ctx context.Context, userId, sequence int64) (*pb.KeyResult, error)
	ListKeyResults(ctx context.Context, userId int64) ([]*pb.KeyResult, error)
	ListKeyResultsById(ctx context.Context, id []int64) ([]*pb.KeyResult, error)
	ListKeyResultsByObjectiveId(ctx context.Context, objectiveId int64) ([]*pb.KeyResult, error)
	CreateKeyResult(ctx context.Context, keyResult *pb.KeyResult) (int64, error)
	UpdateKeyResult(ctx context.Context, keyResult *pb.KeyResult) error
	DeleteKeyResult(ctx context.Context, id int64) error
	DeleteKeyResultBySequence(ctx context.Context, userId, sequence int64) error
	AggregateObjectiveValue(ctx context.Context, id int64) error
	AggregateKeyResultValue(ctx context.Context, id int64) error
	CreateKeyResultValue(ctx context.Context, keyResultValue *pb.KeyResultValue) (int64, error)
	GetKeyResultValues(ctx context.Context, keyResultId int64) ([]*pb.KeyResultValue, error)
}

type MysqlOkrRepository struct {
	locker *global.Locker
	id     *global.ID
	db     *mysql.Conn
}

func NewMysqlOkrRepository(locker *global.Locker, id *global.ID, db *mysql.Conn) OkrRepository {
	return &MysqlOkrRepository{locker: locker, id: id, db: db}
}

func (r *MysqlOkrRepository) GetObjectiveByID(ctx context.Context, id int64) (*pb.Objective, error) {
	var objective pb.Objective
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&objective).Error
	if err != nil {
		return nil, err
	}
	return &objective, nil
}

func (r *MysqlOkrRepository) GetObjectiveBySequence(ctx context.Context, userId, sequence int64) (*pb.Objective, error) {
	var objective pb.Objective
	err := r.db.WithContext(ctx).Where("user_id = ? AND sequence = ?", userId, sequence).First(&objective).Error
	if err != nil {
		return nil, err
	}
	return &objective, nil
}

func (r *MysqlOkrRepository) ListObjectives(ctx context.Context, userId int64) ([]*pb.Objective, error) {
	var objectives []*pb.Objective
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Order("id DESC").Find(&objectives).Error
	if err != nil {
		return nil, err
	}
	return objectives, nil
}

func (r *MysqlOkrRepository) CreateObjective(ctx context.Context, objective *pb.Objective) (int64, error) {
	l, err := r.locker.Acquire(fmt.Sprintf("chatbot:objective:create:%d", objective.UserId))
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = l.Release()
	}()

	// sequence
	sequence := int64(0)
	var max pb.Objective
	err = r.db.WithContext(ctx).Where("user_id = ?", objective.UserId).Order("sequence DESC").Take(&max).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	if max.Sequence > 0 {
		sequence = max.Sequence
	}
	sequence += 1

	objective.Id = r.id.Generate(ctx)
	objective.Sequence = sequence
	objective.CreatedAt = time.Now().Unix()
	objective.UpdatedAt = time.Now().Unix()
	err = r.db.WithContext(ctx).Create(&objective).Error
	if err != nil {
		return 0, err
	}
	return objective.Id, nil
}

func (r *MysqlOkrRepository) UpdateObjective(ctx context.Context, objective *pb.Objective) error {
	return r.db.WithContext(ctx).Model(&pb.Objective{}).
		Where("user_id = ? AND sequence = ?", objective.UserId, objective.Sequence).
		UpdateColumns(map[string]interface{}{
			"title":       objective.Title,
			"memo":        objective.Memo,
			"motive":      objective.Motive,
			"feasibility": objective.Feasibility,
			"is_plan":     objective.IsPlan,
			"plan_start":  objective.PlanStart,
			"plan_end":    objective.PlanEnd,
			"updated_at":  time.Now().Unix(),
		}).Error
}

func (r *MysqlOkrRepository) DeleteObjective(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.Objective{}).Error
}

func (r *MysqlOkrRepository) DeleteObjectiveBySequence(ctx context.Context, userId, sequence int64) error {
	return r.db.WithContext(ctx).Where("user_id = ? AND sequence = ?", userId, sequence).Delete(&pb.Objective{}).Error
}

func (r *MysqlOkrRepository) GetKeyResultByID(ctx context.Context, id int64) (*pb.KeyResult, error) {
	var keyResult pb.KeyResult
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&keyResult).Error
	if err != nil {
		return nil, err
	}
	return &keyResult, nil
}

func (r *MysqlOkrRepository) GetKeyResultBySequence(ctx context.Context, userId, sequence int64) (*pb.KeyResult, error) {
	var keyResult pb.KeyResult
	err := r.db.WithContext(ctx).Where("user_id = ? AND sequence = ?", userId, sequence).First(&keyResult).Error
	if err != nil {
		return nil, err
	}
	return &keyResult, nil
}

func (r *MysqlOkrRepository) ListKeyResults(ctx context.Context, userId int64) ([]*pb.KeyResult, error) {
	var keyResult []*pb.KeyResult
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Order("id DESC").Find(&keyResult).Error
	if err != nil {
		return nil, err
	}
	return keyResult, nil
}

func (r *MysqlOkrRepository) ListKeyResultsById(ctx context.Context, id []int64) ([]*pb.KeyResult, error) {
	var keyResult []*pb.KeyResult
	err := r.db.WithContext(ctx).Where("id IN ?", id).Order("id DESC").Find(&keyResult).Error
	if err != nil {
		return nil, err
	}
	return keyResult, nil
}

func (r *MysqlOkrRepository) ListKeyResultsByObjectiveId(ctx context.Context, objectiveId int64) ([]*pb.KeyResult, error) {
	var keyResult []*pb.KeyResult
	err := r.db.WithContext(ctx).Where("objective_id = ?", objectiveId).Order("id DESC").Find(&keyResult).Error
	if err != nil {
		return nil, err
	}
	return keyResult, nil
}

func (r *MysqlOkrRepository) CreateKeyResult(ctx context.Context, keyResult *pb.KeyResult) (int64, error) {
	l, err := r.locker.Acquire(fmt.Sprintf("chatbot:key_result:create:%d", keyResult.UserId))
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = l.Release()
	}()

	// sequence
	sequence := int64(0)
	var max pb.KeyResult
	err = r.db.WithContext(ctx).Where("user_id = ?", keyResult.UserId).Order("sequence DESC").Take(&max).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	if max.Sequence > 0 {
		sequence = max.Sequence
	}
	sequence += 1

	keyResult.Id = r.id.Generate(ctx)
	keyResult.Sequence = sequence
	keyResult.CreatedAt = time.Now().Unix()
	keyResult.UpdatedAt = time.Now().Unix()
	err = r.db.WithContext(ctx).Create(&keyResult).Error
	if err != nil {
		return 0, err
	}

	// init value record
	if keyResult.CurrentValue > 0 {
		err = r.db.WithContext(ctx).Create(&pb.KeyResultValue{
			Id:          r.id.Generate(ctx),
			KeyResultId: keyResult.Id,
			Value:       keyResult.CurrentValue,
			CreatedAt:   time.Now().Unix(),
		}).Error
		if err != nil {
			return 0, err
		}
	}

	return keyResult.Id, nil
}

func (r *MysqlOkrRepository) UpdateKeyResult(ctx context.Context, keyResult *pb.KeyResult) error {
	return r.db.WithContext(ctx).Model(&pb.KeyResult{}).
		Where("user_id = ? AND sequence = ?", keyResult.UserId, keyResult.Sequence).
		UpdateColumns(map[string]interface{}{
			"title":        keyResult.Title,
			"memo":         keyResult.Memo,
			"target_value": keyResult.TargetValue,
			"value_mode":   keyResult.ValueMode,
			"updated_at":   time.Now().Unix(),
		}).Error
}

func (r *MysqlOkrRepository) DeleteKeyResult(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&pb.KeyResult{}).Error
}

func (r *MysqlOkrRepository) DeleteKeyResultBySequence(ctx context.Context, userId, sequence int64) error {
	return r.db.WithContext(ctx).Where("user_id = ? AND sequence = ?", userId, sequence).Delete(&pb.KeyResult{}).Error
}

func (r *MysqlOkrRepository) AggregateObjectiveValue(ctx context.Context, id int64) error {
	result := pb.KeyResult{}
	err := r.db.WithContext(ctx).Model(&pb.KeyResult{}).Where("objective_id = ?", id).
		Select("SUM(current_value) as current_value, SUM(target_value) as target_value").Take(&result).Error
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Model(&pb.Objective{}).Where("id = ?", id).UpdateColumns(map[string]interface{}{
		"current_value": result.CurrentValue,
		"total_value":   result.TargetValue,
		"updated_at":    time.Now().Unix(),
	}).Error
}

func (r *MysqlOkrRepository) AggregateKeyResultValue(ctx context.Context, id int64) error {
	keyResult, err := r.GetKeyResultByID(ctx, id)
	if err != nil {
		return err
	}
	var value sql.NullInt64
	switch keyResult.ValueMode {
	case enum.ValueSumMode:
		err = r.db.WithContext(ctx).Model(&pb.KeyResultValue{}).Where("key_result_id = ?", id).
			Select("SUM(value) as value").Pluck("value", &value).Error
	case enum.ValueLastMode:
		err = r.db.WithContext(ctx).Model(&pb.KeyResultValue{}).Where("key_result_id = ?", id).
			Order("created_at DESC").Limit(1).Pluck("value", &value).Error
	case enum.ValueAvgMode:
		err = r.db.WithContext(ctx).Model(&pb.KeyResultValue{}).Where("key_result_id = ?", id).
			Select("AVG(value) as value").Pluck("value", &value).Error
	case enum.ValueMaxMode:
		err = r.db.WithContext(ctx).Model(&pb.KeyResultValue{}).Where("key_result_id = ?", id).
			Select("MAX(value) as value").Pluck("value", &value).Error
	}
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Model(&pb.KeyResult{}).Where("id = ?", id).UpdateColumns(map[string]interface{}{
		"current_value": value.Int64,
		"updated_at":    time.Now().Unix(),
	}).Error
}

func (r *MysqlOkrRepository) CreateKeyResultValue(ctx context.Context, keyResultValue *pb.KeyResultValue) (int64, error) {
	keyResultValue.Id = r.id.Generate(ctx)
	keyResultValue.CreatedAt = time.Now().Unix()
	err := r.db.WithContext(ctx).Create(&keyResultValue).Error
	if err != nil {
		return 0, err
	}
	return keyResultValue.Id, nil
}

func (r *MysqlOkrRepository) GetKeyResultValues(ctx context.Context, keyResultId int64) ([]*pb.KeyResultValue, error) {
	var values []*pb.KeyResultValue
	err := r.db.WithContext(ctx).Where("key_result_id = ?", keyResultId).Order("id DESC").Find(&values).Error
	if err != nil {
		return nil, err
	}
	return values, nil
}
