package plugin

import (
	// Include all plugins.
	_ "github.com/tsundata/assistant/internal/app/chatbot/bot/plugin/any"
	_ "github.com/tsundata/assistant/internal/app/chatbot/bot/plugin/end"
	_ "github.com/tsundata/assistant/internal/app/chatbot/bot/plugin/filter"
	_ "github.com/tsundata/assistant/internal/app/chatbot/bot/plugin/save"
)
