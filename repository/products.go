package repository

import (
	"context"
	"reza/scrapper-test/model"
)

func (p *PostgresRepository) Create(ctx context.Context, req model.CreateRequest) error {
	query := `INSERT INTO products (name, description, price, rating, merchant_name, image_link)
		VALUES (:name, :description, :price, :rating, :merchant_name, :image_link)`

	args := map[string]any{
		"name":          req.Name,
		"description":   req.Description,
		"price":         req.Price,
		"rating":        req.Rating,
		"merchant_name": req.MerchantName,
		"image_link":    req.ImageLink,
	}

	_, err := ExecStatementContext(ctx, p.Conn, query, args)
	return err
}
