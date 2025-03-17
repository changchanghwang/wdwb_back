package infrastructure

import (
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/domain"
	"gorm.io/gorm"
)

type HoldingRepository interface {
	Save(db *gorm.DB, filings []*domain.Holding) error
}
