package server

import (
	"log"
	"net/http"
	"omh-simple-app/database"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()
	startHttpServer()
}

func startHttpServer() {
	r := mux.NewRouter()
	initRoutes(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         os.Getenv("API_SERVER_HOSTNAME") + ":" + os.Getenv("API_SERVER_PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
