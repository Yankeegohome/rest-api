package nlab

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Nlab struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

func New(config *Config) *Nlab {
	return &Nlab{
		config: config,
	}
}

func (s *Nlab) Open() error {

	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return err
	}

	s.db = db

	return nil
}

func (s *Nlab) Close() {
	s.db.Close()
}

func (n *Nlab) User() *UserRepository {
	if n.userRepository != nil {
		return n.userRepository
	}
	n.userRepository = &UserRepository{
		nlab: n,
	}
	return n.userRepository
}
