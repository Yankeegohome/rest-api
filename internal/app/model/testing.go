package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Login:    "kostia",
		Password: "passowrd",
	}
}
