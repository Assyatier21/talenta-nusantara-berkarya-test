package helper

import (
	"fmt"
	"net/http"
	"strings"

	"tnb-technical-test/models"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func UseCustomValidatorHandler(c *echo.Echo) {
	c.Validator = &CustomValidator{validator: validator.New()}

	c.HTTPErrorHandler = func(err error, c echo.Context) {
		var MessageValidation []string
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required", "required_with", "required_without", "required_unless":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s is required",
						err.Field()))
				case "email":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s is not valid email",
						err.Field()))
				case "gte":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param()))
				case "lte":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param()))
				case "numeric|eq=*":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value must be numeric or *",
						err.Field()))
				case "oneof":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value not one of %s",
						err.Field(), strings.ReplaceAll(err.Param(), " ", "/")))
				}
			}
			_ = WriteResponse(c, models.StandardResponseReq{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Data:    MessageValidation,
				Error:   nil,
			})
		}
	}

}

func BindValidate(c echo.Context, req interface{}) error {
	var err = c.Bind(req)
	if err != nil {
		return err
	}
	return nil
}

func IsPalindrome(text string) bool {
	text = strings.ToLower(strings.ReplaceAll(text, " ", ""))

	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-1-i] {
			return false
		}
	}

	return true
}
