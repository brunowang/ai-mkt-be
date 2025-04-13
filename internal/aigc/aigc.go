package aigc

import "github.com/google/wire"

// ProviderSet is aigc providers.
var ProviderSet = wire.NewSet(NewKlingSDK)
