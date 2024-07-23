package models

type RequestCredentials struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	GrantType string `json:"grant_type" form:"grant_type"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
type Authority struct {
	Id        int    `json:"id"`
	Authority string `json:"authority"`
}
