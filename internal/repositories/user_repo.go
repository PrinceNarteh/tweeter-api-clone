package repositories

import (
	"context"

	"github.com/PrinceNarteh/tweeter/internal/models"
)

var _ UserRepo = (*userRepo)(nil)

type UserRepo interface {
	GetByUsername(context.Context, string) (models.User, error)
	GetByEmail(context.Context, string) (models.User, error)
}

type userRepo struct {
}

func (ur *userRepo) GetByUsername(ctx context.Context, username string) (models.User, error) {
	return models.User{}, nil
}

func (ur *userRepo) GetByEmail(ctx context.Context, username string) (models.User, error) {
	return models.User{}, nil
}
