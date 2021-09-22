package listener

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/user/repository"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"go.uber.org/zap"
)

func RegisterEventHandler(ctx context.Context, bus event.Bus, logger log.Logger, repo repository.UserRepository, nlpClient pb.NLPSvcClient) error {
	err := bus.Subscribe(ctx, event.RoleChangeExpSubject, func(msg *nats.Msg) {
		var role pb.Role
		err := json.Unmarshal(msg.Data, &role)
		if err != nil {
			logger.Error(err, zap.Any("event", event.RoleChangeExpSubject))
			return
		}
		err = repo.ChangeRoleExp(ctx, role.UserId, role.Exp)
		if err != nil {
			logger.Error(err, zap.Any("event", event.RoleChangeExpSubject))
			return
		}
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, event.RoleChangeAttrSubject, func(msg *nats.Msg) {
		var data pb.AttrChange
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(err, zap.Any("event", event.RoleChangeAttrSubject))
			return
		}

		res, err := nlpClient.Classifier(ctx, &pb.TextRequest{Text: data.Content})
		if err != nil {
			logger.Error(err, zap.Any("event", event.RoleChangeAttrSubject))
			return
		}

		err = repo.ChangeRoleAttr(ctx, data.UserId, res.Text, 1)
		if err != nil {
			logger.Error(err, zap.Any("event", event.RoleChangeAttrSubject))
			return
		}
	})
	if err != nil {
		return err
	}

	return nil
}
