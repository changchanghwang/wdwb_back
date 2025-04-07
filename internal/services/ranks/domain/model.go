package domain

import (
	"net/http"
	"time"

	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
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
	ddd.Aggregate
	Id            uuid.UUID `json:"id" gorm:"column:id; type:varchar(36); primaryKey"`
	Type          RankType  `json:"type" gorm:"type:int; not null;"`
	Year          int       `json:"year" gorm:"type:int; not null;"`
	Quarter       int       `json:"quarter" gorm:"type:int; not null;"`
	Rank          int       `json:"rank" gorm:"type:int; not null;"`
	Value         int       `json:"value" gorm:"type:bigint; not null;"`
	Tickers       []string  `json:"tickers" gorm:"type:text;serializer:json"`
	Name          string    `json:"name" gorm:"type:varchar(255); not null;"`
	InvestorId    string    `json:"investorId" gorm:"type:varchar(36); column:investorId; not null;"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt" gorm:"type:datetime; column:lastUpdatedAt; not null;"`
}

func (s *Rank) TableName() string {
	return "rank"
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
