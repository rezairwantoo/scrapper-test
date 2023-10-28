package usecase

import (
	"context"
	"reza/scrapper-test/helpers"
	"reza/scrapper-test/model"
	"reza/scrapper-test/model/constant"

	zlog "github.com/rs/zerolog/log"
)

func (u *Usecase) Create(ctx context.Context, req model.CreateRequest) (model.CreateResponse, error) {
	var (
		CreateRespData model.ResponseDataCreate
		err            error
		resp           model.CreateResponse
	)

	if err = u.postgreSQL.Create(ctx, req); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Create products")
		resp.Message = constant.ErrCreate
		return resp, err
	}

	if err = helpers.WriteCsv(req); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Create csv products")
		resp.Message = constant.ErrCreate
		return resp, err
	}

	CreateRespData.Status = true
	return model.CreateResponse{
		Message: constant.SuccessCreate,
		Data:    CreateRespData,
	}, nil
}
