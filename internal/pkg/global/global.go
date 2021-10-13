package global

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewID, NewLocker)
