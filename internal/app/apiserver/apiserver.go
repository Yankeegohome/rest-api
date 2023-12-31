package apiserver

import (
	"database/sql"
	"github.com/gorilla/sessions"
	"net/http"
	"rest-api/internal/app/nlab/sqlnlab"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	nlab := sqlnlab.New(db)
	sessionNlab := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(nlab, sessionNlab)
	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
