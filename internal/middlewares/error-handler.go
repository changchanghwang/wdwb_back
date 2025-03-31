package middlewares

import (
	"fmt"

	"github.com/changchanghwang/wdwb_back/internal/libs/base"
	"github.com/changchanghwang/wdwb_back/internal/libs/translate"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/gofiber/fiber/v2"
)

type ErrorHandler struct {
	translator *translate.Translator
}

func NewErrorHandler(translator *translate.Translator) *ErrorHandler {
	return &ErrorHandler{
		translator: translator,
	}
}

func (h *ErrorHandler) Middleware(ctx *fiber.Ctx, err error) error {
	locale := ctx.Locals("language").(string)
	if e, ok := err.(*fiber.Error); ok {
		translatedMessage := h.translator.Translate("error-message", locale, "ERRC500", false)
		return ctx.Status(e.Code).JSON(base.ErrorResponse{ErrorMessage: translatedMessage})
	}

	e := applicationError.UnWrap(err)
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	//TODO: log error with something (e.g. Sentry, ELK, File, etc.)
	fmt.Println(e.Stack)

	translatedMessage := h.translator.Translate("error-message", locale, e.ErrorCode, false)
	return ctx.Status(e.StatusCode).JSON(base.ErrorResponse{ErrorMessage: translatedMessage})
}
