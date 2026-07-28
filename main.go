package main

import (
	"fmt"
	"net/http"

	"database/sql"
	"github.com/edsown/urlzinha/business/repository"
	"github.com/edsown/urlzinha/handler"
	"github.com/edsown/urlzinha/service"
	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "./app.db")
	if err != nil {
		fmt.Errorf("error with database creation %s", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Errorf("error pinging the database %w", err)
	}

	query := `CREATE TABLE IF NOT EXISTS urls ( id INTEGER PRIMARY KEY AUTOINCREMENT, original_url TEXT NOT NULL, short_url TEXT UNIQUE, created_at DATETIME DEFAULT CURRENT_TIMESTAMP )`
	_, err = db.Exec(query)
	if err != nil {
		fmt.Errorf("error creating table %w", err)
	}

	fmt.Println("config loaded")
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	mux := http.NewServeMux()
	creationHandler := handler.NewHandler(svc)
	mux.HandleFunc("GET /shorten", func(w http.ResponseWriter, r *http.Request) {
		creationHandler.HandleCreate(w, r)
	})
	http.ListenAndServe(":8080", mux)

}
