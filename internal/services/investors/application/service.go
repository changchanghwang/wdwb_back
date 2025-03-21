package application

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	"github.com/changchanghwang/wdwb_back/internal/services/investors/command"
	"github.com/changchanghwang/wdwb_back/internal/services/investors/domain"
	investorInfra "github.com/changchanghwang/wdwb_back/internal/services/investors/infrastructure"
	"github.com/changchanghwang/wdwb_back/internal/services/investors/response"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type InvestorService struct {
	ddd.ApplicationService
	investorRepository investorInfra.InvestorRepository
}

func New(
	investorRepository investorInfra.InvestorRepository,
	db *gorm.DB,
) *InvestorService {
	return &InvestorService{
		ApplicationService: ddd.ApplicationService{Manager: db},
		investorRepository: investorRepository,
	}
}

func (s *InvestorService) List() (*response.ListResponse, error) {
	var (
		investors []*domain.Investor
		count     int
	)

	var eg errgroup.Group

	eg.Go(func() error {
		var err error
		investors, err = s.investorRepository.FindAll(nil)
		return err
	})
	eg.Go(func() error {
		var err error
		count, err = s.investorRepository.Count(nil)
		return err
	})

	if err := eg.Wait(); err != nil {
		return nil, applicationError.Wrap(err)
	}

	res := &response.ListResponse{
		Items: make([]*response.RetrieveResponse, len(investors)),
		Count: count,
	}

	for i, investor := range investors {
		res.Items[i] = &response.RetrieveResponse{
			Id:           investor.Id.String(),
			Name:         investor.Name,
			CompanyName:  investor.CompanyName,
			Cik:          investor.Cik,
			HoldingValue: investor.HoldingValue,
		}
	}

	return res, nil
}

func (s *InvestorService) Retrieve(command *command.RetrieveCommand) (*response.RetrieveResponse, error) {
	investor, err := s.investorRepository.FindOneOrFail(nil, command.Id)
	if err != nil {
		return nil, applicationError.Wrap(err)
	}

	res := &response.RetrieveResponse{
		Id:           investor.Id.String(),
		Name:         investor.Name,
		CompanyName:  investor.CompanyName,
		Cik:          investor.Cik,
		HoldingValue: investor.HoldingValue,
	}

	return res, nil
}
