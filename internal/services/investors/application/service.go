package application

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	"github.com/changchanghwang/wdwb_back/internal/libs/translate"
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
	translator         *translate.Translator
}

func New(
	investorRepository investorInfra.InvestorRepository,
	translator *translate.Translator,
	db *gorm.DB,
) *InvestorService {
	return &InvestorService{
		ApplicationService: ddd.ApplicationService{Manager: db},
		investorRepository: investorRepository,
		translator:         translator,
	}
}

func (s *InvestorService) List(locale string) (*response.InvestorListResponse, error) {
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

	res := &response.InvestorListResponse{
		Items: make([]*response.InvestorRetrieveResponse, len(investors)),
		Count: count,
	}

	for i, investor := range investors {
		res.Items[i] = &response.InvestorRetrieveResponse{
			Id:           investor.Id,
			Name:         s.translator.Translate("investors", locale, investor.Name, true),
			CompanyName:  s.translator.Translate("companies", locale, investor.CompanyName, true),
			Cik:          investor.Cik,
			HoldingValue: investor.HoldingValue,
			Url:          investor.Url,
		}
	}

	return res, nil
}

func (s *InvestorService) Retrieve(locale string, command *command.RetrieveCommand) (*response.InvestorRetrieveResponse, error) {
	investor, err := s.investorRepository.FindOneOrFail(nil, command.Id)
	if err != nil {
		return nil, applicationError.Wrap(err)
	}

	res := &response.InvestorRetrieveResponse{
		Id:           investor.Id,
		Name:         s.translator.Translate("investors", locale, investor.Name, true),
		CompanyName:  s.translator.Translate("companies", locale, investor.CompanyName, true),
		Cik:          investor.Cik,
		HoldingValue: investor.HoldingValue,
		Url:          investor.Url,
	}

	return res, nil
}
