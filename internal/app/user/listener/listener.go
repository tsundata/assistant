package listener

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/internal/app/user/repository"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/model"
	"log"
)

func RegisterEventHandler(bus *event.Bus, repo repository.UserRepository) error {
	err := bus.Subscribe(event.ChangeExpSubject, func(msg *nats.Msg) {
		var role model.Role
		err := json.Unmarshal(msg.Data, &role)
		if err != nil {
			log.Println(err)
			return
		}
		err = repo.ChangeRoleExp(role.UserID, role.Exp)
		if err != nil {
			log.Println(err)
			return
		}
	})
	if err != nil {
		return err
	}

	return nil
}
