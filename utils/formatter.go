package utils

import (
	"encoding/json"

	"github.com/mrizalr/cafecashierpt2/models"
)

func ResponseFormatter(httpstatuscode int, httpstatus string, responseType string, data interface{}) []byte {
	res := map[string]interface{}{
		"code":       httpstatuscode,
		"status":     httpstatus,
		responseType: data,
	}

	jsonResponse, _ := json.Marshal(res)
	return jsonResponse
}

func FormatToCreateNewAdminResponse(a models.Admin) *models.CreateNewAdminResponse {
	return &models.CreateNewAdminResponse{
		ID:       a.ID,
		Username: a.Username,
		Role:     a.Role,
	}
}
