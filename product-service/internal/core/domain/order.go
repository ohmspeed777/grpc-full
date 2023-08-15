package domain

import (
	"time"
)

type Order struct {
	ID        string
	Items     []*OrderItem
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
}

type OrderItem struct {
	Quantity  uint
	ProductID string
	Product   *Product
}
