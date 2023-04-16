package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedidosya/golan-rest-simple/handlers"
)

func setupRoutes(router *mux.Router) {
	// First enable CORS. If you don't need cors, comment the next line
	enableCORS(router)
	router.HandleFunc("/", handlers.IndexRoute)

	//Category Routes
	router.HandleFunc("/categories", handlers.GetCategory).Methods(http.MethodGet)
	router.HandleFunc("/categoriesPost", handlers.CreateCategory).Methods(http.MethodPost)
	router.HandleFunc("/deletecategories", handlers.DeleteCategory).Methods(http.MethodDelete)
	router.HandleFunc("/updatecategories", handlers.UpdateCategory).Methods(http.MethodPut)
	//Categorie Template Boostrap
	router.HandleFunc("/listcategories", handlers.Temp_listcategories).Methods(http.MethodGet)
	router.HandleFunc("/createcategory", handlers.Temp_createcategorie).Methods(http.MethodGet)
	router.HandleFunc("/updatecategory", handlers.Temp_updatecategory).Methods(http.MethodGet)
	router.HandleFunc("/deletecategory", handlers.Temp_deletecategory).Methods(http.MethodGet)

	//Product Routes
	router.HandleFunc("/products", handlers.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/productsbycategory", handlers.GetProductsByCategory).Methods(http.MethodGet)
	router.HandleFunc("/productsPost", handlers.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/deleteproducts", handlers.DeleteProduct).Methods(http.MethodDelete)
	router.HandleFunc("/updateproducts", handlers.UpdateProduct).Methods(http.MethodPut)
	//Categorie Template Boostrap
	router.HandleFunc("/listproducts", handlers.Temp_listproducts).Methods(http.MethodGet)
	router.HandleFunc("/createproduct", handlers.Temp_createproduct).Methods(http.MethodGet)
	router.HandleFunc("/updateproduct", handlers.Temp_updateproduct).Methods(http.MethodGet)
	router.HandleFunc("/deleteproduct", handlers.Temp_deleteproduct).Methods(http.MethodGet)
	//Response JSON categories + products by category
	router.HandleFunc("/all", handlers.GetAll).Methods(http.MethodGet)

}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}
func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

// Helper functions for respond with 200 or 500 code
func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
