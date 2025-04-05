package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/db"
	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	"github.com/changchanghwang/wdwb_back/internal/services/ranks/domain"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RankRepositoryImpl struct {
	ddd.Repository[*domain.Rank]
}

func New(manager *gorm.DB) RankRepository {
	return &RankRepositoryImpl{ddd.Repository[*domain.Rank]{Manager: manager}}
}

func (r *RankRepositoryImpl) FindOneOrFail(manager *gorm.DB, id uuid.UUID) (*domain.Rank, error) {
	if manager == nil {
		manager = r.Manager
	}

	var rank *domain.Rank
	if err := manager.Where("id = ?", id).First(&rank).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to findById. %s", err.Error()), "")
	}
	if rank == nil {
		return nil, applicationError.New(http.StatusNotFound, fmt.Sprintf("Rank(%s) not found", id.String()), "")
	}

	return rank, nil
}

func (r *RankRepositoryImpl) Save(manager *gorm.DB, ranks []*domain.Rank) error {
	if manager == nil {
		manager = r.Manager
	}

	if err := manager.Save(ranks).Error; err != nil {
		return applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to save. %s", err.Error()), "")
	}
	return nil
}

func (r *RankRepositoryImpl) Find(manager *gorm.DB, conditions *RankQueryConditions, options *db.FindOptions, orderOptions *db.OrderOptions) ([]*domain.Rank, error) {
	if manager == nil {
		manager = r.Manager
	}

	manager = manager.Scopes(applyConditions(conditions), db.ApplyOptions(options, orderOptions))

	var ranks []*domain.Rank
	if err := manager.Find(&ranks).Error; err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to find. %s", err.Error()), "")
	}
	return ranks, nil
}
