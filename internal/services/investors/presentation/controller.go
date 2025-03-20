package presentation

import (
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/base"
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

// List godoc
// @Summary Get investors list
// @Description Get investors list
// @Tags investors
// @Accept json
// @Produce json
// @Success 200 {object} base.BaseResponse{data=response.ListResponse} "Successfully get investors list"
// @Failure 400 {object} error "Bad request"
// @Failure 500 {object} error "Internal server error"
// @Router /investors [get]
func (c *InvestorController) List(ctx *fiber.Ctx) error {
	res, err := c.investorService.List()
	if err != nil {
		return applicationError.Wrap(err)
	}

	return ctx.Status(http.StatusOK).JSON(base.NewResponse(res))
}
