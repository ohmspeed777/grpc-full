package domain

import "time"

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type SignUpReq struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type SignInReq struct {
	Email    string
	Password string
}

type SignInRes struct {
	Token        string
	RefreshToken string
}

type Order struct {
	ID        string
	Items     []*OrderItem
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	User      *User
}

type OrderItem struct {
	Quantity  uint
	ProductID string
	Product   *Product
}
