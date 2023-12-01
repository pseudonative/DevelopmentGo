package repository

import (
	"context"
	"database/sql"

	"github.com/pseudonative/my-rest-api/pkg/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	return nil, nil
}
