package domain

import (
	"fmt"
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/google/uuid"
)

type Holding struct {
	ddd.Aggregate
	Id         uuid.UUID `json:"id" gorm:"column:id; type:varchar(32); primaryKey"`
	Name       string    `json:"name" gorm:"type:varchar(255); not null;"`
	Cik        string    `json:"cik" gorm:"type:varchar(20); not null;"`
	Cusip      string    `json:"cusip" gorm:"type:varchar(20);  not null;"`
	InvestorId uuid.UUID `json:"investorId" gorm:"type:varchar(36); column:investorId; not null;"`
	StockId    uuid.UUID `json:"stockId" gorm:"type:varchar(36); column:stockId; not null;"`
	Value      int       `json:"value" gorm:"type:int; not null;"`
	Shares     int       `json:"shares" gorm:"type:int; not null;"`
	Year       int       `json:"year" gorm:"type:int; not null;"`
	Quarter    int       `json:"quarter" gorm:"type:int; not null;"`
}

func (h *Holding) TableName() string {
	return "holding"
}

func New(name, cik, cusip string, investorId, stockId uuid.UUID, value, shares, year, quarter int) (*Holding, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("failed to generate uuid: %v", err), "")
	}

	return &Holding{
		Id:         id,
		Name:       name,
		Cik:        cik,
		Cusip:      cusip,
		InvestorId: investorId,
		StockId:    stockId,
		Value:      value,
		Shares:     shares,
		Year:       year,
		Quarter:    quarter,
	}, nil
}
