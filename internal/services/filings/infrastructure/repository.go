package infrastructure

import (
	"github.com/changchanghwang/wdwb_back/internal/services/filings/domain"
	"gorm.io/gorm"
)

type FilingRepository interface {
	FindByAccessionNumbers(db *gorm.DB, accessionNumbers []string) ([]*domain.Filing, error)
	Save(db *gorm.DB, filings []*domain.Filing) error
}
