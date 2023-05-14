package api

import (
	"tnb-technical-test/internal/usecase"

	"github.com/labstack/echo/v4"
)

type DeliveryHandler interface {
	HelloDev(ctx echo.Context) (err error)
	IsPalindrome(ctx echo.Context) (err error)
	GetLanguages(ctx echo.Context) (err error)
	GetLanguageByID(ctx echo.Context) (err error)
	AddLanguage(ctx echo.Context) (err error)
	UpdateLanguage(ctx echo.Context) (err error)
	DeleteLanguageByID(ctx echo.Context) (err error)
}

type handler struct {
	usecase usecase.UsecaseHandler
}

func NewHandler(usecase usecase.UsecaseHandler) DeliveryHandler {
	return &handler{
		usecase: usecase,
	}
}
