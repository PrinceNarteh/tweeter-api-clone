package models

import (
	"strings"

	"github.com/PrinceNarteh/tweeter/internal/utils"
)

type LoginInput struct{}

type RegisterInput struct {
	Username        string `json:"username"        validate:"required"`
	Email           string `json:"email"           validate:"required,email"`
	Password        string `json:"password"        validate:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=6"`
}

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
	User        User   `json:"user"`
}

func (in *RegisterInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)

	in.Username = strings.TrimSpace(in.Username)
}

func (in RegisterInput) Validate() map[string]string {
	return utils.ValidateStruct(in)
}
