package investors

import (
	"github.com/changchanghwang/wdwb_back/internal/services/investors/application"
	"github.com/changchanghwang/wdwb_back/internal/services/investors/infrastructure"
	"github.com/changchanghwang/wdwb_back/internal/services/investors/presentation"
	"github.com/google/wire"
)

var InvestorSet = wire.NewSet(
	infrastructure.New,
	application.New,
	presentation.New,
)
