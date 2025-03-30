package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	"github.com/changchanghwang/wdwb_back/internal/services/stocks/domain"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockRepositoryImpl struct {
	ddd.Repository[domain.Stock]
}

func New(manager *gorm.DB) StockRepository {
	return &StockRepositoryImpl{ddd.Repository[domain.Stock]{Manager: manager}}
}

func (r *StockRepositoryImpl) FindOneOrFail(db *gorm.DB, id uuid.UUID) (*domain.Stock, error) {
	if db == nil {
		db = r.Manager
	}

	var stock *domain.Stock
	if err := db.Where("id = ?", id).First(&stock).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to findById. %s", err.Error()), "")
	}
	if stock == nil {
		return nil, applicationError.New(http.StatusNotFound, fmt.Sprintf("Stock(%s) not found", id.String()), "")
	}

	return stock, nil
}

func (r *StockRepositoryImpl) Save(db *gorm.DB, stocks []*domain.Stock) error {
	if db == nil {
		db = r.Manager
	}

	if err := db.Save(stocks).Error; err != nil {
		return applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to save. %s", err.Error()), "")
	}
	return nil
}

func (r *StockRepositoryImpl) FindAll(db *gorm.DB) ([]*domain.Stock, error) {
	if db == nil {
		db = r.Manager
	}

	var stocks []*domain.Stock
	if err := db.Find(&stocks).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to findAll. %s", err.Error()), "")
	}
	return stocks, nil
}

func (r *StockRepositoryImpl) FindByCusips(db *gorm.DB, cusips []string) ([]*domain.Stock, error) {
	if db == nil {
		db = r.Manager
	}

	var stocks []*domain.Stock
	if err := db.Where("cusip IN ?", cusips).Find(&stocks).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to findByCusips. %s", err.Error()), "")
	}

	return stocks, nil
}
