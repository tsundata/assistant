package plugin

import (
	// Include all plugins.
	_ "github.com/tsundata/assistant/internal/app/bot/plugin/any"
	_ "github.com/tsundata/assistant/internal/app/bot/plugin/end"
	_ "github.com/tsundata/assistant/internal/app/bot/plugin/filter"
	_ "github.com/tsundata/assistant/internal/app/bot/plugin/save"
)
