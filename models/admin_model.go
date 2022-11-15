package models

type CreateNewAdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role_id"`
}

type CreateNewAdminResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role_id"`
}

type AdminDataToken struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role_id"`
}

type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
