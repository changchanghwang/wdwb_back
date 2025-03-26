package presentation

import (
	"fmt"
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/base"
	"github.com/changchanghwang/wdwb_back/internal/libs/validate"
	"github.com/changchanghwang/wdwb_back/internal/services/investors/application"
	"github.com/changchanghwang/wdwb_back/internal/services/investors/command"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type InvestorController struct {
	investorService *application.InvestorService
}

func New(investorService *application.InvestorService) *InvestorController {
	return &InvestorController{investorService: investorService}
}

func (c *InvestorController) Route(r fiber.Router) {
	r.Get("/", c.List)
	r.Get("/:id", c.Retrieve)
}

// List godoc
// @Summary Get investors list
// @Description Get investors list
// @Tags investors
// @Accept json
// @Produce json
// @Success 200 {object} base.BaseResponse{data=response.ListResponse} "Successfully get investors list"
// @Failure 400 {object} base.ErrorResponse{data=string} "Bad request"
// @Failure 500 {object} base.ErrorResponse{data=string} "Internal server error"
// @Router /investors [get]
func (c *InvestorController) List(ctx *fiber.Ctx) error {
	res, err := c.investorService.List()
	if err != nil {
		return applicationError.Wrap(err)
	}

	return ctx.Status(http.StatusOK).JSON(base.NewResponse(res))
}

// Retrieve godoc
// @Summary Retrieve investor
// @Description Retrieve investor
// @Tags investors
// @Accept json
// @Produce json
// @Param id path string true "Investor ID"
// @Success 200 {object} base.BaseResponse{data=response.RetrieveResponse} "Successfully retrieve investor"
// @Failure 400 {object} base.ErrorResponse{data=string} "Bad request"
// @Failure 404 {object} base.ErrorResponse{data=string} "Not found"
// @Failure 500 {object} base.ErrorResponse{data=string} "Internal server error"
// @Router /investors/{id} [get]
func (c *InvestorController) Retrieve(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return applicationError.New(http.StatusBadRequest, fmt.Sprintf("uuid parse error: %s", err.Error()), "Invalid investor Id")
	}

	retrieveCommand := &command.RetrieveCommand{
		Id: id,
	}

	if err := validate.ValidateStruct(retrieveCommand); err != nil {
		return err
	}

	res, err := c.investorService.Retrieve(retrieveCommand)
	if err != nil {
		return applicationError.Wrap(err)
	}

	return ctx.Status(http.StatusOK).JSON(base.NewResponse(res))
}
