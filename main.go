package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/mmorpg-server/db"
	"github.com/mmorpg-server/handlers"
)

func main() {
	store, err := db.NewPostGresStorage()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	server := handlers.NewServer(":3000", store)
	server.RunServer()
	fmt.Println("Server Online!")
}