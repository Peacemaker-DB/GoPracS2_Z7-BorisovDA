package main

import (
	"log"
	"net/http"
	"os"

	"example.com/gopracs2-z7-borisovda/services/tasks/internal/server"
)

func main() {
	port := os.Getenv("TASKS_PORT")
	if port == "" {
		port = "8082"
	}

	addr := ":" + port
	log.Println("tasks service started on", addr)

	if err := http.ListenAndServe(addr, server.NewRouter()); err != nil {
		log.Fatal(err)
	}
}