package listener

import (
	"context"
	"encoding/json"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/user/repository"
	"github.com/tsundata/assistant/internal/pkg/event"
)

func RegisterEventHandler(ctx context.Context, bus event.Bus, repo repository.UserRepository, middle pb.MiddleSvcClient) error {
	err := bus.Subscribe(ctx, enum.User, event.RoleChangeExpSubject, func(msg *event.Msg) error {
		var role pb.Role
		err := json.Unmarshal(msg.Data, &role)
		if err != nil {
			return err
		}
		err = repo.ChangeRoleExp(ctx, role.UserId, role.Exp)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.User, event.RoleChangeAttrSubject, func(msg *event.Msg) error {
		var data pb.AttrChange
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}

		res, err := middle.Classifier(ctx, &pb.TextRequest{Text: data.Content})
		if err != nil {
			return err
		}

		err = repo.ChangeRoleAttr(ctx, data.UserId, res.Text, 1)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
