package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/services/investors/domain"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"gorm.io/gorm"
)

type InvestorRepositoryImpl struct {
	manager *gorm.DB
}

func New(manager *gorm.DB) InvestorRepository {
	return &InvestorRepositoryImpl{manager: manager}
}

func (r *InvestorRepositoryImpl) FindAll(db *gorm.DB) ([]*domain.Investor, error) {
	if db == nil {
		db = r.manager
	}

	var investors []*domain.Investor
	if err := db.Find(&investors).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to findAll. %s", err.Error()), "")
	}

	return investors, nil
}
