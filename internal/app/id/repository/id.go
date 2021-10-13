package repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"gorm.io/gorm"
)

type IdRepository interface {
	GetOrCreateNode(ctx context.Context, node *pb.Node) (*pb.Node, error)
}

type MysqlIdRepository struct {
	db *mysql.Conn
}

func NewMysqlMiddleRepository(db *mysql.Conn) IdRepository {
	return &MysqlIdRepository{db: db}
}

func (r *MysqlIdRepository) GetOrCreateNode(ctx context.Context, node *pb.Node) (*pb.Node, error) {
	var find pb.Node
	err := r.db.WithContext(ctx).Where("ip = ? AND port = ?", node.Ip, node.Port).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if find.Id <= 0 {
		err = r.db.WithContext(ctx).Create(&node).Error
		if err != nil {
			return nil, err
		}
	}

	return &find, nil
}
