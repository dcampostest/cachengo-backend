package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedidosya/golan-rest-simple/handlers"
)

func SetupRoutes(routes *mux.Router) {
	// First enable CORS. If you don't need cors, comment the next line
	EnableCORS(routes)
	routes.HandleFunc("/", handlers.IndexRoute)

	//Category Routes
	routes.HandleFunc("/categories", handlers.GetCategory).Methods(http.MethodGet, http.MethodOptions)
	routes.HandleFunc("/categoriesPost", handlers.CreateCategory).Methods(http.MethodPost, http.MethodOptions)
	routes.HandleFunc("/deletecategories/{id}", handlers.DeleteCategory).Methods(http.MethodDelete, http.MethodOptions)
	routes.HandleFunc("/updatecategories/{id}", handlers.UpdateCategory).Methods(http.MethodPut, http.MethodOptions)
	//Categorie Template Boostrap
	routes.HandleFunc("/listcategories", handlers.Temp_listcategories).Methods(http.MethodGet, http.MethodOptions)
	routes.HandleFunc("/createcategory", handlers.Temp_createcategorie).Methods(http.MethodGet, http.MethodOptions)
	routes.HandleFunc("/updatecategory", handlers.Temp_updatecategory).Methods(http.MethodGet, http.MethodOptions)
	routes.HandleFunc("/deletecategory", handlers.Temp_deletecategory).Methods(http.MethodGet, http.MethodOptions)

	//Product Routes
	routes.HandleFunc("/products", handlers.GetProducts).Methods(http.MethodGet, http.MethodOptions)
	routes.HandleFunc("/productsbycategory/{id}", handlers.GetProductsByCategory).Methods(http.MethodGet, http.MethodOptions)
	routes.HandleFunc("/productsPost", handlers.CreateProduct).Methods(http.MethodPost, http.MethodOptions)
	routes.HandleFunc("/deleteproducts/{id}", handlers.DeleteProduct).Methods(http.MethodDelete, http.MethodOptions)
	routes.HandleFunc("/updateproducts/{id}", handlers.UpdateProduct).Methods(http.MethodPut, http.MethodOptions)
	//Categorie Template Boostrap
	routes.HandleFunc("/listproducts", handlers.Temp_listproducts).Methods(http.MethodGet, http.MethodOptions)
	routes.HandleFunc("/createproduct", handlers.Temp_createproduct).Methods(http.MethodGet, http.MethodOptions)
	routes.HandleFunc("/updateproduct", handlers.Temp_updateproduct).Methods(http.MethodGet, http.MethodOptions)
	routes.HandleFunc("/deleteproduct", handlers.Temp_deleteproduct).Methods(http.MethodGet, http.MethodOptions)
	//Response JSON categories + products by category
	routes.HandleFunc("/all", handlers.GetAll).Methods(http.MethodGet)
}

func EnableCORS(routes *mux.Router) {
	routes.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	}).Methods(http.MethodOptions)
	routes.Use(MiddlewareCors)
}
func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}
