package sync

import (
	"github.com/changchanghwang/wdwb_back/internal/services/sync/application"
	"github.com/changchanghwang/wdwb_back/internal/services/sync/presentation"
	"github.com/google/wire"
)

var SyncSet = wire.NewSet(
	presentation.New,
	application.New,
)
