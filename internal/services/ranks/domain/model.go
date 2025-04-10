package domain

import (
	"net/http"
	"time"

	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/google/uuid"
)

type RankType int

const (
	TopBuyQuarter RankType = iota
	TopSellQuarter
	TopHoldingQuarter
	TopBuyYear
	TopSellYear
	TopHoldingYear
)

var RankTypes = []RankType{
	TopBuyQuarter,
	TopSellQuarter,
	TopHoldingQuarter,
	TopBuyYear,
	TopSellYear,
	TopHoldingYear,
}

type Rank struct {
	Id            uuid.UUID `json:"id"`
	Type          RankType  `json:"type"`
	Year          int       `json:"year"`
	Quarter       int       `json:"quarter"`
	Rank          int       `json:"rank"`
	Value         int       `json:"value"`
	Tickers       []string  `json:"tickers"`
	Name          string    `json:"name"`
	InvestorId    string    `json:"investorId"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
}

func New(rankType RankType, year, quarter, rank, value int, tickers []string, name, investorId string) (*Rank, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, "Failed to create. Can not generate uuid.", "")
	}

	return &Rank{
		Id:            id,
		Type:          rankType,
		Year:          year,
		Quarter:       quarter,
		Rank:          rank,
		Value:         value,
		Tickers:       tickers,
		Name:          name,
		InvestorId:    investorId,
		LastUpdatedAt: time.Now(),
	}, nil
}
