package presentation

import (
	"net/http"

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
	r.Post("/", c.Rank)
}

func (c *RankController) Rank(ctx *fiber.Ctx) error {
	var command *commands.RankCommand

	ctx.BodyParser(&command)

	err := c.rankService.Rank(*command)
	if err != nil {
		return applicationError.Wrap(err)
	}

	message := "랭킹이 성공적으로 생성되었습니다"

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": message})
}
