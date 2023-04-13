package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedidosya/golan-rest-simple/handlers"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodDelete,
			http.MethodPut,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	router.HandleFunc("/", handlers.IndexRoute)

	//Category Routes
	router.HandleFunc("/categories", handlers.GetCategory).Methods("GET", "OPTIONS")
	router.HandleFunc("/categories", handlers.CreateCategory).Methods("POST", "OPTIONS")
	router.HandleFunc("/deletecategories", handlers.DeleteCategory).Methods("POST", "OPTIONS")
	router.HandleFunc("/updatecategories", handlers.UpdateCategory).Methods("POST", "OPTIONS")
	//Categorie Template Boostrap
	router.HandleFunc("/listcategories", handlers.Temp_listcategories)
	router.HandleFunc("/createcategory", handlers.Temp_createcategorie)
	router.HandleFunc("/updatecategory", handlers.Temp_updatecategory)
	router.HandleFunc("/deletecategory", handlers.Temp_deletecategory)

	//Product Routes
	router.HandleFunc("/products", handlers.GetProducts).Methods("GET", "OPTIONS")
	router.HandleFunc("/productsbycategory", handlers.GetProductsByCategory).Methods("GET", "OPTIONS")
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST", "OPTIONS")
	router.HandleFunc("/deleteproducts", handlers.DeleteProduct).Methods("POST", "OPTIONS")
	router.HandleFunc("/updateproducts", handlers.UpdateProduct).Methods("POST", "OPTIONS")
	//Categorie Template Boostrap
	router.HandleFunc("/listproducts", handlers.Temp_listproducts)
	router.HandleFunc("/createproduct", handlers.Temp_createproduct)
	router.HandleFunc("/updateproduct", handlers.Temp_updateproduct)
	router.HandleFunc("/deleteproduct", handlers.Temp_deleteproduct)

	fmt.Println("Server started on port ", 3000)
	handler := cors.Handler(router)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
