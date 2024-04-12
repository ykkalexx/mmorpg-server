package handlers

import "github.com/gorilla/mux"

// add database once ready
type Server struct {
	listenAddr string
}

func newServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (s *Server) runServer() error {
	router := mux.NewRouter()
	router.handleFunc("/player", s.handlePlayer).Methods("POST")
}