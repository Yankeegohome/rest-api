package nlab_test

import (
	"github.com/stretchr/testify/assert"
	"rest-api/internal/app/model"
	"rest-api/internal/app/nlab"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := nlab.Testnlab(t, databaseURL)
	defer teardown("usr")
	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByLogin(t *testing.T) {
	s, teardown := nlab.Testnlab(t, databaseURL)
	defer teardown("usr")

	login := "kostia"
	_, err := s.User().FindByLogin(login)
	assert.Error(t, err)
	u := model.TestUser(t)
	u.Login = login
	s.User().Create(u)
	u, err = s.User().FindByLogin(login)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
