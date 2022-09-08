package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@mail.ru",
		Password: "user_password",
	}
}
