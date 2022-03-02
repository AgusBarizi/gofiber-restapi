package repository

import (
	"gorm.io/gorm"
	"m3gaplazma/gofiber-restapi/exception"
	"m3gaplazma/gofiber-restapi/model/domain"
)

type ProductRepository interface {
	FindAll() []domain.Product
	FindById(productId int) (domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
	Update(product domain.Product) (domain.Product, error)
	Delete(product domain.Product) error
}

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (repository ProductRepositoryImpl) FindAll() []domain.Product {
	var products []domain.Product
	err := repository.DB.Find(&products).Error
	exception.PanicIfError(err)
	return products
}

func (repository ProductRepositoryImpl) FindById(productId int) (domain.Product, error) {
	var product domain.Product
	err := repository.DB.Where("id=?", productId).First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repository ProductRepositoryImpl) Create(product domain.Product) (domain.Product, error) {
	err := repository.DB.Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repository ProductRepositoryImpl) Update(product domain.Product) (domain.Product, error) {
	err := repository.DB.Save(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repository ProductRepositoryImpl) Delete(product domain.Product) error {
	err := repository.DB.Delete(&product).Error
	return err
}
