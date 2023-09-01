package model_test

import (
	"github.com/SharpDevOps10/GoPractice/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.Validate())
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)

}
