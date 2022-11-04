package models

type CreateNewAdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type CreateNewAdminResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type AdminDataToken struct {
	Id       int
	Username string
	Role     string
}
