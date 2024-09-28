package main

import (
	"fmt"
	"net/http"
	"table-ms/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Started Table Management Microservice")
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetTables).Methods(http.MethodGet)

}
