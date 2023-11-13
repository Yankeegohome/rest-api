package sqlnlab

import (
	"database/sql"
	"rest-api/internal/app/model"
	"rest-api/internal/app/nlab"
)

type UserRepository struct {
	nlab *Nlab
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	return r.nlab.db.QueryRow(
		"INSERT INTO NLAB.USR(ID, LOGIN, PASS, TEXT, STATUS) VALUES (nextval('nlab.s_usr'), $1, $2, $3, $4) RETURNING id",
		u.Login,
		u.Pass,
		u.Text,
		1,
	).Scan(&u.ID)
}

func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	u := &model.User{}
	if err := r.nlab.db.QueryRow(
		"select id, login, text, pass from nlab.usr where login = $1",
		login,
	).Scan(
		&u.ID,
		&u.Login,
		&u.Text,
		&u.Pass,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nlab.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}
