package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	"github.com/changchanghwang/wdwb_back/internal/services/investors/domain"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvestorRepositoryImpl struct {
	ddd.Repository[*domain.Investor]
}

func New(manager *gorm.DB) InvestorRepository {
	return &InvestorRepositoryImpl{ddd.Repository[*domain.Investor]{Manager: manager}}
}

func (r *InvestorRepositoryImpl) FindAll(manager *gorm.DB) ([]*domain.Investor, error) {
	if manager == nil {
		manager = r.Manager
	}

	var investors []*domain.Investor
	if err := manager.Find(&investors).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to findAll. %s", err.Error()), "")
	}

	return investors, nil
}

func (r *InvestorRepositoryImpl) Count(manager *gorm.DB) (int, error) {
	if manager == nil {
		manager = r.Manager
	}

	var count int64
	if err := manager.Model(&domain.Investor{}).Count(&count).Error; err != nil {
		return 0, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to count. %s", err.Error()), "")
	}

	return int(count), nil
}

func (r *InvestorRepositoryImpl) FindOneOrFail(manager *gorm.DB, id uuid.UUID) (*domain.Investor, error) {
	if manager == nil {
		manager = r.Manager
	}

	var investor domain.Investor
	if err := manager.Where("id = ?", id).First(&investor).Error; err != nil {
		return nil, applicationError.New(http.StatusNotFound, fmt.Sprintf("Failed to findOneOrFail. %s", err.Error()), "")
	}

	return &investor, nil
}
