package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/db"
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

func (r *FilingRepositoryImpl) Find(manager *gorm.DB, conditions *FilingQueryConditions, options *db.FindOptions, orderOptions *db.OrderOptions) ([]*domain.Filing, error) {
	if manager == nil {
		manager = r.Manager
	}

	manager = manager.Scopes(applyConditions(conditions), db.ApplyOptions(options, orderOptions))

	var filings []*domain.Filing
	if err := manager.Find(&filings).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to find all. %s", err.Error()), "")
	}

	return filings, nil
}
