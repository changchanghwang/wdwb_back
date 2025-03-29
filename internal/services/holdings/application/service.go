package application

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	secClient "github.com/changchanghwang/wdwb_back/internal/libs/sec-client"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/commands"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/domain"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/infrastructure"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/response"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type HoldingService struct {
	*ddd.ApplicationService
	holdingRepository infrastructure.HoldingRepository
	secClient         *secClient.SecClient
}

func New(holdingRepository infrastructure.HoldingRepository, db *gorm.DB, secClient *secClient.SecClient) *HoldingService {
	return &HoldingService{
		holdingRepository:  holdingRepository,
		ApplicationService: &ddd.ApplicationService{Manager: db},
		secClient:          secClient,
	}
}

func (s *HoldingService) List(command *commands.ListCommand) (*response.ListResponse, error) {
	var (
		holdings []*domain.Holding
		count    int
	)

	var eg errgroup.Group

	eg.Go(func() error {
		var err error
		holdings, err = s.holdingRepository.FindAll(nil)
		return err
	})
	eg.Go(func() error {
		var err error
		count, err = s.holdingRepository.Count(nil)
		return err
	})

	if err := eg.Wait(); err != nil {
		return nil, applicationError.Wrap(err)
	}

	res := &response.ListResponse{
		Items: holdings,
		Count: count,
	}

	return res, nil
}
