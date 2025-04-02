package infrastructure

import (
	"github.com/changchanghwang/wdwb_back/internal/services/filings/domain"
	"gorm.io/gorm"
)

type FilingRepository interface {
	FindByAccessionNumbers(manager *gorm.DB, accessionNumbers []string) ([]*domain.Filing, error)
	Save(manager *gorm.DB, filings []*domain.Filing) error
}
