package mysql

import (
	"context"
	"database/sql"

	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/mrizalr/cafecashierpt2/models"
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

func (r *mysqlAdminRepository) FindByID(ctx context.Context, ID int) (models.Admin, error) {
	result := new(models.Admin)
	query := `SELECT a.id, a.username, a.role FROM admins a JOIN admin_roles ar ON a.role_id = ar.id WHERE id = ?`
	err := r.db.QueryRowContext(ctx, query, ID).Scan(&result.ID, &result.Username, &result.Role)
	return *result, err
}

func (r *mysqlAdminRepository) FindByUsername(ctx context.Context, username string) (models.Admin, error) {
	result := new(models.Admin)
	query := `SELECT a.id, a.username, a.password, ar.role FROM admins a JOIN admin_roles ar ON a.role_id = ar.id WHERE username = ?`
	err := r.db.QueryRowContext(ctx, query, username).Scan(&result.ID, &result.Username, &result.Password, &result.Role)
	return *result, err
}

func (r *mysqlAdminRepository) FindAll(ctx context.Context) ([]models.Admin, error) {
	result := []models.Admin{}
	query := `SELECT a.id, a.username, ar.role FROM admins a JOIN admin_roles ar ON a.role_id = ar.id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		admin := models.Admin{}
		err := rows.Scan(&admin.ID, &admin.Username, &admin.Role)
		if err != nil {
			return nil, err
		}

		result = append(result, admin)
	}

	return result, nil
}

func (r *mysqlAdminRepository) FindAdminRoleByID(ctx context.Context, ID int) (string, error) {
	role := ""
	query := `SELECT role FROM admin_roles WHERE id = ?`
	err := r.db.QueryRowContext(ctx, query, ID).Scan(&role)

	return role, err
}
