package services

import (
	"context"

	"github.com/pseudonative/my-rest-api/internal/repository"
	"github.com/pseudonative/my-rest-api/pkg/models"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, user models.User) (int, error) {
	return s.Repo.CreateUser(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id int) (*models.User, error) {
	return s.Repo.GetUserBYID(ctx, id)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return s.Repo.GetAllUsers(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, user models.User) error {
	return s.Repo.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.Repo.DeleteUser(ctx, id)
}
