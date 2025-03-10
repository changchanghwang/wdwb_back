package domain

import (
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/google/uuid"
)

type Stock struct {
	ddd.SoftDeletableAggregate
	Id           uuid.UUID `json:"id" gorm:"column:id;primaryKey; type:varchar(36); not null;"`
	Name         string    `json:"name" gorm:"type:varchar(255); not null;"`
	Cusip        string    `json:"cusip" gorm:"type:varchar(20); not null;"`
	Ticker       string    `json:"ticker" gorm:"type:varchar(20); not null;"`
	Exchange     string    `json:"exchange" gorm:"type:varchar(50); not null;"`
	Cik          string    `json:"cik" gorm:"type:varchar(20); not null;"`
	Sic          string    `json:"sic" gorm:"type:varchar(20); not null;"`
	Industry     string    `json:"industry" gorm:"type:varchar(255); not null;"`
	Sector       string    `json:"sector" gorm:"type:varchar(255); not null;"`
	IsDelisted   bool      `json:"isDelisted" gorm:"type:boolean; column:isDelisted; not null;"`
	Location     string    `json:"location" gorm:"type:varchar(255); not null;"`
	SecId        string    `json:"secId" gorm:"type:varchar(100); column:secId; not null;"`
	SicSector    string    `json:"sicSector" gorm:"type:varchar(255); column:sicSector; not null;"`
	SicIndustry  string    `json:"sicIndustry" gorm:"type:varchar(255); column:sicIndustry; not null;"`
	FamaIndustry string    `json:"famaIndustry" gorm:"type:varchar(255); column:famaIndustry; not null;"`
	Currency     string    `json:"currency" gorm:"type:varchar(100); not null;"`
}

func (s *Stock) TableName() string {
	return "stock"
}

func New(name, cusip, ticker, exchange, cik, sic, industry, sector, location, secId, sicSector, sicIndustry, famaIndustry, currency string, isDelisted bool) (*Stock, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, "Failed to create. Can not generate uuid.", "")
	}

	return &Stock{
		Id:           id,
		Name:         name,
		Cusip:        cusip,
		Ticker:       ticker,
		Exchange:     exchange,
		Cik:          cik,
		Sic:          sic,
		Industry:     industry,
		Sector:       sector,
		IsDelisted:   isDelisted,
		Location:     location,
		SecId:        secId,
		SicSector:    sicSector,
		SicIndustry:  sicIndustry,
		FamaIndustry: famaIndustry,
		Currency:     currency,
	}, nil
}
