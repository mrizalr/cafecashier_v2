package http

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/mrizalr/cafecashierpt2/domain/mocks"
	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/stretchr/testify/assert"
)

func TestAddNewAdmin(t *testing.T) {
	t.Run("Test Success Add New Admin", func(t *testing.T) {
		body := strings.NewReader(`
			{
				"username" : "admin2",
				"password" : "admin123",
				"role" : "super super admin"
			}	
		`)

		req := httptest.NewRequest(http.MethodPost, "/admin", body)
		ctx := context.WithValue(context.Background(), "admin-data", models.AdminDataToken{
			Id:       1,
			Username: "owner",
			Role:     "super admin",
		})

		w := httptest.NewRecorder()

		mockUcase := new(mocks.AdminUcase)
		mockUcase.On("Add", ctx, &models.CreateNewAdminRequest{
			Username: "admin2",
			Password: "admin123",
			Role:     "super super admin",
		}).Return(domain.Admin{
			ID:       2,
			Username: "admin2",
			Role:     "super super admin",
		}, nil)

		adminUcase := AdminHandler{mockUcase}

		adminUcase.AddNewAdmin(w, req.WithContext(ctx))

		res := w.Result()
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		expect := `{"code":201,"data":{"id":2,"username":"admin2","role":"super super admin"},"status":"CREATED"}`

		mockUcase.AssertExpectations(t)
		assert.Equal(t, 201, res.StatusCode)
		assert.Equal(t, expect, string(data))
	})

	t.Run("Test Non Authorized", func(t *testing.T) {
		body := strings.NewReader(`
			{
				"username" : "admin2",
				"password" : "admin123",
				"role" : "super super admin"
			}	
		`)

		req := httptest.NewRequest(http.MethodPost, "/admin", body)
		w := httptest.NewRecorder()

		mockUcase := new(mocks.AdminUcase)
		adminUcase := AdminHandler{mockUcase}

		adminUcase.AddNewAdmin(w, req)

		res := w.Result()
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		expect := `{"code":403,"errors":"can't access this endpoint","status":"FORBIDDEN"}`

		mockUcase.AssertExpectations(t)
		assert.Equal(t, 403, res.StatusCode)
		assert.Equal(t, expect, string(data))
	})
}
