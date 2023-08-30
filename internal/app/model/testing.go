package model

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Email:    "bulldog@gmail.com",
		Password: "aboba322",
	}
}
