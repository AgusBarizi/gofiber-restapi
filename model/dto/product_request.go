package dto

type CreateProductRequest struct {
	Name       string `json:"name" form:"name"`
	Sku        string `json:"sku" form:"sku"`
	Price      int64  `json:"price" form:"price"`
	Stock      int64  `json:"stock" form:"stock"`
	CategoryId int    `json:"category_id" form:"category_id"`
	Image      string `json:"image" form:"image"`
}

type UpdateProductRequest struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Sku        string `json:"sku"`
	Price      int64  `json:"price"`
	Stock      int64  `json:"stock"`
	CategoryId int    `json:"category_id"`
}
