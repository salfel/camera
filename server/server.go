package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"camera-server/handlers"
	"camera-server/services/broadcast"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
	}

	hub := broadcast.NewHub()
	go hub.Run()

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      handlers.HandleRoutes(hub),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
