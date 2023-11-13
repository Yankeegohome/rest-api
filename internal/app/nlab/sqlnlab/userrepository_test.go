package sqlnlab_test

import (
	"github.com/stretchr/testify/assert"
	"rest-api/internal/app/model"
	"rest-api/internal/app/nlab"
	"rest-api/internal/app/nlab/sqlnlab"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlnlab.TestDB(t, databaseURL)
	defer teardown("usr")
	s := sqlnlab.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByLogin(t *testing.T) {
	db, teardown := sqlnlab.TestDB(t, databaseURL)
	defer teardown("usr")
	s := sqlnlab.New(db)
	login := "kostia"
	_, err := s.User().FindByLogin(login)
	assert.EqualError(t, err, nlab.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Login = login
	s.User().Create(u)
	u, err = s.User().FindByLogin(login)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
