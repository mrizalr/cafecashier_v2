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

func (r *mysqlAdminRepository) Add(ctx context.Context, admin *domain.Admin) (insertedID int64, err error) {
	query := `INSERT INTO ADMIN(username, password) VALUES (?,?)`
	sqlRes, err := r.db.ExecContext(ctx, query, admin.Username, admin.Password)
	if err != nil {
		return
	}

	insertedID, err = sqlRes.LastInsertId()
	return
}

func (r *mysqlAdminRepository) FindByID(ctx context.Context, ID int) (res domain.Admin, err error) {
	query := `SELECT id, username FROM admin WHERE id = ?`
	err = r.db.QueryRowContext(ctx, query, ID).Scan(&res.ID, &res.Username)
	return
}
