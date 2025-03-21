package infrastructure

import (
	"github.com/changchanghwang/wdwb_back/internal/services/investors/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvestorRepository interface {
	FindAll(db *gorm.DB) ([]*domain.Investor, error)
	Count(db *gorm.DB) (int, error)
	FindOneOrFail(db *gorm.DB, id uuid.UUID) (*domain.Investor, error)
}
