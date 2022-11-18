package models

type Admin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type CreateNewAdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role_id"`
}

type CreateNewAdminResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type AdminDataToken struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
