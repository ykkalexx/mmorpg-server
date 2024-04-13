package main

import (
	"fmt"

	"github.com/mmorpg-server/handlers"
)

func main() {
	server := handlers.NewServer(":3000")
	server.RunServer()
	fmt.Println("Server Online!")
}