package store_test

import (
	"github.com/SharpDevOps10/GoPractice/internal/app/model"
	"github.com/SharpDevOps10/GoPractice/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("userss")

	u, err := s.User().Create(&model.User{
		Email: "bulldog@gmail.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("userss")
	email := "bulldog@gmail.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

}
