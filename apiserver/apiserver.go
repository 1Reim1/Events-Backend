package apiserver

import (
	"Events-Backend/config"
	"Events-Backend/storage"
	"github.com/gorilla/mux"
	"net/http"
)

type APIServer struct {
	config  *config.Config
	router  *mux.Router
	storage storage.Storage
}

func NewAPIServer(config *config.Config, storage storage.Storage) *APIServer {
	return &APIServer{
		config:  config,
		router:  mux.NewRouter(),
		storage: storage,
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureRouter() {

}
