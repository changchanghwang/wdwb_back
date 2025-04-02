package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	"github.com/changchanghwang/wdwb_back/internal/services/filings/domain"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"gorm.io/gorm"
)

type FilingRepositoryImpl struct {
	ddd.Repository[*domain.Filing]
}

func New(manager *gorm.DB) FilingRepository {
	return &FilingRepositoryImpl{ddd.Repository[*domain.Filing]{Manager: manager}}
}

func (r *FilingRepositoryImpl) FindByAccessionNumbers(manager *gorm.DB, accessionNumbers []string) ([]*domain.Filing, error) {
	if manager == nil {
		manager = r.Manager
	}

	var filings []*domain.Filing
	if err := manager.Where("accessionNumber IN ?", accessionNumbers).Find(&filings).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*domain.Filing{}, nil
		}

		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to findByAccessionNumbers. %s", err.Error()), "")
	}

	return filings, nil
}
