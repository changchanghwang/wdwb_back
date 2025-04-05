package infrastructure

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/db"
	"github.com/changchanghwang/wdwb_back/internal/services/ranks/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RankRepository interface {
	Find(manager *gorm.DB, conditions *RankQueryConditions, options *db.FindOptions, orderOptions *db.OrderOptions) ([]*domain.Rank, error)
	FindOneOrFail(manager *gorm.DB, id uuid.UUID) (*domain.Rank, error)
	Save(manager *gorm.DB, ranks []*domain.Rank) error
}
