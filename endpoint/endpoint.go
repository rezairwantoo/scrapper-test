package endpoint

import (
	"net/http"
	"reza/scrapper-test/model"
	"reza/scrapper-test/usecase"

	"github.com/labstack/echo/v4"
	zlog "github.com/rs/zerolog/log"
)

func MakeCreateProductEndpoint(u usecase.ProductUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			createRequest model.CreateRequest
			err           error
			resp          model.CreateResponse
		)

		if err = c.Bind(&createRequest); err != nil {
			zlog.Info().Interface("error", err).Msg("bad request")
			return c.String(http.StatusBadRequest, "bad request")
		}

		if err = c.Validate(createRequest); err != nil {
			zlog.Info().Interface("error", err).Msg("Validate Param Create")
			return err
		}

		if resp, err = u.Create(c.Request().Context(), createRequest); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}
