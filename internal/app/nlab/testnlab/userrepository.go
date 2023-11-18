package testnlab

import (
	"rest-api/internal/app/model"
	"rest-api/internal/app/nlab"
)

type UserRepository struct {
	nlab  *Nlab
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	u.ID = len(r.users) + 1
	r.users[u.ID] = u
	return nil
}

func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	for _, u := range r.users {
		if u.Login == login {
			return u, nil
		}

	}
	return nil, nlab.ErrRecordNotFound

}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, nlab.ErrRecordNotFound
	}
	return u, nil
}
