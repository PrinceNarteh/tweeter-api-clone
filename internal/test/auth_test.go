package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PrinceNarteh/tweeter/internal/errs"
	"github.com/PrinceNarteh/tweeter/internal/models"
)

func TestRegisterInput_Sanitize(t *testing.T) {
	input := models.RegisterInput{
		Username:        "  John Doe  ",
		Email:           "JOHN.doe@email.com",
		Password:        "secret_password",
		ConfirmPassword: "secret_password",
	}

	want := models.RegisterInput{
		Username:        "John Doe",
		Email:           "john.doe@email.com",
		Password:        "secret_password",
		ConfirmPassword: "secret_password",
	}

	input.Sanitize()

	require.Equal(t, want, input)
}

func TestRegisterInput_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		input models.RegisterInput
		err   error
	}{
		{
			name: "valid input",
			input: models.RegisterInput{
				Username:        "John Doe",
				Email:           "john.doe@email.com",
				Password:        "secret_password",
				ConfirmPassword: "secret_password",
			},
			err: nil,
		},
		{
			name: "invalid email",
			input: models.RegisterInput{
				Username:        "John Doe",
				Email:           "john.doe@email",
				Password:        "secret_password",
				ConfirmPassword: "secret_password",
			},
			err: errs.ErrValidation,
		},
		{
			name: "username too short",
			input: models.RegisterInput{
				Username:        "J",
				Email:           "john.doe@email.com",
				Password:        "secret_password",
				ConfirmPassword: "secret_password",
			},
			err: errs.ErrValidation,
		},
		{
			name: "passwords do not match",
			input: models.RegisterInput{
				Username:        "John Doe",
				Email:           "john.doe@email.com",
				Password:        "secret_password",
				ConfirmPassword: "wrong_password",
			},
			err: errs.ErrValidation,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
