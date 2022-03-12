package controller

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"m3gaplazma/gofiber-restapi/exception"
	"m3gaplazma/gofiber-restapi/helper"
	"m3gaplazma/gofiber-restapi/model/dto"
	"m3gaplazma/gofiber-restapi/service"
	"m3gaplazma/gofiber-restapi/validation"
	"strconv"
)

type ProductController interface {
	FindAllProducts(ctx *fiber.Ctx) error
	FindProductDetail(ctx *fiber.Ctx) error
	CreateProduct(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
}

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(service *service.ProductServiceImpl) *ProductControllerImpl {
	return &ProductControllerImpl{ProductService: service}
}

func (controller *ProductControllerImpl) FindAllProducts(ctx *fiber.Ctx) error {
	log.Println("find all products")
	result := controller.ProductService.FindAllProducts()
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: result})
}

func (controller *ProductControllerImpl) FindProductDetail(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("productId"))
	exception.PanicIfError(err)

	var result dto.ProductDetailResponse
	result, err = controller.ProductService.FindProductDetail(id)
	exception.PanicIfError(err)

	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: result})
}

func (controller *ProductControllerImpl) CreateProduct(ctx *fiber.Ctx) error {
	request := new(dto.CreateProductRequest)
	err := ctx.BodyParser(request)
	exception.PanicIfError(err)
	validation.CreateProductValidation(*request)

	var result dto.ProductResponse
	result, err = controller.ProductService.CreateProduct(*request)
	exception.PanicIfError(err)
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: result})
}

func (controller *ProductControllerImpl) UpdateProduct(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("productId"))
	exception.PanicIfError(err)

	request := new(dto.UpdateProductRequest)
	err = ctx.BodyParser(request)
	exception.PanicIfError(err)

	request.Id = int64(id)
	validation.UpdateProductValidation(*request)

	var result dto.ProductResponse
	result, err = controller.ProductService.UpdateProduct(*request)
	exception.PanicIfError(err)
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: result})
}

func (controller *ProductControllerImpl) DeleteProduct(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("productId"))
	exception.PanicIfError(err)

	err = controller.ProductService.DeleteProduct(id)
	exception.PanicIfError(err)
	return helper.SuccessResponse(ctx, dto.ApiResponse{Message: "product was deleted"})
}
