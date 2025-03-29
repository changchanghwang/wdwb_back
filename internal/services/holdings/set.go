package holdings

import (
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/application"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/infrastructure"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/presentation"
	"github.com/google/wire"
)

var HoldingSet = wire.NewSet(
	infrastructure.New,
	application.New,
	presentation.New,
)
