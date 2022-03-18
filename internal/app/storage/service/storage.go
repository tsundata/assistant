package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/storage/fs"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/util"
	"io"
)

type Storage struct {
	conf *config.AppConfig
	rdb  *redis.Client
}

const MaxFileSize = 1 << 20

func NewStorage(conf *config.AppConfig, rdb *redis.Client) *Storage {
	return &Storage{conf: conf, rdb: rdb}
}

func (s *Storage) UploadFile(stream pb.StorageSvc_UploadFileServer) error {
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
	uuid := util.UUID()

	f, err := fs.FS(s.conf)
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s.%s", uuid, fileType)
	err = f.Put(path, fileData.Bytes(), false)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&pb.FileReply{Path: f.FullPath(path)})
}

func (s *Storage) AbsolutePath(_ context.Context, payload *pb.TextRequest) (*pb.TextReply, error) {
	f, err := fs.FS(s.conf)
	if err != nil {
		return nil, err
	}
	return &pb.TextReply{Text: f.AbsolutePath(payload.Text)}, nil
}
