package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Fikkie007/rest-api/internal/config"
	"github.com/Fikkie007/rest-api/internal/database"
)

func main() {

	cfg := config.Load()

	db, err := database.NewPostgres(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("Database connected")

	mux := http.NewServeMux()

	log.Println("Server running on http://localhost:" + cfg.App.Port)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, `{"status" : "ok"}`)
	})

	if err := http.ListenAndServe(":"+cfg.App.Port, mux); err != nil {
		log.Println(err)
	}

}
