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
	stockInfra "github.com/changchanghwang/wdwb_back/internal/services/stocks/infrastructure"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/changchanghwang/wdwb_back/pkg/util"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type HoldingService struct {
	*ddd.ApplicationService
	holdingRepository infrastructure.HoldingRepository
	stockRepository   stockInfra.StockRepository
	secClient         *secClient.SecClient
	translator        *translate.Translator
}

func New(
	holdingRepository infrastructure.HoldingRepository,
	stockRepository stockInfra.StockRepository,
	translator *translate.Translator,
	db *gorm.DB,
	secClient *secClient.SecClient,
) *HoldingService {
	return &HoldingService{
		holdingRepository:  holdingRepository,
		stockRepository:    stockRepository,
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
		count, err = s.holdingRepository.Count(nil, conditions, &db.FindOptions{GroupBy: "cik"})
		return err
	})

	if err := eg.Wait(); err != nil {
		return nil, applicationError.Wrap(err)
	}

	holdingGroup := util.GroupBy(holdings, func(holding *domain.Holding) string {
		return holding.Cik
	})

	res := &response.HoldingListResponse{
		Items: make([]*response.HoldingRetrieveResponse, 0, len(holdingGroup)),
		Count: count,
	}

	for _, holdings := range holdingGroup {
		firstHolding := holdings[0]
		value := 0
		shares := 0
		for _, holding := range holdings {
			value += holding.Value
			shares += holding.Shares
		}

		res.Items = append(res.Items, &response.HoldingRetrieveResponse{
			Id:         firstHolding.Cik,
			InvestorId: firstHolding.InvestorId,
			Name:       s.translator.Translate("companies", locale, firstHolding.Name, false),
			Year:       firstHolding.Year,
			Quarter:    firstHolding.Quarter,
			Value:      value,
			Shares:     shares,
			Translated: firstHolding.Name != s.translator.Translate("companies", locale, firstHolding.Name, false),
		})
	}

	return res, nil
}
