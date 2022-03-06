package service

import (
	"m3gaplazma/gofiber-restapi/model/domain"
	"m3gaplazma/gofiber-restapi/model/dto"
	"m3gaplazma/gofiber-restapi/model/mapper"
	"m3gaplazma/gofiber-restapi/repository"
)

type CategoryService interface {
	FindAllCategories() []dto.CategoryResponse
	FindCategoryById(categoryId int) (dto.CategoryResponse, error)
	CreateCategory(request dto.CreateCategoryRequest) (dto.CategoryResponse, error)
	UpdateCategory(request dto.UpdateCategoryRequest) (dto.CategoryResponse, error)
	DeleteCategory(categoryId int) error
}

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{CategoryRepository: repository}
}

func (service *CategoryServiceImpl) FindAllCategories() []dto.CategoryResponse {
	var categories []domain.Category
	categories = service.CategoryRepository.FindAll()
	return mapper.ToCategoryResponses(categories)
}

func (service *CategoryServiceImpl) FindCategoryById(categoryId int) (dto.CategoryResponse, error) {
	category, err := service.CategoryRepository.FindById(categoryId)
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	return mapper.ToCategoryResponse(category), nil
}

func (service *CategoryServiceImpl) CreateCategory(request dto.CreateCategoryRequest) (dto.CategoryResponse, error) {
	category := domain.Category{
		Name: request.Name,
	}
	result, err := service.CategoryRepository.Create(category)
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	return mapper.ToCategoryResponse(result), nil
}

func (service *CategoryServiceImpl) UpdateCategory(request dto.UpdateCategoryRequest) (dto.CategoryResponse, error) {
	category, err := service.CategoryRepository.FindById(request.Id)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	category.Name = request.Name
	result, err := service.CategoryRepository.Update(category)
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	return mapper.ToCategoryResponse(result), nil
}

func (service *CategoryServiceImpl) DeleteCategory(categoryId int) error {
	category, err := service.CategoryRepository.FindById(categoryId)
	if err != nil {
		return err
	}
	err = service.CategoryRepository.Delete(category)
	if err != nil {
		return err
	}
	return nil
}
