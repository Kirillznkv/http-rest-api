package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

// Store ...
type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// MigrateUP ...
func (s *Store) MigrateUP() error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(s.db, s.config.MigrateDirPath); err != nil {
		return err
	}
	return nil
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close
func (s *Store) Close() {
	s.db.Close()
}

// User ...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
