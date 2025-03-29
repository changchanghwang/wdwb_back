package presentation

import (
	"net/http"

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

func (c *HoldingController) List(ctx *fiber.Ctx) error {
	command := &commands.ListCommand{}
	if err := ctx.QueryParser(command); err != nil {
		return applicationError.Wrap(err)
	}

	res, err := c.holdingService.List(command)
	if err != nil {
		return applicationError.Wrap(err)
	}

	return ctx.Status(http.StatusOK).JSON(res)
}
