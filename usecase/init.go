package usecase

import (
	"reza/scrapper-test/repository"
)

// Usecase contain all methods for usecase
type Usecase struct {
	postgreSQL repository.PostgresSQLRepository
}

func NewProductUsecase(
	repositoryPostgres repository.PostgresSQLRepository,
) ProductUsecase {
	return &Usecase{
		postgreSQL: repositoryPostgres,
	}
}
