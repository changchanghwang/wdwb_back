package infrastructure

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/db"
	"github.com/changchanghwang/wdwb_back/internal/services/stocks/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockRepository interface {
	Find(manager *gorm.DB, conditions *StockQueryConditions, options *db.FindOptions, orderOptions *db.OrderOptions) ([]*domain.Stock, error)
	FindOneOrFail(manager *gorm.DB, id uuid.UUID) (*domain.Stock, error)
	Save(manager *gorm.DB, stocks []*domain.Stock) error
}
