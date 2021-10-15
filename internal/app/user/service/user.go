package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/user/repository"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/image/font/gofont/goregular"
	"gorm.io/gorm"
	"image/png"
	"time"
)

type User struct {
	conf   *config.AppConfig
	logger log.Logger
	rdb    *redis.Client
	repo   repository.UserRepository
}

func NewUser(conf *config.AppConfig, logger log.Logger, rdb *redis.Client, repo repository.UserRepository) *User {
	return &User{conf: conf, logger: logger, rdb: rdb, repo: repo}
}

func (s *User) Login(ctx context.Context, payload *pb.LoginRequest) (*pb.AuthReply, error) {
	find, err := s.repo.GetByName(ctx, payload.Username)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Error(err)
		}
		return &pb.AuthReply{State: false}, errors.New("username/password invalid")
	}

	err = bcrypt.CompareHashAndPassword(util.StringToByte(find.Password), util.StringToByte(payload.Password))
	if err != nil {
		s.logger.Error(err)
		return &pb.AuthReply{State: false}, errors.New("username/password invalid")
	}

	reply, err := s.GetAuthToken(ctx, &pb.AuthRequest{Id: find.Id})
	if err != nil {
		s.logger.Error(err)
		return &pb.AuthReply{State: false}, errors.New("username/password invalid")
	}

	return &pb.AuthReply{State: true, Token: reply.Token}, nil
}

func (s *User) GetAuthToken(_ context.Context, payload *pb.AuthRequest) (*pb.AuthReply, error) {
	// jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  payload.Id,
		"nbf": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(util.StringToByte(s.conf.Jwt.Secret))
	if err != nil {
		return nil, err
	}

	return &pb.AuthReply{Token: tokenString}, nil
}

func (s *User) Authorization(_ context.Context, payload *pb.AuthRequest) (*pb.AuthReply, error) {
	// jwt
	token, err := jwt.Parse(payload.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return util.StringToByte(s.conf.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &pb.AuthReply{
			State: true,
			Id:    int64(claims["id"].(float64)),
		}, nil
	}

	return &pb.AuthReply{
		State: false,
	}, nil
}

func (s *User) GetRole(ctx context.Context, payload *pb.RoleRequest) (*pb.RoleReply, error) {
	find, err := s.repo.GetRole(ctx, int(payload.Id))
	if err != nil {
		return nil, err
	}
	return &pb.RoleReply{Role: &pb.Role{
		Profession:  find.Profession,
		Exp:         find.Exp,
		Level:       find.Level,
		Strength:    find.Strength,
		Culture:     find.Culture,
		Environment: find.Environment,
		Charisma:    find.Charisma,
		Talent:      find.Talent,
		Intellect:   find.Intellect,
	}}, nil
}

func (s *User) GetRoleImage(ctx context.Context, payload *pb.RoleRequest) (*pb.BytesReply, error) {
	find, err := s.repo.GetRole(ctx, int(payload.Id))
	if err != nil {
		return nil, err
	}
	if find.Id <= 0 {
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

	return &pb.BytesReply{Data: buf.Bytes()}, nil
}

func (s *User) CreateUser(ctx context.Context, payload *pb.UserRequest) (*pb.UserReply, error) {
	user := &pb.User{
		Nickname: payload.User.GetNickname(),
		Mobile:   payload.User.GetMobile(),
		Remark:   payload.User.GetRemark(),
	}
	id, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.UserReply{User: &pb.User{
		Id: id,
	}}, nil
}

func (s *User) GetUser(ctx context.Context, payload *pb.UserRequest) (*pb.UserReply, error) {
	find, err := s.repo.GetByID(ctx, payload.User.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.UserReply{
		User: &pb.User{
			Id:       find.Id,
			Nickname: find.Nickname,
			Mobile:   find.Mobile,
			Remark:   find.Remark,
		},
	}, nil
}

func (s *User) GetUserByName(ctx context.Context, payload *pb.UserRequest) (*pb.UserReply, error) {
	find, err := s.repo.GetByName(ctx, payload.User.GetNickname())
	if err != nil {
		return nil, err
	}

	return &pb.UserReply{
		User: &pb.User{
			Id:       find.Id,
			Nickname: find.Nickname,
			Mobile:   find.Mobile,
			Remark:   find.Remark,
		},
	}, nil
}

func (s *User) GetUsers(ctx context.Context, _ *pb.UserRequest) (*pb.UsersReply, error) {
	items, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var res []*pb.User
	for _, item := range items {
		res = append(res, &pb.User{
			Id:       item.Id,
			Nickname: item.Nickname,
			Mobile:   item.Mobile,
			Remark:   item.Remark,
		})
	}

	return &pb.UsersReply{Users: res}, nil
}

func (s *User) UpdateUser(ctx context.Context, payload *pb.UserRequest) (*pb.StateReply, error) {
	err := s.repo.Update(ctx, &pb.User{
		Id:       payload.User.GetId(),
		Nickname: payload.User.GetNickname(),
		Mobile:   payload.User.GetMobile(),
		Remark:   payload.User.GetRemark(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *User) BindDevice(ctx context.Context, payload *pb.DeviceRequest) (*pb.StateReply, error) {
	devices, err := s.repo.ListDevice(ctx, payload.Device.GetUserId())
	if err != nil {
		return nil, err
	}
	isBind := false
	for _, item := range devices {
		if item.Name == payload.Device.GetName() {
			isBind = true
		}
	}
	if isBind {
		return &pb.StateReply{State: true}, nil
	}
	_, err = s.repo.CreateDevice(ctx, payload.Device)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}
