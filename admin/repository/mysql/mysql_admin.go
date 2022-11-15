package mysql

import (
	"context"
	"database/sql"

	"github.com/mrizalr/cafecashierpt2/domain"
)

type mysqlAdminRepository struct {
	db *sql.DB
}

func NewMysqlArticleRepository(db *sql.DB) *mysqlAdminRepository {
	return &mysqlAdminRepository{db}
}

func (r *mysqlAdminRepository) Add(ctx context.Context, admin *domain.Admin) (int64, error) {
	query := `INSERT INTO admins (username, password, role_id) VALUES (?,?,?)`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	sqlRes, err := stmt.ExecContext(ctx, admin.Username, admin.Password, admin.Role)
	if err != nil {
		return 0, err
	}

	return sqlRes.LastInsertId()
}

func (r *mysqlAdminRepository) FindByID(ctx context.Context, ID int) (domain.Admin, error) {
	result := new(domain.Admin)
	query := `SELECT id, username, role_id FROM admins WHERE id = ?`
	err := r.db.QueryRowContext(ctx, query, ID).Scan(&result.ID, &result.Username, &result.Role)
	return *result, err
}

func (r *mysqlAdminRepository) FindByUsername(ctx context.Context, username string) (domain.Admin, error) {
	result := new(domain.Admin)
	query := `SELECT id, username, password, role_id FROM admins WHERE username = ?`
	err := r.db.QueryRowContext(ctx, query, username).Scan(&result.ID, &result.Username, &result.Password, &result.Role)
	return *result, err
}
