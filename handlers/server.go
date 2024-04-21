package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmorpg-server/db"
)

// add database once ready
type Server struct {
	ListenAddr  string
    store       db.Storage
}

func NewServer(listenAddr string, store db.Storage) *Server {
    return &Server{
    ListenAddr: listenAddr,
    store: store,
    }
}

func (s *Server) RunServer() error {
    router := mux.NewRouter()
    router.HandleFunc("/player/{id}", makeHTTPHandleFunc(s.handleGetPlayerByID))
    router.HandleFunc("/player", makeHTTPHandleFunc(s.handlePlayer))
    log.Println("Server is running on port: " + s.ListenAddr)
    return http.ListenAndServe(s.ListenAddr, router)
}