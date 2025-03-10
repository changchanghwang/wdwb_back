package stocks

import (
	"github.com/changchanghwang/wdwb_back/internal/services/stocks/application"
	"github.com/changchanghwang/wdwb_back/internal/services/stocks/infrastructure"
	"github.com/changchanghwang/wdwb_back/internal/services/stocks/presentation"
	"github.com/google/wire"
)

var StockSet = wire.NewSet(
	presentation.New,
	application.New,
	infrastructure.New,
)
