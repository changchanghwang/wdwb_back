package infrastructure

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HoldingQueryConditions struct {
	InvestorIds []uuid.UUID
	Years       []int
	Quarters    []int
}

func applyConditions(conditions *HoldingQueryConditions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if conditions == nil {
			return db
		}

		if len(conditions.InvestorIds) > 0 {
			db = db.Where("investorId IN ?", conditions.InvestorIds)
		}

		if len(conditions.Years) > 0 {
			db = db.Where("year IN ?", conditions.Years)
		}

		if len(conditions.Quarters) > 0 {
			db = db.Where("quarter IN ?", conditions.Quarters)
		}

		return db
	}
}
