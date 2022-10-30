package teststore_test

import (
	"github.com/Kirillznkv/new_rest_api/internal/app/model"
	"github.com/Kirillznkv/new_rest_api/internal/app/store"
	"github.com/Kirillznkv/new_rest_api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()
	u, err := s.User().Find(0)
	assert.Nil(t, u)
	assert.EqualError(t, err, store.ErrUserNotFound.Error())

	u = model.TestUser(t)
	err = s.User().Create(u)
	assert.NoError(t, err)

	u, err = s.User().Find(u.ID)
	assert.NotNil(t, u)
	assert.NoError(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "zenich99@gmail.com"
	u, err := s.User().FindByEmail(email)
	assert.Nil(t, u)
	assert.EqualError(t, err, store.ErrUserNotFound.Error())

	u = model.TestUser(t)
	u.Email = email
	err = s.User().Create(u)
	assert.NoError(t, err)

	u, err = s.User().FindByEmail(email)
	assert.NotNil(t, u)
	assert.NoError(t, err)
}
