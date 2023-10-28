package main

import (
	"context"
	"errors"
	"log"
	"reza/scrapper-test/config"
	"reza/scrapper-test/endpoint"
	"reza/scrapper-test/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	ctx := context.Background()
	config.LoadConfigFile(ctx)
	settings, err := config.NewSettings(ctx)
	if err != nil {
		errWrap := errors.New("initialize settings, err: " + err.Error())
		log.Fatalln("initialize settings error", errWrap)
	}

	settings.Load(settings.SetPostgresRepo(settings))
	usecaseProducts := usecase.NewProductUsecase(settings.PostgresSQLProvider)
	e.POST("/products", endpoint.MakeCreateProductEndpoint(usecaseProducts))
	e.Logger.Fatal(e.Start(":1323"))
}
