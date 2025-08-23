package services

import (
	"context"

	"github.com/PrinceNarteh/tweeter/internal/models"
)

type AuthService interface {
	Login(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, input models.RegisterInput) (models.AuthResponse, error)
}

var _ AuthService = (*authService)(nil)

type authService struct {
	// Add any dependencies needed for the service, such as a user repository
	// userRepo UserRepository
}

func (s *authService) Login(ctx context.Context, username, password string) (string, error) {
	// Implement login logic here
	// For example, check the username and password against the database
	return "", nil // Replace with actual token or error
}

func (s *authService) Register(ctx context.Context, input models.RegisterInput) (models.AuthResponse, error) {
	input.Sanitize()

	if err := input.Validate(); err != nil {
		return models.AuthResponse{}, err
	}

	response := models.AuthResponse{}
	return response, nil // Replace with actual user ID or error
}
