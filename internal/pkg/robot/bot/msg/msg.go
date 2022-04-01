package msg

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

func BotActionMsg(rule []bot.ActionRule, actionId string) pb.MsgPayload {
	for _, item := range rule {
		if item.ID == actionId {
			var option []string
			for k := range item.OptionFunc {
				option = append(option, k)
			}
			return pb.ActionMsg{
				ID:     item.ID,
				Title:  item.Title,
				Option: option,
				Value:  "",
			}
		}
	}
	return nil
}

func BotFormMsg(rule []bot.FormRule, formId string) pb.MsgPayload {
	for _, item := range rule {
		if item.ID == formId {
			var field []pb.FormField
			for _, fieldItem := range item.Field {
				field = append(field, pb.FormField{
					Key:      fieldItem.Key,
					Type:     string(fieldItem.Type),
					Required: fieldItem.Required,
					Value:    fieldItem.Value,
				})
			}

			return pb.FormMsg{
				ID:    item.ID,
				Title: item.Title,
				Field: field,
			}
		}
	}
	return nil
}
