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

func (r *UserRepository) GetUserBYID(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1;`
	var user models.User
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	query := `SELECT id, name, email FROM users;`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user models.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3;`
	_, err := r.DB.ExecContext(ctx, query, user.Name, user.Email, user.ID)
	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1;`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}
