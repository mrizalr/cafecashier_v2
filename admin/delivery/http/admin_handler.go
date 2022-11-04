package http

import (
	"encoding/json"
	"errors"
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
	adminData := r.Context().Value("admin-data").(models.AdminDataToken)

	if adminData.Role != "super admin" {
		w.WriteHeader(http.StatusForbidden)
		res := utils.ResponseFormatter(http.StatusForbidden, "FORBIDDEN", "errors", errors.New("RESTRICTED AREA ! \nyou can't access this page"))
		w.Write(res)

		return
	}

	request := models.CreateNewAdminRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := utils.ResponseFormatter(http.StatusBadRequest, "BAD_REQUEST", "errors", err.Error())
		w.Write(res)

		return
	}

	result, err := h.ucaseAdmin.Add(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		res := utils.ResponseFormatter(http.StatusBadGateway, "BAD_GATEWAY", "errors", err.Error())
		w.Write(res)

		return
	}

	w.WriteHeader(http.StatusCreated)
	res := utils.ResponseFormatter(http.StatusCreated, "CREATED", "data", utils.FormatToCreateNewAdminResponse(result))
	w.Write(res)
}
