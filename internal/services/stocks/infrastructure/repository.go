package infrastructure

import (
	"github.com/changchanghwang/wdwb_back/internal/services/stocks/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockRepository interface {
	FindAll(manager *gorm.DB) ([]*domain.Stock, error)
	FindByCusips(manager *gorm.DB, cusips []string) ([]*domain.Stock, error)
	FindOneOrFail(manager *gorm.DB, id uuid.UUID) (*domain.Stock, error)
	Save(manager *gorm.DB, stocks []*domain.Stock) error
}
