package testnlab

import (
	"rest-api/internal/app/model"
	"rest-api/internal/app/nlab"
)

type Nlab struct {
	userRepository *UserRepository
}

func New() *Nlab {
	return &Nlab{}
}

func (n *Nlab) User() nlab.UserRepository {
	if n.userRepository != nil {
		return n.userRepository
	}
	n.userRepository = &UserRepository{
		nlab:  n,
		users: make(map[int]*model.User),
	}
	return n.userRepository
}
