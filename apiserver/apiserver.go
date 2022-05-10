package apiserver

import (
	"Events-Backend/config"
	"github.com/gorilla/mux"
	"net/http"
)

type APIServer struct {
	config *config.Config
	router *mux.Router
}

func NewAPIServer(config *config.Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureRouter() {

}
