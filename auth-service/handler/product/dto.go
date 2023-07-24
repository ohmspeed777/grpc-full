package product

import "time"

type Product struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
	Price     float64   `json:"price"`
}
