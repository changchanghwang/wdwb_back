package presentation

import (
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/libs/validate"
	"github.com/changchanghwang/wdwb_back/internal/services/ranks/application"
	"github.com/changchanghwang/wdwb_back/internal/services/ranks/commands"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/gofiber/fiber/v2"
)

type RankController struct {
	rankService *application.RankService
}

func New(rankService *application.RankService) *RankController {
	return &RankController{rankService: rankService}
}

func (c *RankController) Route(r fiber.Router) {
	r.Get("/", c.Rank)
}

// Rank godoc
// @Summary Rank
// @Description Rank
// @Tags ranks
// @Accept json
// @Produce json
// @Param command body commands.RankCommand true "Rank command"
// @Success 200 {object} response.RankResponse "Successfully retrieve investor"
// @Failure 400 {object} base.ErrorResponse{errorMessage=string} "Bad request"
// @Failure 404 {object} base.ErrorResponse{errorMessage=string} "Not found"
// @Failure 500 {object} base.ErrorResponse{errorMessage=string} "Internal server error"
// @Router /ranks [get]
func (c *RankController) Rank(ctx *fiber.Ctx) error {
	language := ctx.Locals("language").(string)

	command := &commands.RankCommand{}

	if err := ctx.QueryParser(command); err != nil {
		return applicationError.Wrap(err)
	}

	if err := validate.ValidateStruct(command); err != nil {
		return applicationError.Wrap(err)
	}

	rankResponse, err := c.rankService.Rank(language, *command)
	if err != nil {
		return applicationError.Wrap(err)
	}

	return ctx.Status(http.StatusOK).JSON(rankResponse)
}
