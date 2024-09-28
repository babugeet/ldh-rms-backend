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
	router.HandleFunc("/tables", handlers.GetTables).Methods(http.MethodGet)
	http.ListenAndServe(":4000", router)
}
