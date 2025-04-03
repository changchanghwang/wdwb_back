package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/db"
	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	"github.com/changchanghwang/wdwb_back/internal/services/stocks/domain"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockRepositoryImpl struct {
	ddd.Repository[*domain.Stock]
}

func New(manager *gorm.DB) StockRepository {
	return &StockRepositoryImpl{ddd.Repository[*domain.Stock]{Manager: manager}}
}

func (r *StockRepositoryImpl) FindOneOrFail(manager *gorm.DB, id uuid.UUID) (*domain.Stock, error) {
	if manager == nil {
		manager = r.Manager
	}

	var stock *domain.Stock
	if err := manager.Where("id = ?", id).First(&stock).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to findById. %s", err.Error()), "")
	}
	if stock == nil {
		return nil, applicationError.New(http.StatusNotFound, fmt.Sprintf("Stock(%s) not found", id.String()), "")
	}

	return stock, nil
}

func (r *StockRepositoryImpl) Save(manager *gorm.DB, stocks []*domain.Stock) error {
	if manager == nil {
		manager = r.Manager
	}

	if err := manager.Save(stocks).Error; err != nil {
		return applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to save. %s", err.Error()), "")
	}
	return nil
}

func (r *StockRepositoryImpl) Find(manager *gorm.DB, conditions *StockQueryConditions, options *db.FindOptions, orderOptions *db.OrderOptions) ([]*domain.Stock, error) {
	if manager == nil {
		manager = r.Manager
	}

	manager = manager.Scopes(applyConditions(conditions), db.ApplyOptions(options, orderOptions))

	var stocks []*domain.Stock
	if err := manager.Find(&stocks).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to find. %s", err.Error()), "")
	}
	return stocks, nil
}
