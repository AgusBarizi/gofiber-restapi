package mapper

import (
	"m3gaplazma/gofiber-restapi/model/domain"
	"m3gaplazma/gofiber-restapi/model/dto"
)

func ToProductResponse(product domain.Product) dto.ProductResponse {
	return dto.ProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Sku:   product.Sku,
		Price: product.Price,
		Stock: product.Stock,
	}
}

func ToProductResponses(products []domain.Product) []dto.ProductResponse {
	var productResponses []dto.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
