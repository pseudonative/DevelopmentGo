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
