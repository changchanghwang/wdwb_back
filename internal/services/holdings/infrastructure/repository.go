package infrastructure

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/db"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/domain"
	"gorm.io/gorm"
)

type HoldingRepository interface {
	Save(db *gorm.DB, holdings []*domain.Holding) error
	Find(db *gorm.DB, conditions *HoldingQueryConditions, options *db.FindOptions, orderOptions *db.OrderOptions) ([]*domain.Holding, error)
	Count(db *gorm.DB, conditions *HoldingQueryConditions) (int, error)
}
