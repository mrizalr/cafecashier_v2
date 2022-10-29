package models

type CreateNewAdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
