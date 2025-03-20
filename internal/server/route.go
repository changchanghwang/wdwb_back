package server

import (
	investors "github.com/changchanghwang/wdwb_back/internal/services/investors/presentation"
	stocks "github.com/changchanghwang/wdwb_back/internal/services/stocks/presentation"
	sync "github.com/changchanghwang/wdwb_back/internal/services/sync/presentation"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	stocks   *stocks.StockController
	sync     *sync.SyncController
	investor *investors.InvestorController
}

func NewRoute(
	stocks *stocks.StockController,
	sync *sync.SyncController,
	investor *investors.InvestorController,
) *Route {
	return &Route{
		stocks:   stocks,
		sync:     sync,
		investor: investor,
	}
}

func (r *Route) Route(app *fiber.App) {
	// TODO: swagger

	// stocks
	stocksGroup := app.Group("/stocks")
	r.stocks.Route(stocksGroup)

	// sync
	syncGroup := app.Group("/sync")
	r.sync.Route(syncGroup)

	// investors
	investorsGroup := app.Group("/investors")
	r.investor.Route(investorsGroup)
}
