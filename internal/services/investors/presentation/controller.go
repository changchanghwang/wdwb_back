package presentation

import (
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/services/investors/application"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/gofiber/fiber/v2"
)

type InvestorController struct {
	investorService *application.InvestorService
}

func New(investorService *application.InvestorService) *InvestorController {
	return &InvestorController{investorService: investorService}
}

func (c *InvestorController) Route(r fiber.Router) {
	r.Get("/", c.List)
}

func (c *InvestorController) List(ctx *fiber.Ctx) error {
	res, err := c.investorService.List()
	if err != nil {
		return applicationError.Wrap(err)
	}

	return ctx.Status(http.StatusOK).JSON(res)
}
