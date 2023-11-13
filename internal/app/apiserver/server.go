package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"rest-api/internal/app/nlab"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	nlab   nlab.Nlab
}

func newServer(nlab nlab.Nlab) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		nlab:   nlab,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
