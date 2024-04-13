/*
Handle players accounts
Create players accounts
delete players accounts
update players accounts
fetch all players accounts
*/

package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmorpg-server/models"
)


func (s *Server) handlePlayer(w http.ResponseWriter, r *http.Request) error {
    if r.Method == "GET" {
        return s.handleGetPlayerByID(w, r)
    }
    if r.Method == "POST" {
        return s.handleCreatePlayer(w, r)
    }
    if r.Method == "DELETE" {
        return s.handleDeletePlayer(w, r)
    }
    if r.Method == "PUT" {
        return s.handleUpdatePlayer(w, r)
    }
    return fmt.Errorf("Method not allowed %s", r.Method)
}

func (s *Server) handleGetPlayerByID(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	// db.get(id)
	fmt.Println("ID: ", id)

	return WriteJSON(w, http.StatusOK, &models.Player{})
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



