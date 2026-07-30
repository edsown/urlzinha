package main

import (
	"net/http"

	"database/sql"
	"github.com/edsown/urlzinha/business/repository"
	"github.com/edsown/urlzinha/handler"
	"github.com/edsown/urlzinha/service"
	"log"
	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "./app.db")
	// TODO: fly.io
	if err != nil {
		log.Fatal("error creating the database %w", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("error pinging the server", err)
	}

	query := `CREATE TABLE IF NOT EXISTS urls ( id INTEGER PRIMARY KEY AUTOINCREMENT, original_url TEXT NOT NULL, short_url TEXT UNIQUE, created_at DATETIME DEFAULT CURRENT_TIMESTAMP )`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("error creating table %w", err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	mux := http.NewServeMux()
	creationHandler := handler.NewHandler(svc)
	mux.HandleFunc("GET /shorten", func(w http.ResponseWriter, r *http.Request) {
		creationHandler.HandleCreate(w, r)
	})
	log.Print("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error starting server: %v", err)
	}

}
