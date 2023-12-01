package services

import (
	"context"

	"github.com/pseudonative/my-rest-api/internal/repository"
	"github.com/pseudonative/my-rest-api/pkg/models"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (us *UserService) GetUser(ctx context.Context, id int) (*models.User, error) {
	return us.Repo.GetUserByID(ctx, id)
}
