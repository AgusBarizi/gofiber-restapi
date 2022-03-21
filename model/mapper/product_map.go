package mapper

import (
	"m3gaplazma/gofiber-restapi/model/domain"
	"m3gaplazma/gofiber-restapi/model/dto"
)

func ToProductResponse(product domain.Product) dto.ProductResponse {
	return dto.ProductResponse{
		Id:         product.Id,
		Name:       product.Name,
		Sku:        product.Sku,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryId: product.CategoryId,
		Image:      product.Image,
	}
}

func ToProductDetailResponse(product domain.Product) dto.ProductDetailResponse {
	productDetail := dto.ProductDetailResponse{
		Id:         product.Id,
		Name:       product.Name,
		Sku:        product.Sku,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryId: product.CategoryId,
		Image:      product.Image,
	}
	if product.Category != nil {
		productDetail.Category = ToCategoryResponse(*product.Category)
	}
	return productDetail
}

func ToProductResponses(products []domain.Product) []dto.ProductResponse {
	productResponses := []dto.ProductResponse{}
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}

func ToProductDetailResponses(products []domain.Product) []dto.ProductDetailResponse {
	productResponses := []dto.ProductDetailResponse{}
	for _, product := range products {
		productResponses = append(productResponses, ToProductDetailResponse(product))
	}
	return productResponses
}
