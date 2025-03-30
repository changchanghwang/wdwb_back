package application

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/db"
	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	secClient "github.com/changchanghwang/wdwb_back/internal/libs/sec-client"
	"github.com/changchanghwang/wdwb_back/internal/libs/translate"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/commands"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/domain"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/infrastructure"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/response"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type HoldingService struct {
	*ddd.ApplicationService
	holdingRepository infrastructure.HoldingRepository
	secClient         *secClient.SecClient
	translator        *translate.Translator
}

func New(holdingRepository infrastructure.HoldingRepository, translator *translate.Translator, db *gorm.DB, secClient *secClient.SecClient) *HoldingService {
	return &HoldingService{
		holdingRepository:  holdingRepository,
		ApplicationService: &ddd.ApplicationService{Manager: db},
		secClient:          secClient,
		translator:         translator,
	}
}

func (s *HoldingService) List(locale string, command *commands.ListCommand) (*response.HoldingListResponse, error) {
	var (
		holdings []*domain.Holding
		count    int
	)

	var eg errgroup.Group

	conditions := &infrastructure.HoldingQueryConditions{
		InvestorIds: []uuid.UUID{command.InvestorId},
		Years:       []int{command.Year},
		Quarters:    []int{command.Quarter},
	}

	eg.Go(func() error {
		var err error
		holdings, err = s.holdingRepository.Find(nil, conditions, nil, &db.OrderOptions{OrderBy: "value", Direction: "desc"})
		return err
	})
	eg.Go(func() error {
		var err error
		count, err = s.holdingRepository.Count(nil, conditions)
		return err
	})

	if err := eg.Wait(); err != nil {
		return nil, applicationError.Wrap(err)
	}

	res := &response.HoldingListResponse{
		Items: make([]*response.HoldingRetrieveResponse, len(holdings)),
		Count: count,
	}

	for i, holding := range holdings {
		res.Items[i] = &response.HoldingRetrieveResponse{
			Id:         holding.Id,
			InvestorId: holding.InvestorId,
			Name:       s.translator.Translate("companies", locale, holding.Name),
			Year:       holding.Year,
			Quarter:    holding.Quarter,
			Value:      holding.Value,
			Shares:     holding.Shares,
		}
	}

	return res, nil
}
