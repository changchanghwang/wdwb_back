package presentation

import (
	"net/http"

	"github.com/changchanghwang/wdwb_back/internal/config"
	"github.com/changchanghwang/wdwb_back/internal/services/sync/application"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/gofiber/fiber/v2"
)

type SyncController struct {
	syncService *application.SyncService
}

func New(syncService *application.SyncService) *SyncController {
	return &SyncController{syncService: syncService}
}

func (c *SyncController) Route(r fiber.Router) {
	r.Post("/", c.Sync)
}

func (c *SyncController) Sync(ctx *fiber.Ctx) error {
	var body struct {
		Secret string `json:"secret"`
	}

	ctx.BodyParser(&body)

	if body.Secret != config.SyncSecret {
		return applicationError.New(http.StatusForbidden, "Invalid Secret.", "Only Admin Can run Sync.")
	}

	message, err := c.syncService.Sync()
	if err != nil {
		return applicationError.Wrap(err)
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": message})
}
