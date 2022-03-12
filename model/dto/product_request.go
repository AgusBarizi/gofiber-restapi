package dto

type CreateProductRequest struct {
	Name       string `json:"name"`
	Sku        string `json:"sku"`
	Price      int64  `json:"price"`
	Stock      int64  `json:"stock"`
	CategoryId int    `json:"category_id"`
}

type UpdateProductRequest struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Sku        string `json:"sku"`
	Price      int64  `json:"price"`
	Stock      int64  `json:"stock"`
	CategoryId int    `json:"category_id"`
}
