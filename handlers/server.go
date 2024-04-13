package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// add database once ready
type Server struct {
	ListenAddr string
}

func NewServer(listenAddr string) *Server {
    return &Server{ListenAddr: listenAddr}
}

func (s *Server) RunServer() error {
    router := mux.NewRouter()
    router.HandleFunc("/player/{id}", makeHTTPHandleFunc(s.handleGetPlayerByID))
    router.HandleFunc("/player", makeHTTPHandleFunc(s.handlePlayer))
    log.Println("Server is running on port: " + s.ListenAddr)
    return http.ListenAndServe(s.ListenAddr, router)
}