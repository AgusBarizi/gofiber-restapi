package controller

import (
	"github.com/gofiber/fiber/v2"
	"m3gaplazma/gofiber-restapi/exception"
	"m3gaplazma/gofiber-restapi/helper"
	"m3gaplazma/gofiber-restapi/model/dto"
	"m3gaplazma/gofiber-restapi/service"
	"m3gaplazma/gofiber-restapi/validation"
	"strconv"
)

type CategoryController interface {
	FindAllCategories(ctx *fiber.Ctx) error
	FindCategoryById(ctx *fiber.Ctx) error
	CreateCategory(ctx *fiber.Ctx) error
	UpdateCategory(ctx *fiber.Ctx) error
	DeleteCategory(ctx *fiber.Ctx) error
}

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(service *service.CategoryServiceImpl) *CategoryControllerImpl {
	return &CategoryControllerImpl{CategoryService: service}
}

func (controller *CategoryControllerImpl) FindAllCategories(ctx *fiber.Ctx) error {
	result := controller.CategoryService.FindAllCategories()
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: result})
}

func (controller *CategoryControllerImpl) FindCategoryById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("categoryId"))
	exception.PanicIfError(err)

	var result dto.CategoryResponse
	result, err = controller.CategoryService.FindCategoryById(id)
	exception.PanicIfError(err)
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: result})
}

func (controller *CategoryControllerImpl) CreateCategory(ctx *fiber.Ctx) error {
	request := new(dto.CreateCategoryRequest)
	err := ctx.BodyParser(request)
	exception.PanicIfError(err)
	validation.CreateCategoryValidation(*request)

	var result dto.CategoryResponse
	result, err = controller.CategoryService.CreateCategory(*request)
	exception.PanicIfError(err)
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: result})
}

func (controller *CategoryControllerImpl) UpdateCategory(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("categoryId"))
	exception.PanicIfError(err)

	request := new(dto.UpdateCategoryRequest)
	err = ctx.BodyParser(request)
	exception.PanicIfError(err)

	request.Id = id
	validation.UpdateCategoryValidation(*request)

	var result dto.CategoryResponse
	result, err = controller.CategoryService.UpdateCategory(*request)
	exception.PanicIfError(err)
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: result})
}

func (controller *CategoryControllerImpl) DeleteCategory(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("categoryId"))
	exception.PanicIfError(err)

	err = controller.CategoryService.DeleteCategory(id)
	exception.PanicIfError(err)
	return helper.SuccessResponse(ctx, dto.ApiResponse{Message: "category was deleted"})
}
