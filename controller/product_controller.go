package controller

import (
	"github.com/gofiber/fiber/v2"
	"m3gaplazma/gofiber-restapi/exception"
	"m3gaplazma/gofiber-restapi/helper"
	"m3gaplazma/gofiber-restapi/model/domain"
	"m3gaplazma/gofiber-restapi/model/dto"
	"m3gaplazma/gofiber-restapi/model/mapper"
	"m3gaplazma/gofiber-restapi/repository"
	"m3gaplazma/gofiber-restapi/validation"
	"strconv"
)

type ProductController interface {
	FindAllProducts(ctx *fiber.Ctx) error
	FindProductById(ctx *fiber.Ctx) error
	CreateProduct(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
}

type ProductControllerImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductController(productRepository *repository.ProductRepositoryImpl) *ProductControllerImpl {
	return &ProductControllerImpl{ProductRepository: productRepository}
}

func (controller *ProductControllerImpl) FindAllProducts(ctx *fiber.Ctx) error {
	var products []domain.Product
	products = controller.ProductRepository.FindAll()

	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: mapper.ToProductResponses(products)})
}

func (controller *ProductControllerImpl) FindProductById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("productId"))
	exception.PanicIfError(err)

	var product domain.Product
	product, err = controller.ProductRepository.FindById(id)
	exception.PanicIfError(err)
	
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: mapper.ToProductResponse(product)})
}

func (controller *ProductControllerImpl) CreateProduct(ctx *fiber.Ctx) error {
	createProductRequest := new(dto.CreateProductRequest)
	err := ctx.BodyParser(createProductRequest)
	exception.PanicIfError(err)
	validation.CreateProductValidation(*createProductRequest)

	product := domain.Product{
		Name:  createProductRequest.Name,
		Sku:   createProductRequest.Sku,
		Stock: createProductRequest.Stock,
		Price: createProductRequest.Price,
	}
	product, err = controller.ProductRepository.Create(product)
	exception.PanicIfError(err)

	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: mapper.ToProductResponse(product)})
}

func (controller *ProductControllerImpl) UpdateProduct(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("productId"))
	exception.PanicIfError(err)

	request := new(dto.UpdateProductRequest)
	err = ctx.BodyParser(request)
	exception.PanicIfError(err)

	request.Id = int64(id)
	validation.UpdateProductValidation(*request)

	var product domain.Product
	product, err = controller.ProductRepository.FindById(id)
	exception.PanicIfError(exception.NotFoundError{Message: "product not found"})

	product.Name = request.Name
	product.Sku = request.Sku
	product.Stock = request.Stock
	product.Price = request.Price
	product, err = controller.ProductRepository.Update(product)
	exception.PanicIfError(err)

	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: mapper.ToProductResponse(product)})
}

func (controller *ProductControllerImpl) DeleteProduct(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("productId"))
	exception.PanicIfError(err)

	var product domain.Product
	product, err = controller.ProductRepository.FindById(id)
	exception.PanicIfError(err)

	err = controller.ProductRepository.Delete(product)
	exception.PanicIfError(err)

	return helper.SuccessResponse(ctx, dto.ApiResponse{Message: "product was deleted"})
}
