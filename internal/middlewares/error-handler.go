package middlewares

import (
	"fmt"

	"github.com/changchanghwang/wdwb_back/internal/libs/base"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).JSON(base.ErrorResponse{Data: e.Message})
	}

	e := applicationError.UnWrap(err)
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	//TODO: log error with something (e.g. Sentry, ELK, File, etc.)
	fmt.Println(e.Stack)

	return ctx.Status(e.Code).JSON(base.ErrorResponse{Data: e.ClientMessage})
}
