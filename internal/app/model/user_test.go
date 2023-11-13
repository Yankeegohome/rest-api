package model_test

import (
	"github.com/stretchr/testify/assert"
	"rest-api/internal/app/model"
	"testing"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.Pass)

}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty login",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Login = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Passoword = "short"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Passoword = ""
				return u
			},
			isValid: false,
		},
		{
			name: "valid Pass",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Passoword = ""
				u.Pass = "asdASD123"
				return u
			},
			isValid: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
