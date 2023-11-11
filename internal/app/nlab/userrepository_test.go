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
	u, err := s.User().Create(&model.User{
		Login:  "kostia",
		Pass:   "admin",
		Text:   "Ян Константин Эдуардович",
		Status: 1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByLogin(t *testing.T) {
	s, teardown := nlab.Testnlab(t, databaseURL)
	defer teardown("usr")

	login := "kostia"
	_, err := s.User().FindByLogin(login)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Login:  "kostia",
		Pass:   "admin",
		Text:   "Ян Константин Эдуардович",
		Status: 1,
	})

	u, err := s.User().FindByLogin(login)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
