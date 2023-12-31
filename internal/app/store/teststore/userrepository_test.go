package teststore_test

import (
	"github.com/SharpDevOps10/GoPractice/internal/app/model"
	"github.com/SharpDevOps10/GoPractice/internal/app/store"
	"github.com/SharpDevOps10/GoPractice/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.ID)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "bulldog@gmail.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u1 := model.TestUser(t)
	u1.Email = email

	s.User().Create(u1)

	u2, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindById(t *testing.T) {
	s := teststore.New()
	u1 := model.TestUser(t)
	s.User().Create(u1)

	u2, err := s.User().FindById(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
