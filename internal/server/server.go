package server

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/translate"
	"github.com/changchanghwang/wdwb_back/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Server struct {
	app *fiber.App
}

func New(
	route *Route,
	translator *translate.Translator,
) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.NewErrorHandler(translator).Middleware,
	})

	// request logger
	app.Use(requestid.New(), logger.New(logger.Config{
		Format:     "${time} | ${pid} | ${locals:requestid} | ${status} - ${method} ${path}\u200b\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "UTC",
	}))

	app.Use(middlewares.LanguageMiddleware())

	route.Route(app)

	return &Server{
		app: app,
	}
}

func (s *Server) Run(addr string) error {
	return s.app.Listen(addr)
}
