package presentation

import (
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/validate"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/application"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings/commands"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/gofiber/fiber/v2"
)

type HoldingController struct {
	holdingService *application.HoldingService
}

func New(holdingService *application.HoldingService) *HoldingController {
	return &HoldingController{
		holdingService: holdingService,
	}
}

func (c *HoldingController) Route(r fiber.Router) {
	r.Get("/", c.List)
}

// List godoc
// @Summary Get holdings list
// @Description Get holdings list
// @Tags holdings
// @Accept json
// @Produce json
// @Param query query commands.ListCommand true "List command"
// @Success 200 {object} response.HoldingListResponse "Successfully get holdings list"
// @Failure 400 {object} base.ErrorResponse{errorMessage=string} "Bad request"
// @Failure 500 {object} base.ErrorResponse{errorMessage=string} "Internal server error"
// @Router /holdings [get]
func (c *HoldingController) List(ctx *fiber.Ctx) error {
	language := ctx.Locals("language").(string)
	command := &commands.ListCommand{}
	if err := ctx.QueryParser(command); err != nil {
		return applicationError.Wrap(err)
	}

	if err := validate.ValidateStruct(command); err != nil {
		return applicationError.Wrap(err)
	}

	res, err := c.holdingService.List(language, command)
	if err != nil {
		return applicationError.Wrap(err)
	}

	return ctx.Status(http.StatusOK).JSON(res)
}
