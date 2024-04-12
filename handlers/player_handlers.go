/*
Handle players accounts
Create players accounts
delete players accounts
update players accounts
fetch players accounts
*/

package handlers

import "net/http"

type apiFunction func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(fn apiFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (s *Server) handlePlayer(w http.ResponseWriter, r *http.Request) error {
	return nil
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



