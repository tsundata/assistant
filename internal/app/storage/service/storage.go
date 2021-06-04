package service

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"go.etcd.io/etcd/clientv3"
	"io"
	"io/ioutil"
)

type Storage struct {
	path string
	etcd *clientv3.Client
	db   *sqlx.DB
	rdb  *redis.Client
}

const MaxFileSize = 1 << 20

func NewStorage(path string, etcd *clientv3.Client, db *sqlx.DB, rdb *redis.Client) *Storage {
	return &Storage{path: path, etcd: etcd, db: db, rdb: rdb}
}

func (s *Storage) UploadFile(stream pb.Storage_UploadFileServer) error {
	fileData := bytes.Buffer{}
	fileSize := 0
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	fileType := req.GetInfo().GetFileType()
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}

		chuck := req.GetChuck()
		size := len(chuck)

		fileSize += size
		if fileSize > MaxFileSize {
			return fmt.Errorf("image is too large: %d > %d", fileSize, MaxFileSize)
		}
		_, err = fileData.Write(chuck)
		if err != nil {
			return err
		}
	}

	// store
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s.%s", uuid, fileType)
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s", s.path, path), fileData.Bytes(), 0644)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&pb.FileReply{Path: path})
}
