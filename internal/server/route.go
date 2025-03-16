package server

import (
	stocks "github.com/changchanghwang/wdwb_back/internal/services/stocks/presentation"
	sync "github.com/changchanghwang/wdwb_back/internal/services/sync/presentation"
	"github.com/gofiber/fiber/v2"
)

func route(
	r *fiber.App,
	stockController *stocks.StockController,
	syncController *sync.SyncController,
) {

	// TODO: swagger

	// stocks
	stocksGroup := r.Group("/stocks")
	stockController.Route(stocksGroup)

	// sync
	syncGroup := r.Group("/sync")
	syncController.Route(syncGroup)
}
