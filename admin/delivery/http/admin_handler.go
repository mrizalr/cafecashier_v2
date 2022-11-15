package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/mrizalr/cafecashierpt2/admin/delivery/http/middleware"
	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/mrizalr/cafecashierpt2/utils"
)

type AdminHandler struct {
	ucaseAdmin domain.AdminUseCase
}

func NewAdminHandler(m *http.ServeMux, ucaseAdmin domain.AdminUseCase) {
	handler := AdminHandler{ucaseAdmin}

	m.Handle("/admin", middleware.BasicAuth(http.HandlerFunc(handler.AddNewAdmin)))
}

func (h *AdminHandler) AddNewAdmin(w http.ResponseWriter, r *http.Request) {
	utils.SetContentTypeJSON(w)
	adminData, ok := r.Context().Value("admin-data").(models.AdminDataToken)

	if adminData.Role != 1 || !ok {
		log.Printf("admins with roles other than \"super admin\" cannot access this endpoint, current admin role %v\n", adminData.Role)
		err := errors.New("can't access this endpoint")

		w.WriteHeader(http.StatusForbidden)
		w.Write(utils.ResponseFormatter(http.StatusForbidden, "FORBIDDEN", "errors", err.Error()))

		return
	}

	request := models.CreateNewAdminRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ResponseFormatter(http.StatusBadRequest, "BAD_REQUEST", "errors", err.Error()))

		return
	}

	result, err := h.ucaseAdmin.Add(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write(utils.ResponseFormatter(http.StatusBadGateway, "BAD_GATEWAY", "errors", err.Error()))

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(utils.ResponseFormatter(http.StatusCreated, "CREATED", "data", utils.FormatToCreateNewAdminResponse(result)))
}
