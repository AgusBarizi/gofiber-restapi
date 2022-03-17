package service

import (
	"errors"
	"m3gaplazma/gofiber-restapi/helper"
	"m3gaplazma/gofiber-restapi/model/dto"
	"m3gaplazma/gofiber-restapi/model/mapper"
	"m3gaplazma/gofiber-restapi/repository"
)

type UserService interface {
	FindAllUsers() []dto.UserResponse
	CreateUser(request dto.CreateUserRequest) (dto.UserResponse, error)
	FindUserById(userId int64) (dto.UserResponse, error)
	Login(request dto.UserLoginRequest) (dto.UserLoginResponse, error)
	ActivateUser(requst dto.ActivateUserRequest) (dto.UserResponse, error)
	ChangeUserPassword(request dto.ChangeUserPasswordRequest) (dto.UserResponse, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: repository,
	}
}

func (service UserServiceImpl) FindAllUsers() []dto.UserResponse {
	users := service.UserRepository.FindAll()
	return mapper.ToUserResponses(users)
}

func (service UserServiceImpl) CreateUser(request dto.CreateUserRequest) (dto.UserResponse, error) {
	user, err := service.UserRepository.FindByEmail(request.Email)
	if err == nil {
		return dto.UserResponse{}, errors.New("user already registered")
	}

	password, err := helper.HashingPassword(request.Password)
	if err != nil {
		return dto.UserResponse{}, err
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = password
	user.IsActive = 1

	result, err := service.UserRepository.Create(user)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return mapper.ToUserResponse(result), nil
}

func (service UserServiceImpl) FindUserById(userId int64) (dto.UserResponse, error) {
	user, err := service.UserRepository.FindById(userId)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return mapper.ToUserResponse(user), nil
}

func (service UserServiceImpl) Login(request dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	userResponse := dto.UserLoginResponse{}

	user, err := service.UserRepository.FindByEmail(request.Email)
	if err != nil {
		return userResponse, err
	}

	if !helper.VerifyPassword(user.Password, request.Password) {
		return userResponse, errors.New("password not match")
	}

	if user.IsActive != 1 {
		return userResponse, errors.New("user not active")
	}

	user.ApiToken = helper.RandomString(30)
	user, err = service.UserRepository.Update(user)
	if err != nil {
		return userResponse, err
	}
	return mapper.ToUserLoginResponse(user), nil
}

func (service UserServiceImpl) ActivateUser(request dto.ActivateUserRequest) (dto.UserResponse, error) {
	userResponse := dto.UserResponse{}

	user, err := service.UserRepository.FindById(request.Id)
	if err != nil {
		return userResponse, err
	}

	user.IsActive = request.IsActive
	user, err = service.UserRepository.Update(user)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return mapper.ToUserResponse(user), nil
}

func (service UserServiceImpl) ChangeUserPassword(request dto.ChangeUserPasswordRequest) (dto.UserResponse, error) {
	userResponse := dto.UserResponse{}

	user, err := service.UserRepository.FindById(request.Id)
	if err != nil {
		return userResponse, err
	}
	if user.Password != request.OldPassword {
		return userResponse, errors.New("old password not match")
	}

	user.Password = request.NewPassword
	user, err = service.UserRepository.Update(user)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return mapper.ToUserResponse(user), nil
}
