package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "modernc.org/sqlite"

	"github.com/edsown/urlzinha/business/repository"
	"github.com/edsown/urlzinha/handler"
	"github.com/edsown/urlzinha/service"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	dbPath := getEnv("DB_PATH", "./app.db")
	port := getEnv("PORT", "8080")

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("error pinging database: %v", err)
	}

	if err := initSchema(db); err != nil {
		log.Fatalf("error initializing schema: %v", err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /shorten", func(w http.ResponseWriter, r *http.Request) {
		h.HandleCreate(w, r)
	})
	mux.HandleFunc("GET /{shortUrl}", func(w http.ResponseWriter, r *http.Request) {
		h.HandleRetrieve(w, r)
	})

	log.Printf("Starting server on :%s", port)
	if err := http.ListenAndServe(":"+port, corsMiddleware(mux)); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func initSchema(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		original_url TEXT NOT NULL,
		short_url TEXT UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := db.Exec(query)
	return err
}
