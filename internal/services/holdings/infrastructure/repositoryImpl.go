package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/db"
	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/domain"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"gorm.io/gorm"
)

type HoldingRepositoryImpl struct {
	ddd.Repository[*domain.Holding]
}

func New(manager *gorm.DB) HoldingRepository {
	return &HoldingRepositoryImpl{ddd.Repository[*domain.Holding]{Manager: manager}}
}

func (r *HoldingRepositoryImpl) Find(manager *gorm.DB, conditions *HoldingQueryConditions, options *db.FindOptions, orderOptions *db.OrderOptions) ([]*domain.Holding, error) {
	if manager == nil {
		manager = r.Manager
	}

	manager = manager.Scopes(applyConditions(conditions), db.ApplyOptions(options, orderOptions))

	var holdings []*domain.Holding
	if err := manager.Find(&holdings).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to find all. %s", err.Error()), "")
	}

	return holdings, nil
}

func (r *HoldingRepositoryImpl) Count(manager *gorm.DB, conditions *HoldingQueryConditions, options *db.FindOptions) (int, error) {
	if manager == nil {
		manager = r.Manager
	}

	manager = manager.Scopes(applyConditions(conditions), db.ApplyOptions(options, nil))

	var count int64
	if err := manager.Model(&domain.Holding{}).Count(&count).Error; err != nil {
		return 0, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to count. %s", err.Error()), "")
	}

	return int(count), nil
}
