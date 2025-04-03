package infrastructure

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/db"
	"github.com/changchanghwang/wdwb_back/internal/services/filings/domain"
	"gorm.io/gorm"
)

type FilingRepository interface {
	Find(manager *gorm.DB, conditions *FilingQueryConditions, options *db.FindOptions, orderOptions *db.OrderOptions) ([]*domain.Filing, error)
	Save(manager *gorm.DB, filings []*domain.Filing) error
}
