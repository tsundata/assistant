package global

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/discovery"
	"github.com/tsundata/assistant/internal/pkg/util"
	"strconv"
	"strings"
)

func Id(ctx context.Context, middle pb.MiddleSvcClient, conf *config.AppConfig, service string) (int64, error) {
	ip := util.GetLocalIP4()
	svcAddr := discovery.SvcAddr(conf, service)
	s := strings.Split(svcAddr, ":")
	if len(s) != 2 {
		return 0, errors.New("error service addr")
	}
	port, err := strconv.ParseInt(s[1], 10, 64)
	if err != nil {
		return 0, err
	}
	idReply, err := middle.GetGlobalId(ctx, &pb.IdRequest{Ip: ip, Port: port})
	if err != nil {
		return 0, err
	}
	return idReply.Id, nil
}
