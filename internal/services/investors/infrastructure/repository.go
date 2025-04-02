package infrastructure

import (
	"github.com/changchanghwang/wdwb_back/internal/services/investors/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvestorRepository interface {
	FindAll(manager *gorm.DB) ([]*domain.Investor, error)
	Count(manager *gorm.DB) (int, error)
	FindOneOrFail(manager *gorm.DB, id uuid.UUID) (*domain.Investor, error)
}
