package repository

import (
	"gorm.io/gorm"
	"m3gaplazma/gofiber-restapi/model/domain"
)

type CategoryRepository interface {
	FindAll() []domain.Category
	FindById(categoryId int) (domain.Category, error)
	Create(category domain.Category) (domain.Category, error)
	Update(category domain.Category) (domain.Category, error)
	Delete(category domain.Category) error
}

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}

func (repository CategoryRepositoryImpl) FindAll() []domain.Category {
	var categories []domain.Category
	err := repository.DB.Find(&categories).Error
	if err != nil {
		categories = []domain.Category{}
	}
	return categories
}

func (repository CategoryRepositoryImpl) FindById(categoryId int) (domain.Category, error) {
	var category domain.Category
	err := repository.DB.Where("id=?", categoryId).First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (repository CategoryRepositoryImpl) Create(category domain.Category) (domain.Category, error) {
	err := repository.DB.Create(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (repository CategoryRepositoryImpl) Update(category domain.Category) (domain.Category, error) {
	err := repository.DB.Save(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (repository CategoryRepositoryImpl) Delete(category domain.Category) error {
	err := repository.DB.Delete(&category).Error
	return err
}
