package migrate

import (
	"embed"
)

//go:embed migrations/*.sql
var Fs embed.FS
