package response

import "github.com/changchanghwang/wdwb_back/internal/services/ranks/domain"

type RankResponse struct {
	TopBuyQuarter     []*domain.Rank `json:"topBuyQuarter"`     // buy ranking per quarter
	TopSellQuarter    []*domain.Rank `json:"topSellQuarter"`    // sell ranking per quarter
	TopHoldingQuarter []*domain.Rank `json:"topHoldingQuarter"` // holding ranking per quarter
	TopBuyYear        []*domain.Rank `json:"topBuyYear"`        // buy ranking per year
	TopSellYear       []*domain.Rank `json:"topSellYear"`       // sell ranking per year
	TopHoldingYear    []*domain.Rank `json:"topHoldingYear"`    // holding ranking per year
}
