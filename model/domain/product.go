package domain

import "time"

type Product struct {
	Id         int64     `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Sku        string    `json:"sku"`
	Price      int64     `json:"price"`
	Stock      int64     `json:"stock"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	CategoryId int       `json:"category_id"`
	Category   *Category `json:"category"`
	Image      string    `json:"image"`
}
