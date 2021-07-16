package listener

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/api/model"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/user/repository"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
)

func RegisterEventHandler(bus event.Bus, logger log.Logger, repo repository.UserRepository, nlpClient pb.NLPClient) error {
	err := bus.Subscribe(context.Background(), event.ChangeExpSubject, func(msg *nats.Msg) {
		var role model.Role
		err := json.Unmarshal(msg.Data, &role)
		if err != nil {
			logger.Error(err)
			return
		}
		err = repo.ChangeRoleExp(role.UserID, role.Exp)
		if err != nil {
			logger.Error(err)
			return
		}
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(context.Background(), event.ChangeAttrSubject, func(msg *nats.Msg) {
		var data model.AttrChange
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(err)
			return
		}

		res, err := nlpClient.Classifier(context.Background(), &pb.TextRequest{Text: data.Content})
		if err != nil {
			logger.Error(err)
			return
		}

		err = repo.ChangeRoleAttr(data.UserID, res.Text, 1)
		if err != nil {
			logger.Error(err)
			return
		}
	})
	if err != nil {
		return err
	}

	return nil
}
