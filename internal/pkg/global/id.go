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

func ID(ctx context.Context, conf *config.AppConfig, client pb.IdSvcClient, service string) int64 {
	ip := util.GetLocalIP4()
	svcAddr := discovery.SvcAddr(conf, service)
	s := strings.Split(svcAddr, ":")
	if len(s) != 2 {
		panic("error service addr")
	}
	port, err := strconv.ParseInt(s[1], 10, 64)
	if err != nil {
		panic(err)
	}
	idReply, err := client.GetGlobalId(ctx, &pb.IdRequest{Ip: ip, Port: port})
	if err != nil {
		panic(err)
	}
	return idReply.Id
}
