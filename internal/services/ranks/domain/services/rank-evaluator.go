package services

import (
	"slices"
	"sort"

	holdingDomain "github.com/changchanghwang/wdwb_back/internal/services/holdings/domain"
	holdingInfra "github.com/changchanghwang/wdwb_back/internal/services/holdings/infrastructure"
	rank "github.com/changchanghwang/wdwb_back/internal/services/ranks/domain"
	stockDomain "github.com/changchanghwang/wdwb_back/internal/services/stocks/domain"
	stockInfra "github.com/changchanghwang/wdwb_back/internal/services/stocks/infrastructure"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/changchanghwang/wdwb_back/pkg/util"
	"gorm.io/gorm"
)

type RankEvaluator struct {
	holdingRepository holdingInfra.HoldingRepository
	stockRepository   stockInfra.StockRepository
}

func NewRankEvaluator(holdingRepository holdingInfra.HoldingRepository, stockRepository stockInfra.StockRepository) *RankEvaluator {
	return &RankEvaluator{holdingRepository: holdingRepository, stockRepository: stockRepository}
}

// TODO: idempotency
func (r *RankEvaluator) Evaluate(db *gorm.DB, year int, quarter int) ([]*rank.Rank, error) {
	var result []*rank.Rank
	if quarter != 0 {
		lastQuarterYear := year
		lastQuarter := quarter - 1
		if lastQuarter == 0 {
			lastQuarter = 4
			lastQuarterYear = year - 1
		}

		thisQuarterHoldings, err := r.holdingRepository.Find(db, &holdingInfra.HoldingQueryConditions{Years: []int{year}, Quarters: []int{quarter}}, nil, nil)
		if err != nil {
			return nil, applicationError.Wrap(err)
		}

		lastQuarterHoldings, err := r.holdingRepository.Find(db, &holdingInfra.HoldingQueryConditions{Years: []int{lastQuarterYear}, Quarters: []int{lastQuarter}}, nil, nil)
		if err != nil {
			return nil, applicationError.Wrap(err)
		}

		ranks, err := r.evaluateTopBuyQuarter(db, thisQuarterHoldings, lastQuarterHoldings)
		if err != nil {
			return nil, applicationError.Wrap(err)
		}
		result = append(result, ranks...)
		ranks, err = r.evaluateTopSellQuarter(db, thisQuarterHoldings, lastQuarterHoldings)
		if err != nil {
			return nil, applicationError.Wrap(err)
		}
		result = append(result, ranks...)
		ranks, err = r.evaluateTopHoldingQuarter(db, thisQuarterHoldings)
		if err != nil {
			return nil, applicationError.Wrap(err)
		}
		result = append(result, ranks...)
	} else {
		r.evaluateTopBuyYear(db, year)
		r.evaluateTopSellYear(db, year)
		r.evaluateTopHoldingYear(db, year)
	}

	return result, nil
}

type RankParam struct {
	Cik     string
	Year    int
	Quarter int
	Value   int
	Shares  int
}

func (r *RankEvaluator) evaluateTopBuyQuarter(db *gorm.DB, thisQuarterHoldings, lastQuarterHoldings []*holdingDomain.Holding) (ranks []*rank.Rank, err error) {
	thisQuarterHoldingGroups := util.GroupBy(thisQuarterHoldings, func(holding *holdingDomain.Holding) string {
		return holding.Cik
	})

	lastQuarterHoldingGroups := util.GroupBy(lastQuarterHoldings, func(holding *holdingDomain.Holding) string {
		return holding.Cik
	})

	rankParams := make([]*RankParam, 0)

	for _, holdingGroup := range thisQuarterHoldingGroups {
		var totalValue int
		var totalShares int
		for _, holding := range holdingGroup {
			totalValue += holding.Value
			totalShares += holding.Shares
		}

		lastHoldingGroup := lastQuarterHoldingGroups[holdingGroup[0].Cik]
		rankParam := &RankParam{
			Cik:     holdingGroup[0].Cik,
			Year:    thisQuarterHoldings[0].Year,
			Quarter: thisQuarterHoldings[0].Quarter,
			Value:   totalValue,
			Shares:  totalShares,
		}

		if len(lastHoldingGroup) == 0 {
			rankParams = append(rankParams, rankParam)
			continue
		}

		var lastValue int
		var lastShares int

		for _, lastHolding := range lastHoldingGroup {
			lastValue += lastHolding.Value
			lastShares += lastHolding.Shares
		}

		rankParam.Value -= lastValue
		rankParam.Shares -= lastShares

		rankParams = append(rankParams, rankParam)
	}

	sort.Slice(rankParams, func(i, j int) bool {
		return rankParams[i].Value > rankParams[j].Value
	})

	rankParams = rankParams[:10]

	stocks, err := r.stockRepository.Find(db, &stockInfra.StockQueryConditions{
		Ciks: util.Map(rankParams, func(rankParam *RankParam) string {
			return rankParam.Cik
		}),
	}, nil, nil)

	if err != nil {
		return nil, applicationError.Wrap(err)
	}

	stockGroups := util.GroupBy(stocks, func(stock *stockDomain.Stock) string {
		return stock.Cik
	})

	for i, rankParam := range rankParams {
		stockGroup := stockGroups[rankParam.Cik]
		tickers := make([]string, 0)
		for _, stock := range stockGroup {
			if slices.Contains(tickers, stock.Ticker) {
				continue
			}
			tickers = append(tickers, stock.Ticker)
		}

		rank, err := rank.New(
			rank.TopBuyQuarter,
			thisQuarterHoldings[0].Year,
			thisQuarterHoldings[0].Quarter,
			i+1,
			rankParam.Value,
			tickers,
			stockGroup[0].Name,
			"*",
		)

		if err != nil {
			return nil, applicationError.Wrap(err)
		}

		ranks = append(ranks, rank)
	}

	return ranks, nil
}

