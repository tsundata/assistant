package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/go-redis/redis/v8"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/user/repository"
	"github.com/tsundata/assistant/internal/pkg/util"
	"golang.org/x/image/font/gofont/goregular"
	"image/png"
	"time"
)

const AuthKey = "user:auth:token"

type User struct {
	rdb  *redis.Client
	repo repository.UserRepository
}

func NewUser(rdb *redis.Client, repo repository.UserRepository) *User {
	return &User{rdb: rdb, repo: repo}
}

func (s *User) GetAuthToken(ctx context.Context, _ *pb.TextRequest) (*pb.TextReply, error) {
	var uuid string
	uuid, err := s.rdb.Get(ctx, AuthKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	if errors.Is(err, redis.Nil) {
		uuid, err = util.GenerateUUID()
		if err != nil {
			return nil, err
		}

		status := s.rdb.Set(ctx, AuthKey, uuid, 60*time.Minute)
		if status.Err() != nil {
			return nil, err
		}
	}

	return &pb.TextReply{Text: uuid}, nil
}

func (s *User) Authorization(ctx context.Context, payload *pb.TextRequest) (*pb.StateReply, error) {
	uuid, err := s.rdb.Get(ctx, AuthKey).Result()
	if err != nil {
		return &pb.StateReply{
			State: false,
		}, nil
	}

	return &pb.StateReply{
		State: payload.GetText() == uuid,
	}, nil
}

func (s *User) GetRole(_ context.Context, payload *pb.RoleRequest) (*pb.RoleReply, error) {
	find, err := s.repo.GetRole(int(payload.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.RoleReply{Role: &pb.Role{
		Profession:  find.Profession,
		Exp:         int64(find.Exp),
		Level:       int64(find.Level),
		Strength:    int64(find.Strength),
		Culture:     int64(find.Culture),
		Environment: int64(find.Environment),
		Charisma:    int64(find.Charisma),
		Talent:      int64(find.Talent),
		Intellect:   int64(find.Intellect),
	}}, nil
}

func (s *User) GetRoleImage(_ context.Context, payload *pb.RoleRequest) (*pb.TextReply, error) {
	find, err := s.repo.GetRole(int(payload.GetId()))
	if err != nil {
		return nil, err
	}
	if find.ID <= 0 {
		return nil, errors.New("not role")
	}

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 30})

	dc := gg.NewContext(500, 500)
	dc.SetFontFace(face)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	x := float64(50)
	height := float64(50)
	y := 20 + height
	dc.DrawString(fmt.Sprintf("Profession: %s", find.Profession), x, y)
	y += height
	dc.DrawString(fmt.Sprintf("Level: %d    Exp: %d", find.Level, find.Exp), x, y)
	y += height
	dc.DrawString(fmt.Sprintf("Strength: %d", find.Strength), x, y)
	y += height
	dc.DrawString(fmt.Sprintf("Culture: %d", find.Culture), x, y)
	y += height
	dc.DrawString(fmt.Sprintf("Environment: %d", find.Environment), x, y)
	y += height
	dc.DrawString(fmt.Sprintf("Charisma: %d", find.Charisma), x, y)
	y += height
	dc.DrawString(fmt.Sprintf("Talent: %d", find.Talent), x, y)
	y += height
	dc.DrawString(fmt.Sprintf("Intellect: %d", find.Intellect), x, y)

	buf := new(bytes.Buffer)
	err = png.Encode(buf, dc.Image())
	if err != nil {
		return nil, err
	}

	return &pb.TextReply{Text: util.ImageToBase64(buf.Bytes())}, nil
}
