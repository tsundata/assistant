package keyword

import (
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

func init() {
	bot.RegisterPlugin("keyword", Keyword{})
}
