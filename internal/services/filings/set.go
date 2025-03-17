package filings

import (
	"github.com/changchanghwang/wdwb_back/internal/services/filings/infrastructure"
	"github.com/google/wire"
)

var FilingSet = wire.NewSet(
	infrastructure.New,
)
