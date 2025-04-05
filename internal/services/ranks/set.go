package ranks

import (
	"github.com/changchanghwang/wdwb_back/internal/services/ranks/application"
	"github.com/changchanghwang/wdwb_back/internal/services/ranks/domain/services"
	"github.com/changchanghwang/wdwb_back/internal/services/ranks/infrastructure"
	"github.com/changchanghwang/wdwb_back/internal/services/ranks/presentation"
	"github.com/google/wire"
)

var RankSet = wire.NewSet(
	application.New,
	presentation.New,
	infrastructure.New,
	services.NewRankEvaluator,
)
