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