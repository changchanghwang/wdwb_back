package server

import "github.com/google/wire"

var ServerSet = wire.NewSet(
	NewRoute,
	New,
)
