package repository

import (
	"gorm.io/gorm"
	"m3gaplazma/gofiber-restapi/model/domain"
)

type UserRepository interface {
	FindAll() []domain.User
	FindById(userId int64) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Create(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository UserRepositoryImpl) FindAll() []domain.User {
	var users []domain.User
	err := repository.DB.Find(&users).Error
	if err != nil {
		users = []domain.User{}
	}
	return users
}

func (repository UserRepositoryImpl) FindById(userId int64) (domain.User, error) {
	var user domain.User
	err := repository.DB.Where("id=?", userId).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository UserRepositoryImpl) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	err := repository.DB.Where("email=?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository UserRepositoryImpl) Create(user domain.User) (domain.User, error) {
	err := repository.DB.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository UserRepositoryImpl) Update(user domain.User) (domain.User, error) {
	err := repository.DB.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
