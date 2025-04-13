package lib

import "github.com/google/wire"

// ProviderSet is lib providers.
var ProviderSet = wire.NewSet(NewS3Mgr)
