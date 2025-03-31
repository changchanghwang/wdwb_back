package infrastructure

import (
	"fmt"

	"github.com/changchanghwang/wdwb_back/internal/libs/db"
	"gorm.io/gorm"
)

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

func applyOptions(options *db.FindOptions, orderOptions *db.OrderOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if options != nil {
			if options.Offset != 0 {
				db = db.Offset(options.Offset)
			}
			if options.Limit != 0 {
				db = db.Limit(options.Limit)
			}
			if options.GroupBy != "" {
				db = db.Group(options.GroupBy)
			}
		}

		if orderOptions != nil {
			db = db.Order(fmt.Sprintf("%s %s", orderOptions.OrderBy, orderOptions.Direction))
		}

		return db
	}
}
