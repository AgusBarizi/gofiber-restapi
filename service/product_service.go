package service

import (
	"m3gaplazma/gofiber-restapi/model/domain"
	"m3gaplazma/gofiber-restapi/model/dto"
	"m3gaplazma/gofiber-restapi/model/mapper"
	"m3gaplazma/gofiber-restapi/repository"
)

type ProductService interface {
	FindAllProducts() []dto.ProductResponse
	FindProductById(productId int) (dto.ProductResponse, error)
	CreateProduct(request dto.CreateProductRequest) (dto.ProductResponse, error)
	UpdateProduct(request dto.UpdateProductRequest) (dto.ProductResponse, error)
	DeleteProduct(productId int) error
}

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{ProductRepository: repository}
}

func (service *ProductServiceImpl) FindAllProducts() []dto.ProductResponse {
	var products []domain.Product
	products = service.ProductRepository.FindAll()
	return mapper.ToProductResponses(products)
}

func (service *ProductServiceImpl) FindProductById(productId int) (dto.ProductResponse, error) {
	product, err := service.ProductRepository.FindById(productId)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	return mapper.ToProductResponse(product), nil
}

func (service *ProductServiceImpl) CreateProduct(request dto.CreateProductRequest) (dto.ProductResponse, error) {
	product := domain.Product{
		Name:  request.Name,
		Sku:   request.Sku,
		Stock: request.Stock,
		Price: request.Price,
	}
	result, err := service.ProductRepository.Create(product)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	return mapper.ToProductResponse(result), nil
}

func (service *ProductServiceImpl) UpdateProduct(request dto.UpdateProductRequest) (dto.ProductResponse, error) {
	product, err := service.ProductRepository.FindById(int(request.Id))
	if err != nil {
		return dto.ProductResponse{}, err
	}

	product.Name = request.Name
	product.Sku = request.Sku
	product.Stock = request.Stock
	product.Price = request.Price
	result, err := service.ProductRepository.Update(product)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	return mapper.ToProductResponse(result), nil
}

func (service *ProductServiceImpl) DeleteProduct(productId int) error {
	product, err := service.ProductRepository.FindById(productId)
	if err != nil {
		return err
	}
	err = service.ProductRepository.Delete(product)
	if err != nil {
		return err
	}
	return nil
}
