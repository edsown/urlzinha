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

	db, err := sql.Open("sqlite", "/app.db")
	if err != nil {

		fmt.Errorf("error with database creation %s", err)
	}
	fmt.Println("config loaded")
	repo := repository.NewRepository(db)
	svc := service.NewService()
	mux := http.NewServeMux()
	creationHandler := handler.NewHandler(repo, db, svc)
	mux.HandleFunc("GET /create", func(w http.ResponseWriter, r *http.Request) {
		creationHandler.HandleCreate(w, r)
	})

}
