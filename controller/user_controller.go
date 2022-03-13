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

type UserController interface {
	FindAllUsers(ctx *fiber.Ctx) error
	FindUserById(ctx *fiber.Ctx) error
	RegisterUser(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	ActivateUser(ctx *fiber.Ctx) error
	ChangeUserPassword(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{UserService: userService}
}

func (controller UserControllerImpl) FindAllUsers(ctx *fiber.Ctx) error {
	results := controller.UserService.FindAllUsers()
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: results})
}

func (controller UserControllerImpl) FindUserById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("userId"))
	exception.PanicIfError(err)

	var result dto.UserResponse
	result, err = controller.UserService.FindUserById(int64(id))
	exception.PanicIfError(err)

	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: result})
}

func (controller UserControllerImpl) RegisterUser(ctx *fiber.Ctx) error {
	request := new(dto.CreateUserRequest)
	err := ctx.BodyParser(request)
	exception.PanicIfError(err)
	validation.CreateUserValidation(*request)

	var result dto.UserResponse
	result, err = controller.UserService.CreateUser(*request)
	exception.PanicIfError(err)
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: result})
}

func (controller UserControllerImpl) Login(ctx *fiber.Ctx) error {
	request := new(dto.UserLoginRequest)
	err := ctx.BodyParser(request)
	exception.PanicIfError(err)
	validation.UserLoginValidation(*request)

	var result dto.UserLoginResponse
	result, err = controller.UserService.Login(*request)
	if err != nil {
		return exception.ExceptionError{
			Code:    401,
			Message: err.Error(),
		}
	}
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: result})
}

func (controller UserControllerImpl) ActivateUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("userId"))
	exception.PanicIfError(err)

	request := new(dto.ActivateUserRequest)
	request.Id = int64(id)
	err = ctx.BodyParser(request)
	exception.PanicIfError(err)
	validation.ActivateUserValidation(*request)

	var result dto.UserResponse
	result, err = controller.UserService.ActivateUser(*request)
	exception.PanicIfError(err)
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: result})
}

func (controller UserControllerImpl) ChangeUserPassword(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("userId"))
	exception.PanicIfError(err)

	request := new(dto.ChangeUserPasswordRequest)
	request.Id = int64(id)
	err = ctx.BodyParser(request)
	exception.PanicIfError(err)
	validation.ChangeUserPasswordValidation(*request)

	var result dto.UserResponse
	result, err = controller.UserService.ChangeUserPassword(*request)
	exception.PanicIfError(err)
	return helper.SuccessResponse(ctx, dto.ApiResponse{Data: result})
}
