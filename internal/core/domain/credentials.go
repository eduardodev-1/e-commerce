package domain

type RequestCredentials struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	GrantType string `json:"grant_type" form:"grant_type"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
