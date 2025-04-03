package infrastructure

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockQueryConditions struct {
	Cusips []string
	Ids    []uuid.UUID
}

func applyConditions(conditions *StockQueryConditions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if conditions == nil {
			return db
		}

		if len(conditions.Cusips) > 0 {
			db = db.Where("cusip IN ?", conditions.Cusips)
		}

		if len(conditions.Ids) > 0 {
			db = db.Where("id IN ?", conditions.Ids)
		}

		return db
	}
}
