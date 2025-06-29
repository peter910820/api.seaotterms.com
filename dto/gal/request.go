package gal

type LoginRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type ArticleCreateRequest struct {
}
