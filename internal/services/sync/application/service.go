package application

import (
	"fmt"
	"sort"

	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	secClient "github.com/changchanghwang/wdwb_back/internal/libs/sec-client"
	filingDomain "github.com/changchanghwang/wdwb_back/internal/services/filings/domain"
	filingInfra "github.com/changchanghwang/wdwb_back/internal/services/filings/infrastructure"
	holdingDomain "github.com/changchanghwang/wdwb_back/internal/services/holdings/domain"
	holdingInfra "github.com/changchanghwang/wdwb_back/internal/services/holdings/infrastructure"
	investorInfra "github.com/changchanghwang/wdwb_back/internal/services/investors/infrastructure"
	stockDomain "github.com/changchanghwang/wdwb_back/internal/services/stocks/domain"
	stockInfra "github.com/changchanghwang/wdwb_back/internal/services/stocks/infrastructure"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/changchanghwang/wdwb_back/pkg/util"
	"gorm.io/gorm"
)

type SyncService struct {
	ddd.ApplicationService
	secClient          *secClient.SecClient
	investorRepository investorInfra.InvestorRepository
	filingRepository   filingInfra.FilingRepository
	stockRepository    stockInfra.StockRepository
	holdingRepository  holdingInfra.HoldingRepository
}

func New(
	secClient *secClient.SecClient,
	investorRepository investorInfra.InvestorRepository,
	filingRepository filingInfra.FilingRepository,
	stockRepository stockInfra.StockRepository,
	holdingRepository holdingInfra.HoldingRepository,
	db *gorm.DB,
) *SyncService {
	return &SyncService{
		ApplicationService: ddd.ApplicationService{Manager: db},
		secClient:          secClient,
		investorRepository: investorRepository,
		filingRepository:   filingRepository,
		stockRepository:    stockRepository,
		holdingRepository:  holdingRepository,
	}
}

// TODO: event sourcing 기반으로 변경하는게 좋지 않을까?
func (s *SyncService) Sync() error {
	return s.Manager.Transaction(func(tx *gorm.DB) error {

		// investor 조회
		investors, err := s.investorRepository.FindAll(tx)
		if err != nil {
			return applicationError.Wrap(err)
		}

		for _, investor := range investors {
			//investor 별 SEC 파일 목록 조회
			secFilingList, err := s.secClient.GetFilings(investor.Cik)
			if err != nil {
				return applicationError.Wrap(err)
			}

			accessionNumbers := make([]string, len(secFilingList))
			for i, filing := range secFilingList {
				accessionNumbers[i] = filing.AccessionNumber
			}

			// 13F 목록 조회
			filings, err := s.filingRepository.FindByAccessionNumbers(tx, accessionNumbers)
			if err != nil {
				return applicationError.Wrap(err)
			}

			existingAccessionNumbers := make([]string, len(filings))
			for i, filing := range filings {
				existingAccessionNumbers[i] = filing.AccessionNumber
			}

			newSecFilings := make([]*secClient.FilingDTO, 0)
			for _, filing := range secFilingList {
				exist := util.Includes(existingAccessionNumbers, filing.AccessionNumber)
				if !exist {
					newSecFilings = append(newSecFilings, filing)
				}
			}

			newFilings := make([]*filingDomain.Filing, 0)
			newHoldings := make([]*holdingDomain.Holding, 0)

			sort.Slice(newSecFilings, func(i, j int) bool {
				return newSecFilings[i].ReportDate > newSecFilings[j].ReportDate
			})

			for _, newSecFiling := range newSecFilings {
				holdings, err := s.secClient.ParseInfoTable(newSecFiling.InfoTableLink)
				if len(holdings) == 0 {
					fmt.Printf("no parsed holdings for ulr: %s\n", newSecFiling.InfoTableLink)
				}
				if err != nil {
					return applicationError.Wrap(err)
				}

				cusips := make([]string, len(holdings))
				for i, holding := range holdings {
					cusips[i] = holding.Cusip
				}

				stocks, err := s.stockRepository.FindByCusips(tx, cusips)
				if err != nil {
					return applicationError.Wrap(err)
				}

				filing, err := filingDomain.New("13F", newSecFiling.AccessionNumber, newSecFiling.InfoTableLink, newSecFiling.FilingDate, newSecFiling.ReportDate)
				if err != nil {
					return applicationError.Wrap(err)
				}
				newFilings = append(newFilings, filing)

				stockBy := util.KeyBy(stocks, func(stock *stockDomain.Stock) string {
					return stock.Cusip
				})

				holdingGroupByCusip := util.GroupBy(holdings, func(holding *secClient.HoldingDto) string {
					return holding.Cusip
				})

				for cusip, holdings := range holdingGroupByCusip {
					value := 0
					stockShares := 0

					for _, secHolding := range holdings {
						value += secHolding.Value
						stockShares += secHolding.StockShares
					}

					stock := stockBy[cusip]

					if stock != nil {
						holding, err := holdingDomain.New(
							stock.Name,
							stock.Cik,
							cusip,
							investor.Id,
							stock.Id,
							value,
							stockShares,
							filing.Year,
							filing.Quarter,
						)
						if err != nil {
							return applicationError.Wrap(err)
						}

						newHoldings = append(newHoldings, holding)
					}
				}
			}

			if err := s.filingRepository.Save(tx, newFilings); err != nil {
				return applicationError.Wrap(err)
			}

			chunkedHoldings := util.Chunk(newHoldings, 300)

			for _, chunkedHolding := range chunkedHoldings {
				if err := s.holdingRepository.Save(tx, chunkedHolding); err != nil {
					return applicationError.Wrap(err)
				}
			}
		}

		return nil
	})
}
