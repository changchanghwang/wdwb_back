package domain

import (
	"github.com/changchanghwang/wdtb_back/internal/libs/ddd"
	"github.com/google/uuid"
)

type Investor struct {
	ddd.Aggregate
	Id   uuid.UUID `json:"id" gorm:"column:id;primaryKey"`
	Name string    `json:"name" gorm:"type:varchar(50); not null;"`
	Cik  string    `json:"cik" gorm:"type:varchar(10); not null;"`
}

func (u *Investor) TableName() string {
	return "investor"
}
