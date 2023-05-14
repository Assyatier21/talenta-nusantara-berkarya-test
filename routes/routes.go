package routes

import (
	"net/http"

	"tnb-technical-test/internal/delivery/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(handler api.DeliveryHandler) *echo.Echo {
	e := echo.New()
	useMiddlewares(e)

	g := e.Group("/")
	g.GET("", handler.HelloDev)
	g.GET("palindrome", handler.IsPalindrome)

	language := g.Group("language")
	language.GET("", handler.GetLanguages)
	language.GET("/:id", handler.GetLanguageByID)
	language.POST("", handler.AddLanguage)

	return e
}
func useMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodPut},
	}))
}
