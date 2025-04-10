package application

import (
	"sort"

	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	holdingInfra "github.com/changchanghwang/wdwb_back/internal/services/holdings/infrastructure"
	"github.com/changchanghwang/wdwb_back/internal/services/ranks/commands"
	"github.com/changchanghwang/wdwb_back/internal/services/ranks/domain"
	rankDomainService "github.com/changchanghwang/wdwb_back/internal/services/ranks/domain/services"
	"github.com/changchanghwang/wdwb_back/internal/services/ranks/response"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/changchanghwang/wdwb_back/pkg/util"
	"gorm.io/gorm"
)

type RankService struct {
	*ddd.ApplicationService
	holdingRepository holdingInfra.HoldingRepository
	rankEvaluator     *rankDomainService.RankEvaluator
}

func New(
	holdingRepository holdingInfra.HoldingRepository,
	rankEvaluator *rankDomainService.RankEvaluator,
	db *gorm.DB,
) *RankService {
	return &RankService{
		holdingRepository:  holdingRepository,
		ApplicationService: &ddd.ApplicationService{Manager: db},
		rankEvaluator:      rankEvaluator,
	}
}

func (s *RankService) Rank(language string, command commands.RankCommand) (*response.RankResponse, error) {
	ranks, err := s.rankEvaluator.Evaluate(language, s.Manager, command.Year, command.Quarter)
	if err != nil {
		return nil, applicationError.Wrap(err)
	}

	rankGroups := util.GroupBy(ranks, func(rank *domain.Rank) domain.RankType {
		return rank.Type
	})

	result := &response.RankResponse{}

	for rankType, rankGroup := range rankGroups {
		sort.Slice(rankGroup, func(i, j int) bool {
			return rankGroup[i].Value > rankGroup[j].Value
		})

		switch rankType {
		case domain.TopBuyQuarter:
			result.TopBuyQuarter = rankGroup
		case domain.TopSellQuarter:
			result.TopSellQuarter = rankGroup
		case domain.TopHoldingQuarter:
			result.TopHoldingQuarter = rankGroup
		case domain.TopBuyYear:
			result.TopBuyYear = rankGroup
		case domain.TopSellYear:
			result.TopSellYear = rankGroup
		case domain.TopHoldingYear:
			result.TopHoldingYear = rankGroup
		}
	}

	return result, nil
}
