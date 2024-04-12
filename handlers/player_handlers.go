/*
Handle players accounts
Create players accounts
delete players accounts
update players accounts
fetch players accounts
*/

package handlers

import (
	"fmt"
	"net/http"
)


func (s *Server) handlePlayer(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAllPlayers(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreatePlayer(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeletePlayer(w, r)
	}
	return fmt.Errorf("Method not allowed %s", r.Method)
}

func (s *Server) handleGetPlayerByID(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleCreatePlayer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleDeletePlayer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleUpdatePlayer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleGetAllPlayers(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// func (s *Server) handleGetPlayerItems(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }

// func (s *Server) handleGetPlayerStats(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }



