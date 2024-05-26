package dto

type RequestAuth struct {
	Username      string  `json:"username"`
	Password string `json:"password"`
}

type RequestCreateUser struct {
	Email      string  `json:"email"`
	Password string `json:"password"`
	Username     string `json:"username"`
}