func (r *RankEvaluator) evaluateTopSellQuarter(db *gorm.DB, thisQuarterHoldings, lastQuarterHoldings []*holdingDomain.Holding) (ranks []*rank.Rank, err error) {
	thisQuarterHoldingGroups := util.GroupBy(thisQuarterHoldings, func(holding *holdingDomain.Holding) string {
		return holding.Cik
	})

	lastQuarterHoldingGroups := util.GroupBy(lastQuarterHoldings, func(holding *holdingDomain.Holding) string {
		return holding.Cik
	})

	rankParams := make([]*RankParam, 0)

	for _, holdingGroup := range thisQuarterHoldingGroups {
		var totalValue int
		var totalShares int
		for _, holding := range holdingGroup {
			totalValue += holding.Value
			totalShares += holding.Shares
		}

		lastHoldingGroup := lastQuarterHoldingGroups[holdingGroup[0].Cik]
		rankParam := &RankParam{
			Cik:     holdingGroup[0].Cik,
			Year:    thisQuarterHoldings[0].Year,
			Quarter: thisQuarterHoldings[0].Quarter,
			Value:   totalValue,
			Shares:  totalShares,
		}

		if len(lastHoldingGroup) == 0 {
			rankParams = append(rankParams, rankParam)
			continue
		}

		var lastValue int
		var lastShares int

		for _, lastHolding := range lastHoldingGroup {
			lastValue += lastHolding.Value
			lastShares += lastHolding.Shares
		}

		rankParam.Value -= lastValue
		rankParam.Shares -= lastShares

		rankParams = append(rankParams, rankParam)
	}

	sort.Slice(rankParams, func(i, j int) bool {
		return rankParams[i].Value < rankParams[j].Value
	})

	rankParams = rankParams[:10]

	stocks, err := r.stockRepository.Find(db, &stockInfra.StockQueryConditions{
		Ciks: util.Map(rankParams, func(rankParam *RankParam) string {
			return rankParam.Cik
		}),
	}, nil, nil)

	if err != nil {
		return nil, applicationError.Wrap(err)
	}

	stockGroups := util.GroupBy(stocks, func(stock *stockDomain.Stock) string {
		return stock.Cik
	})

	for i, rankParam := range rankParams {
		stockGroup := stockGroups[rankParam.Cik]
		tickers := make([]string, 0)
		for _, stock := range stockGroup {
			if slices.Contains(tickers, stock.Ticker) {
				continue
			}
			tickers = append(tickers, stock.Ticker)
		}

		rank, err := rank.New(
			rank.TopSellQuarter,
			thisQuarterHoldings[0].Year,
			thisQuarterHoldings[0].Quarter,
			i+1,
			rankParam.Value,
			tickers,
			stockGroup[0].Name,
			"*",
		)

		if err != nil {
			return nil, applicationError.Wrap(err)
		}

		ranks = append(ranks, rank)
	}

	return ranks, nil
}

func (r *RankEvaluator) evaluateTopHoldingQuarter(db *gorm.DB, thisQuarterHoldings []*holdingDomain.Holding) (ranks []*rank.Rank, err error) {
	thisQuarterHoldingGroups := util.GroupBy(thisQuarterHoldings, func(holding *holdingDomain.Holding) string {
		return holding.Cik
	})

	rankParams := make([]*RankParam, 0)

	for _, holdingGroup := range thisQuarterHoldingGroups {
		var totalValue int
		var totalShares int
		for _, holding := range holdingGroup {
			totalValue += holding.Value
			totalShares += holding.Shares
		}

		rankParam := &RankParam{
			Cik:     holdingGroup[0].Cik,
			Year:    thisQuarterHoldings[0].Year,
			Quarter: thisQuarterHoldings[0].Quarter,
			Value:   totalValue,
			Shares:  totalShares,
		}

		rankParams = append(rankParams, rankParam)
	}

	sort.Slice(rankParams, func(i, j int) bool {
		return rankParams[i].Value > rankParams[j].Value
	})

	rankParams = rankParams[:10]

	stocks, err := r.stockRepository.Find(db, &stockInfra.StockQueryConditions{
		Ciks: util.Map(rankParams, func(rankParam *RankParam) string {
			return rankParam.Cik
		}),
	}, nil, nil)

	if err != nil {
		return nil, applicationError.Wrap(err)
	}

	stockGroups := util.GroupBy(stocks, func(stock *stockDomain.Stock) string {
		return stock.Cik
	})

	for i, rankParam := range rankParams {
		stockGroup := stockGroups[rankParam.Cik]
		tickers := make([]string, 0)
		for _, stock := range stockGroup {
			if slices.Contains(tickers, stock.Ticker) {
				continue
			}
			tickers = append(tickers, stock.Ticker)
		}

		rank, err := rank.New(
			rank.TopSellQuarter,
			thisQuarterHoldings[0].Year,
			thisQuarterHoldings[0].Quarter,
			i+1,
			rankParam.Value,
			tickers,
			stockGroup[0].Name,
			"*",
		)

		if err != nil {
			return nil, applicationError.Wrap(err)
		}

		ranks = append(ranks, rank)
	}

	return ranks, nil
}

func (r *RankEvaluator) evaluateTopBuyYear(db *gorm.DB, year int) {

}

func (r *RankEvaluator) evaluateTopSellYear(db *gorm.DB, year int) {

}

func (r *RankEvaluator) evaluateTopHoldingYear(db *gorm.DB, year int) {

}
