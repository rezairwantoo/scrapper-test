package model

type CreateRequest struct {
	Name         string `json:"name" validatestring:"required"`
	Description  string `json:"description" validate:"required"`
	Price        string `json:"price" validate:"required"`
	Rating       string `json:"rating" validate:"required"`
	MerchantName string `json:"merchant_name" validate:"required"`
	ImageLink    string `json:"image_link"`
}
