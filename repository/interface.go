package repository

import (
	"context"
	"reza/scrapper-test/model"
)

type PostgresSQLRepository interface {
	Create(ctx context.Context, req model.CreateRequest) error
}
