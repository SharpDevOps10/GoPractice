package sqlstore_test

import (
	"github.com/SharpDevOps10/GoPractice/internal/app/model"
	"github.com/SharpDevOps10/GoPractice/internal/app/store"
	"github.com/SharpDevOps10/GoPractice/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("userss")
	s := sqlstore.New(db)

	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("userss")

	s := sqlstore.New(db)
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
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("userss")

	s := sqlstore.New(db)

	u1 := model.TestUser(t)
	s.User().Create(u1)

	u2, err := s.User().FindById(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)

}
