package translate

import "github.com/google/wire"

var TranslateSet = wire.NewSet(
	New,
)
