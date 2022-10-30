package teststore

import (
	"github.com/Kirillznkv/new_rest_api/internal/app/model"
	"github.com/Kirillznkv/new_rest_api/internal/app/store"

	_ "github.com/lib/pq"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[int]*model.User),
		}
	}

	return s.userRepository
}
