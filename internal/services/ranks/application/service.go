package application

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	holdingInfra "github.com/changchanghwang/wdwb_back/internal/services/holdings/infrastructure"
	"github.com/changchanghwang/wdwb_back/internal/services/ranks/commands"
	rankDomainService "github.com/changchanghwang/wdwb_back/internal/services/ranks/domain/services"
	rankInfra "github.com/changchanghwang/wdwb_back/internal/services/ranks/infrastructure"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"gorm.io/gorm"
)

type RankService struct {
	*ddd.ApplicationService
	rankRepository    rankInfra.RankRepository
	holdingRepository holdingInfra.HoldingRepository
	rankEvaluator     *rankDomainService.RankEvaluator
}

func New(
	rankRepository rankInfra.RankRepository,
	holdingRepository holdingInfra.HoldingRepository,
	rankEvaluator *rankDomainService.RankEvaluator,
	db *gorm.DB,
) *RankService {
	return &RankService{
		rankRepository:     rankRepository,
		holdingRepository:  holdingRepository,
		ApplicationService: &ddd.ApplicationService{Manager: db},
		rankEvaluator:      rankEvaluator,
	}
}

func (s *RankService) Rank(command commands.RankCommand) error {
	return s.Manager.Transaction(func(tx *gorm.DB) error {
		ranks, err := s.rankEvaluator.Evaluate(tx, command.Year, command.Quarter)
		if err != nil {
			return applicationError.Wrap(err)
		}

		err = s.rankRepository.Save(tx, ranks)
		if err != nil {
			return applicationError.Wrap(err)
		}

		return nil
	})
}
