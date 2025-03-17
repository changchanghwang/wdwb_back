package holdings

import (
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/infrastructure"
	"github.com/google/wire"
)

var HoldingSet = wire.NewSet(
	infrastructure.New,
)
