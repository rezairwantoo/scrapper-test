package usecase

import (
	"context"
	"reza/scrapper-test/helpers"
	"reza/scrapper-test/model"
	"reza/scrapper-test/model/constant"

	zlog "github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func (u *Usecase) CreateScrapper(ctx context.Context, req []model.CreateRequest) (model.CreateResponse, error) {
	var (
		CreateRespData model.ResponseDataCreate
		err            error
		resp           model.CreateResponse
		eg             errgroup.Group
	)

	eg.Go(func() error {
		for _, product := range req {
			if err = u.postgreSQL.Create(ctx, product); err != nil {
				zlog.Info().Interface("error", err.Error()).Msg("Failed Create products")
				resp.Message = constant.ErrCreate
				return err
			}
		}
		return nil
	})

	eg.Go(func() error {
		if err = helpers.WriteCsvBulk(req); err != nil {
			zlog.Info().Interface("error", err.Error()).Msg("Failed Create csv products")
			resp.Message = constant.ErrCreate
			return err
		}
		return nil
	})

	errGroup := eg.Wait()
	if errGroup != nil {
		return resp, errGroup
	}

	CreateRespData.Status = true
	return model.CreateResponse{
		Message: constant.SuccessCreate,
		Data:    CreateRespData,
	}, nil
}
