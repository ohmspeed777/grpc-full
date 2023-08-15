package order

import (
	"app/handler/product"
	"time"
)

type Order struct {
	ID        string       `json:"id"`
	Items     []*OrderItem `json:"items"`
	Total     float64      `json:"total"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	UserID    string       `json:"user_id"`
}

type OrderItem struct {
	Quantity  uint             `json:"quantity"`
	ProductID string           `json:"product_id"`
	Product   *product.Product `json:"product,omitempty"`
}
