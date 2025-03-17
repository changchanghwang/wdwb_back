package investors

import (
	"github.com/changchanghwang/wdwb_back/internal/services/investors/infrastructure"
	"github.com/google/wire"
)

var InvestorSet = wire.NewSet(
	infrastructure.New,
)
