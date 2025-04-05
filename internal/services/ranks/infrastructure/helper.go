package infrastructure

import (
	"github.com/changchanghwang/wdwb_back/internal/services/ranks/domain"
	"gorm.io/gorm"
)

type RankQueryConditions struct {
	Types  []domain.RankType
	Years  []int
	Months []int
}

func applyConditions(conditions *RankQueryConditions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if conditions == nil {
			return db
		}

		if len(conditions.Types) > 0 {
			db = db.Where("type IN ?", conditions.Types)
		}

		if len(conditions.Years) > 0 {
			db = db.Where("year IN ?", conditions.Years)
		}

		if len(conditions.Months) > 0 {
			db = db.Where("month IN ?", conditions.Months)
		}

		return db
	}
}
