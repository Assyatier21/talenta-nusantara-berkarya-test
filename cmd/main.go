package main

import (
	"fmt"

	db "tnb-technical-test/database"
	"tnb-technical-test/internal/delivery/api"
	"tnb-technical-test/internal/repository/language"
	"tnb-technical-test/internal/usecase"
	"tnb-technical-test/routes"
	"tnb-technical-test/utils/helper"
)

func main() {
	db := db.Init()

	repository := language.NewRepository(db)
	usecase := usecase.NewUsecase(repository)
	delivery := api.NewHandler(usecase)
	echo := routes.InitRoutes(delivery)

	helper.UseCustomValidatorHandler(echo)

	host := fmt.Sprintf("%s:%s", "127.0.0.1", "8800")
	_ = echo.Start(host)
}
