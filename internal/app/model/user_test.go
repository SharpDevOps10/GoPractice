package model_test

import (
	"github.com/SharpDevOps10/GoPractice/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	testCase := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "validate",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "encrypted password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedpassword"
				return u
			},
			isValid: true,
		},
		{
			name: "gde email?",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "gde password?",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "short"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}

}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)

}
