package dto

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateCategoryRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
