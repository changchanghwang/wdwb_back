package infrastructure

import (
	"fmt"
	"net/http"

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

func (r *HoldingRepositoryImpl) FindAll(db *gorm.DB) ([]*domain.Holding, error) {
	if db == nil {
		db = r.manager
	}

	var holdings []*domain.Holding
	if err := db.Find(&holdings).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to find all. %s", err.Error()), "")
	}

	return holdings, nil
}

func (r *HoldingRepositoryImpl) Count(db *gorm.DB) (int, error) {
	if db == nil {
		db = r.manager
	}

	var count int64
	if err := db.Model(&domain.Holding{}).Count(&count).Error; err != nil {
		return 0, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to count. %s", err.Error()), "")
	}

	return int(count), nil
}
