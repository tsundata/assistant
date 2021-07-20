package repository

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewRqliteMessageRepository)
