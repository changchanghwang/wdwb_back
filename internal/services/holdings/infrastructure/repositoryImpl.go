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
