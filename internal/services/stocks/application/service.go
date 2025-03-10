package application

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	secClient "github.com/changchanghwang/wdwb_back/internal/libs/sec-client"
	"github.com/changchanghwang/wdwb_back/internal/services/stocks/infrastructure"
	"gorm.io/gorm"
)

type StockService struct {
	*ddd.ApplicationService
	stockRepository infrastructure.StockRepository
	secClient       *secClient.SecClient
}

func New(stockRepository infrastructure.StockRepository, db *gorm.DB, secClient *secClient.SecClient) *StockService {
	return &StockService{
		stockRepository:    stockRepository,
		ApplicationService: &ddd.ApplicationService{Manager: db},
		secClient:          secClient,
	}
}
