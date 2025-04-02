package infrastructure

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/db"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/domain"
	"gorm.io/gorm"
)

type HoldingRepository interface {
	Save(manager *gorm.DB, holdings []*domain.Holding) error
	Find(manager *gorm.DB, conditions *HoldingQueryConditions, options *db.FindOptions, orderOptions *db.OrderOptions) ([]*domain.Holding, error)
	Count(manager *gorm.DB, conditions *HoldingQueryConditions, options *db.FindOptions) (int, error)
}
