package global

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/discovery"
	"github.com/tsundata/assistant/internal/pkg/util"
	"strconv"
	"strings"
)

type ID struct {
	conf   *config.AppConfig
	client pb.IdSvcClient
}

func NewID(conf *config.AppConfig, client pb.IdSvcClient) *ID {
	return &ID{
		conf:   conf,
		client: client,
	}
}

func (i *ID) Generate(ctx context.Context) int64 {
	ip := util.GetLocalIP4()
	svcAddr := discovery.SvcAddr(i.conf, i.conf.Name)
	s := strings.Split(svcAddr, ":")
	if len(s) != 2 {
		panic("error service addr")
	}
	port, err := strconv.ParseInt(s[1], 10, 64)
	if err != nil {
		panic(err)
	}
	idReply, err := i.client.GetGlobalId(ctx, &pb.GetGlobalIdRequest{Ip: ip, Port: port})
	if err != nil {
		panic(err)
	}
	return idReply.Id
}
