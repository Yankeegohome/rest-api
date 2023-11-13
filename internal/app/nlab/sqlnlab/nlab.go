package sqlnlab

import (
	"database/sql"
	_ "github.com/lib/pq"
	"rest-api/internal/app/nlab"
)

type Nlab struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Nlab {
	return &Nlab{
		db: db,
	}
}

func (n *Nlab) User() nlab.UserRepository {
	if n.userRepository != nil {
		return n.userRepository
	}
	n.userRepository = &UserRepository{
		nlab: n,
	}
	return n.userRepository
}
