package models

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PrinceNarteh/tweeter/internal/errs"
)

var (
	UsernameMinLength = 2
	PasswordMinLength = 6
	emailRegex        = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

type LoginInput struct{}

type RegisterInput struct {
	Username        string `json:"username"        validate:"required,min=2"`
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

func (in RegisterInput) Validate() error {
	if len(in.Username) < UsernameMinLength {
		return fmt.Errorf(
			"%w: username not long enough, (%d) characters at least",
			errs.ErrValidation,
			UsernameMinLength,
		)
	}

	if !emailRegex.MatchString(in.Email) {
		return fmt.Errorf("%w: email not valid", errs.ErrValidation)
	}

	if len(in.Password) < PasswordMinLength {
		return fmt.Errorf(
			"%w: password not long enough, (%d) characters at least",
			errs.ErrValidation,
			PasswordMinLength,
		)
	}

	if in.Password != in.ConfirmPassword {
		return fmt.Errorf("%w: passwords do not match", errs.ErrValidation)
	}

	return nil
}
