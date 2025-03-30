package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/db"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/domain"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"gorm.io/gorm"
)

type HoldingRepositoryImpl struct {
	manager *gorm.DB
}

func New(manager *gorm.DB) HoldingRepository {
	return &HoldingRepositoryImpl{manager: manager}
}

func (r *HoldingRepositoryImpl) Save(db *gorm.DB, holdings []*domain.Holding) error {
	if db == nil {
		db = r.manager
	}

	if err := db.Save(holdings).Error; err != nil {
		return applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to save. %s", err.Error()), "")
	}

	return nil
}

func (r *HoldingRepositoryImpl) Find(db *gorm.DB, conditions *HoldingQueryConditions, options *db.FindOptions, orderOptions *db.OrderOptions) ([]*domain.Holding, error) {
	if db == nil {
		db = r.manager
	}

	db = db.Scopes(applyConditions(conditions), applyOptions(options, orderOptions))

	var holdings []*domain.Holding
	if err := db.Find(&holdings).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to find all. %s", err.Error()), "")
	}

	return holdings, nil
}

func (r *HoldingRepositoryImpl) Count(db *gorm.DB, conditions *HoldingQueryConditions) (int, error) {
	if db == nil {
		db = r.manager
	}

	db = db.Scopes(applyConditions(conditions))

	var count int64
	if err := db.Model(&domain.Holding{}).Count(&count).Error; err != nil {
		return 0, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to count. %s", err.Error()), "")
	}

	return int(count), nil
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

func applyOptions(options *db.FindOptions, orderOptions *db.OrderOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if options != nil {
			db = db.Offset(options.Offset).Limit(options.Limit)
		}

		if orderOptions != nil {
			db = db.Order(fmt.Sprintf("%s %s", orderOptions.OrderBy, orderOptions.Direction))
		}

		return db
	}
}
