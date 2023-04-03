package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedidosya/golan-rest-simple/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	//Index Routes
	router.HandleFunc("/", handlers.IndexRoute)

	//Category Routes
	router.HandleFunc("/categories", handlers.GetCategory).Methods("GET")
	router.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", handlers.GetCategoryID).Methods("GET")
	router.HandleFunc("/categories/{id}", handlers.DeleteCategory).Methods("DELETE")
	router.HandleFunc("/categories/{id}", handlers.UpdateCategory).Methods("PUT")

	//Product Routes
	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handlers.GetProductId).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")

	fmt.Println("Server started on port ", 3000)
	log.Fatal(http.ListenAndServe(":3000", router))
}
