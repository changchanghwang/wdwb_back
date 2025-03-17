package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/services/filings/domain"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"gorm.io/gorm"
)

type FilingRepositoryImpl struct {
	manager *gorm.DB
}

func New(manager *gorm.DB) FilingRepository {
	return &FilingRepositoryImpl{manager: manager}
}

func (r *FilingRepositoryImpl) FindByAccessionNumbers(db *gorm.DB, accessionNumbers []string) ([]*domain.Filing, error) {
	if db == nil {
		db = r.manager
	}

	var filings []*domain.Filing
	if err := db.Where("accessionNumber IN ?", accessionNumbers).Find(&filings).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*domain.Filing{}, nil
		}

		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to findByAccessionNumbers. %s", err.Error()), "")
	}

	return filings, nil
}

func (r *FilingRepositoryImpl) Save(db *gorm.DB, filings []*domain.Filing) error {
	if db == nil {
		db = r.manager
	}

	if err := db.Save(filings).Error; err != nil {
		return applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to save. %s", err.Error()), "")
	}

	return nil
}
