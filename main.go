package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Server() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handleHome).Methods("GET")
	router.HandleFunc("/api/products", BrowseProduct).Methods("GET")
	return router
}

func main() {
	router := Server()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleHome(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader((http.StatusOK))
	json.NewEncoder(writer).Encode(map[string]interface{}{
		"status":  200,
		"message": "Hello World",
	})
}

func BrowseProduct(writer http.ResponseWriter, request *http.Request) {
	renderJson(writer, map[string]interface{}{
		"message": "products",
	})
}
