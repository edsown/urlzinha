package main

import (
	"fmt"
	"net/http"

	"github.com/edsown/urlzinha/handler"
)

func main() {

	fmt.Println("config loaded")
	mux := http.NewServeMux()
	creationHandler := handler.NewHandler()
	mux.Handle("/create", creationHandler)

}
