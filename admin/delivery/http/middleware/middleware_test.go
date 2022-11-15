package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("token", "eyJpZCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVfaWQiOjF9")
	w := httptest.NewRecorder()

	handler := BasicAuth(func(w http.ResponseWriter, r *http.Request) {
		adminData := r.Context().Value("admin-data").(models.AdminDataToken)
		assert.Equal(t, adminData.Id, 1)
		assert.Equal(t, adminData.Username, "admin")
		assert.Equal(t, adminData.Role, 1)

		w.WriteHeader(http.StatusOK)
	})

	handler.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Result().StatusCode)
}
