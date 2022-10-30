package sqlstore_test

import (
	"github.com/Kirillznkv/new_rest_api/internal/app/model"
	"github.com/Kirillznkv/new_rest_api/internal/app/store"
	"github.com/Kirillznkv/new_rest_api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u, err := s.User().Find(-1)
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
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
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
