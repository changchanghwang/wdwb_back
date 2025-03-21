package validate

import (
	"fmt"
	"net/http"
	"strings"

	calendarDate "github.com/changchanghwang/wdwb_back/internal/libs/calendar-date"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateCalendarDateFormat(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	_, err := calendarDate.CalendarDate(date).Parse()
	return err == nil
}

func Init() {
	validate = validator.New()
	validate.RegisterValidation("calendardate", ValidateCalendarDateFormat)
}

func ValidateStruct(dto any) error {
	if err := validate.Struct(dto); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMsgs []string
			for _, err := range validationErrors {
				var errorMsg string
				switch err.Tag() {
				case "required":
					errorMsg = fmt.Sprintf("%s는(은) 필수 값입니다.", err.Field())
				case "email":
					errorMsg = fmt.Sprintf("%s는(은) 반드시 이메일 형식이어야 합니다.", err.Field())
				case "gte":
					errorMsg = fmt.Sprintf("%s는(은) 반드시 %s보다 크거나 같아야 합니다.", err.Field(), err.Param())
				case "lte":
					errorMsg = fmt.Sprintf("%s는(은) 반드시 %s보다 작거나 같아야 합니다.", err.Field(), err.Param())
				case "oneof":
					words := strings.Fields(err.Param())
					for i, word := range words {
						words[i] = "'" + word + "'"
					}
					errorMsg = fmt.Sprintf("%s는(은) 반드시 %s 중 하나여야 합니다.", err.Field(), strings.Join(words, ", "))
				case "calendardate":
					errorMsg = fmt.Sprintf("%s는(은) 반드시 YYYY-MM-DD 형식이어야 합니다.", err.Field())
				default:
					errorMsg = fmt.Sprintf("%s는(은) 유효하지 않은 값입니다.", err.Field())
				}
				errorMsgs = append(errorMsgs, errorMsg)
			}
			message := strings.Join(errorMsgs, "\n")

			return applicationError.New(http.StatusBadRequest, message, message)
		}
	}
	return nil
}
