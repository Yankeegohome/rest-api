package testnlab

import (
	"rest-api/internal/app/model"
	"rest-api/internal/app/nlab"
)

type UserRepository struct {
	nlab  *Nlab
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	r.users[u.Login] = u
	u.ID = len(r.users)
	return nil
}

func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	u, ok := r.users[login]
	if !ok {
		return nil, nlab.ErrRecordNotFound
	}
	return u, nil
}
