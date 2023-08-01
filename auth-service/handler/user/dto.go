package user

type SignUpReq struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type SignInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRes struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
