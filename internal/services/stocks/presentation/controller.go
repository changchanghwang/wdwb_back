package presentation

import (
	"github.com/changchanghwang/wdwb_back/internal/services/stocks/application"
	"github.com/gofiber/fiber/v2"
)

type StockController struct {
	stockService *application.StockService
}

func New(stockService *application.StockService) *StockController {
	return &StockController{
		stockService: stockService,
	}
}

func (c *StockController) Route(r fiber.Router) {
}
