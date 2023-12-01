package repository

import (
	"context"
	"database/sql"

	"github.com/pseudonative/my-rest-api/pkg/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(ctx context.Context, user models.User) (int, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.DB.QueryRowContext(ctx, query, user.Name, user.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
