package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mrizalr/cafecashierpt2/domain/mocks"
	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAddNewAdmin(t *testing.T) {
	t.Run("Test Success Add New Admin", func(t *testing.T) {
		body := strings.NewReader(`
			{
				"username" : "admin2",
				"password" : "admin123",
				"role_id" : 1
			}	
		`)

		req := httptest.NewRequest(http.MethodPost, "/admin", body)
		ctx := context.WithValue(context.Background(), "admin-data", models.AdminDataToken{
			ID:       1,
			Username: "owner",
			Role:     "super admin",
		})

		w := httptest.NewRecorder()

		mockUcase := new(mocks.AdminUcase)
		mockUcase.On("Add", ctx, &models.CreateNewAdminRequest{
			Username: "admin2",
			Password: "admin123",
			Role:     1,
		}).Return(models.Admin{
			ID:       2,
			Username: "admin2",
			Role:     "super admin",
		}, nil)

		handler := AdminHandler{mockUcase}

		handler.AddNewAdmin(w, req.WithContext(ctx))

		res := w.Result()
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		expect := `{"code":201,"data":{"id":2,"username":"admin2","role":"super admin"},"status":"CREATED"}`

		mockUcase.AssertExpectations(t)
		assert.Equal(t, 201, res.StatusCode)
		assert.Equal(t, expect, string(data))
	})

	t.Run("Test Non Authorized", func(t *testing.T) {
		body := strings.NewReader(`
			{
				"username" : "admin2",
				"password" : "admin123",
				"role" : "admin"
			}	
		`)

		req := httptest.NewRequest(http.MethodPost, "/admin", body)
		w := httptest.NewRecorder()

		mockUcase := new(mocks.AdminUcase)
		handler := AdminHandler{mockUcase}

		handler.AddNewAdmin(w, req)

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

func TestAdminHandler_Login(t *testing.T) {
	body := strings.NewReader(`
		{
			"username" : "admin",
			"password" : "admin123"
		}
	`)
	token, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	req := httptest.NewRequest(http.MethodPost, "/admin/login", body)
	w := httptest.NewRecorder()

	mockUcase := new(mocks.AdminUcase)
	mockUcase.On("Login", context.Background(), &models.AdminLoginRequest{
		Username: "admin",
		Password: "admin123",
	}).Return(string(token), nil)

	handler := AdminHandler{mockUcase}

	handler.Login(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	expect := fmt.Sprintf(`{"code":200,"data":{"token":"%v"},"status":"SUCCESS"}`, string(token))

	mockUcase.AssertExpectations(t)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, expect, string(data))

}

func TestGetAdmins(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	w := httptest.NewRecorder()

	mockUcase := new(mocks.AdminUcase)
	mockUcase.On("GetAdmins", context.Background()).Return([]models.Admin{{
		ID:       1,
		Username: "owner",
		Role:     "super admin",
	}, {
		ID:       2,
		Username: "finance",
		Role:     "finance",
	}}, nil)

	handler := AdminHandler{mockUcase}
	handler.GetAdmins(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	expect := `{"code":200,"data":[{"id":1,"username":"owner","role":"super admin"},{"id":2,"username":"finance","role":"finance"}],"status":"SUCCESS"}`
	mockUcase.AssertExpectations(t)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, expect, string(data))
}
