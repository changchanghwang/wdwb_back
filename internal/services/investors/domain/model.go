package domain

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	"github.com/google/uuid"
)

type Investor struct {
	ddd.Aggregate
	Id           uuid.UUID `json:"id" gorm:"column:id; type:varchar(36); primaryKey"`
	Name         string    `json:"name" gorm:"type:varchar(255); not null;"`
	CompanyName  string    `json:"companyName" gorm:"type:varchar(255); column:companyName; not null;"`
	Cik          string    `json:"cik" gorm:"type:varchar(20);  not null;"`
	HoldingValue int       `json:"holdingValue" gorm:"type:bigint; column:holdingValue; not null;"`
}

func (u *Investor) TableName() string {
	return "investor"
}
