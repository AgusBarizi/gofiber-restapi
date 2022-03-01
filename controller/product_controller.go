package controller

import (
	"github.com/gofiber/fiber/v2"
	"m3gaplazma/gofiber-restapi/config"
	"m3gaplazma/gofiber-restapi/exception"
	"m3gaplazma/gofiber-restapi/helper"
	"m3gaplazma/gofiber-restapi/model/domain"
	"m3gaplazma/gofiber-restapi/model/dto"
	"m3gaplazma/gofiber-restapi/model/mapper"
	"m3gaplazma/gofiber-restapi/validation"
	"strconv"
)

func FindAllProducts(ctx *fiber.Ctx) error {
	db := config.DB
	var products []domain.Product
	err := db.Find(&products).Error
	exception.PanicIfError(err)

	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: mapper.ToProductResponses(products)})
}

func FindProductById(ctx *fiber.Ctx) error {
	id := ctx.Params("productId")
	db := config.DB

	var product domain.Product
	err := db.Where("id=?", id).First(&product).Error
	if err != nil {
		panic(exception.NotFoundError{Message: "product not found"})
	}
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: mapper.ToProductResponse(product)})
}

func CreateProduct(ctx *fiber.Ctx) error {
	createProductRequest := new(dto.CreateProductRequest)
	err := ctx.BodyParser(createProductRequest)
	exception.PanicIfError(err)
	validation.CreateProductValidation(*createProductRequest)

	product := domain.Product{
		Name:  createProductRequest.Name,
		Sku:   createProductRequest.Sku,
		Stock: createProductRequest.Stock,
	}
	db := config.DB
	err = db.Create(&product).Error
	exception.PanicIfError(err)

	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: mapper.ToProductResponse(product)})
}

func UpdateProduct(ctx *fiber.Ctx) error {
	db := config.DB
	id, err := strconv.Atoi(ctx.Params("productId"))
	exception.PanicIfError(err)

	request := new(dto.UpdateProductRequest)
	err = ctx.BodyParser(request)
	exception.PanicIfError(err)
	request.Id = int64(id)
	validation.UpdateProductValidation(*request)

	product := domain.Product{}
	err = db.Where("id=?", request.Id).First(&product).Error
	if err != nil {
		panic(exception.NotFoundError{Message: "product not found"})
	}
	product.Name = request.Name
	product.Sku = request.Sku
	product.Stock = request.Stock
	err = db.Save(&product).Error
	exception.PanicIfError(err)

	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: mapper.ToProductResponse(product)})
}

func DeleteProduct(ctx *fiber.Ctx) error {
	db := config.DB
	id, err := strconv.Atoi(ctx.Params("productId"))
	exception.PanicIfError(err)

	product := domain.Product{}
	err = db.Where("id=?", id).First(&product).Error
	if err != nil {
		panic(exception.NotFoundError{Message: "product not found"})
	}

	err = db.Delete(&product).Error
	exception.PanicIfError(err)

	return helper.SuccessResponse(ctx, dto.ApiResponse{Message: "product was deleted"})
}
