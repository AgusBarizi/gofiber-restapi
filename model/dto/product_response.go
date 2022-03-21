package dto

type ProductResponse struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Sku        string `json:"sku"`
	Price      int64  `json:"price"`
	Stock      int64  `json:"stock"`
	CategoryId int    `json:"category_id"`
	Image      string `json:"image"`
}

type ProductDetailResponse struct {
	Id         int64       `json:"id"`
	Name       string      `json:"name"`
	Sku        string      `json:"sku"`
	Price      int64       `json:"price"`
	Stock      int64       `json:"stock"`
	CategoryId int         `json:"category_id"`
	Category   interface{} `json:"category"`
	Image      string      `json:"image"`
}
