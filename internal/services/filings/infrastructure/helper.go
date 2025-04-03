package infrastructure

import (
	"gorm.io/gorm"
)

type FilingQueryConditions struct {
	AccessionNumbers []string
}

func applyConditions(conditions *FilingQueryConditions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if conditions == nil {
			return db
		}

		if len(conditions.AccessionNumbers) > 0 {
			db = db.Where("accessionNumber IN ?", conditions.AccessionNumbers)
		}

		return db
	}
}
