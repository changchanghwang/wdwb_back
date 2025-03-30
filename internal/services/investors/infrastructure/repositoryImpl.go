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
	ddd.Repository[domain.Investor]
}

func New(manager *gorm.DB) InvestorRepository {
	return &InvestorRepositoryImpl{ddd.Repository[domain.Investor]{Manager: manager}}
}

func (r *InvestorRepositoryImpl) FindAll(db *gorm.DB) ([]*domain.Investor, error) {
	if db == nil {
		db = r.Manager
	}

	var investors []*domain.Investor
	if err := db.Find(&investors).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to findAll. %s", err.Error()), "")
	}

	return investors, nil
}

func (r *InvestorRepositoryImpl) Count(db *gorm.DB) (int, error) {
	if db == nil {
		db = r.Manager
	}

	var count int64
	if err := db.Model(&domain.Investor{}).Count(&count).Error; err != nil {
		return 0, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to count. %s", err.Error()), "")
	}

	return int(count), nil
}

func (r *InvestorRepositoryImpl) FindOneOrFail(db *gorm.DB, id uuid.UUID) (*domain.Investor, error) {
	if db == nil {
		db = r.Manager
	}

	var investor domain.Investor
	if err := db.Where("id = ?", id).First(&investor).Error; err != nil {
		return nil, applicationError.New(http.StatusNotFound, fmt.Sprintf("Failed to findOneOrFail. %s", err.Error()), "")
	}

	return &investor, nil
}
