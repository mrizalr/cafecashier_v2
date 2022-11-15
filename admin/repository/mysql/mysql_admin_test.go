package mysql

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestAdd(t *testing.T) {
	admin := &domain.Admin{
		Username: "Cashier",
		Password: "Cashier123",
		Role:     3,
	}

	db, mock := NewMock()
	repo := &mysqlAdminRepository{db}

	t.Run("Test Success Add", func(t *testing.T) {
		query := `INSERT INTO ADMIN(username, password, role) VALUES (?,?,?)`
		prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
		prep.ExpectExec().WithArgs(admin.Username, admin.Password, admin.Role).WillReturnResult(sqlmock.NewResult(1, 1))

		lastID, err := repo.Add(context.Background(), admin)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), lastID)
	})

	t.Run("Test Fail Add", func(t *testing.T) {
		query := `INSERT INTO ADMIN(username, password, role) VALUES (?,?,?)`
		prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
		prep.ExpectExec().WithArgs(admin.Username, admin.Password, admin.Role).WillReturnError(errors.New("Duplicate entry"))

		_, err := repo.Add(context.Background(), admin)
		assert.Error(t, err)
	})
}

func TestFindByID(t *testing.T) {
	db, mock := NewMock()
	repo := &mysqlAdminRepository{db}

	t.Run("Test Found Data", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "role"}).
			AddRow(1, "admin", 1)

		query := `SELECT id, username, role FROM admin WHERE id = ?`
		mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

		admin, err := repo.FindByID(context.Background(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, admin)
	})

	t.Run("Test Not Found Data", func(t *testing.T) {
		query := `SELECT id, username, role FROM admin WHERE id = ?`
		rows := sqlmock.NewRows([]string{"id", "username", "role"})
		mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

		admin, err := repo.FindByID(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, domain.Admin{}, admin)
	})
}
