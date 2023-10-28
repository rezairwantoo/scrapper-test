package usecase

import (
	"context"
	"reza/scrapper-test/model"
)

type ProductUsecase interface {
	Create(ctx context.Context, req model.CreateRequest) (model.CreateResponse, error)
	CreateScrapper(ctx context.Context, req []model.CreateRequest) (model.CreateResponse, error)
}
