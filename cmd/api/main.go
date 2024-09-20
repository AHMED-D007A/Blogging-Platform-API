package main

import (
	"log"

	"github.com/AHMED-D007A/Blogging-Platform-API/internal/server"
)

func main() {
	server := server.NewAPIServer(":4000", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
