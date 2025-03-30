package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	DefaultLanguage = "en"
)

func LanguageMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		acceptLanguage := c.Get("Accept-Language")
		if acceptLanguage == "" {
			c.Locals("language", DefaultLanguage)
			return c.Next()
		}

		// Accept-Language 헤더 파싱
		// 예: "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7"
		languages := strings.Split(acceptLanguage, ",")
		for _, lang := range languages {
			// q=0.9 같은 품질 값 제거
			lang = strings.Split(lang, ";")[0]
			// ko-KR -> ko
			lang = strings.Split(lang, "-")[0]

			// 지원하는 언어인지 확인
			if isSupportedLanguage(lang) {
				c.Locals("language", lang)
				return c.Next()
			}
		}

		// 지원하는 언어가 없으면 기본값 사용
		c.Locals("language", DefaultLanguage)
		return c.Next()
	}
}

func isSupportedLanguage(lang string) bool {
	supportedLanguages := []string{"ko", "en"}
	for _, supported := range supportedLanguages {
		if lang == supported {
			return true
		}
	}
	return false
}
