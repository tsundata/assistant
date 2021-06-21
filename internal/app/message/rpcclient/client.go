package rpcclient

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewMiddleClient)
