package domain

import "time"

type Product struct {
	ID        string
	Price     float64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
