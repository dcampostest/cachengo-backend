package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedidosya/golan-rest-simple/handlers"
)

func main() {
	//mux := http.NewServeMux()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.IndexRoute)

	//Category Routes
	router.HandleFunc("/categories", handlers.GetCategory).Methods(http.MethodGet, http.MethodOptions).Name("list Categories")
	router.HandleFunc("/categories", handlers.CreateCategory).Methods(http.MethodPost, http.MethodOptions).Name("Create Categories")
	router.HandleFunc("/deletecategories", handlers.DeleteCategory).Methods(http.MethodPost, http.MethodOptions).Name("delete Categories")
	router.HandleFunc("/updatecategories", handlers.UpdateCategory).Methods(http.MethodPost, http.MethodOptions).Name("Update Categories")
	//Categorie Template Boostrap
	router.HandleFunc("/listcategories", handlers.Temp_listcategories).Methods(http.MethodGet, http.MethodOptions).Name("view list Categories")
	router.HandleFunc("/createcategory", handlers.Temp_createcategorie).Methods(http.MethodGet, http.MethodOptions).Name("view form create category")
	router.HandleFunc("/updatecategory", handlers.Temp_updatecategory).Methods(http.MethodGet, http.MethodOptions).Name("view update category")
	router.HandleFunc("/deletecategory", handlers.Temp_deletecategory).Methods(http.MethodGet, http.MethodOptions).Name("viewdelete category")

	//Product Routes
	router.HandleFunc("/products", handlers.GetProducts).Methods(http.MethodGet, http.MethodOptions).Name("list products")
	router.HandleFunc("/productsbycategory", handlers.GetProductsByCategory).Methods(http.MethodGet, http.MethodOptions).Name("get product list by category")
	router.HandleFunc("/products", handlers.CreateProduct).Methods(http.MethodPost, http.MethodOptions).Name("Create product")
	router.HandleFunc("/deleteproducts", handlers.DeleteProduct).Methods(http.MethodPost, http.MethodOptions).Name("delete product")
	router.HandleFunc("/updateproducts", handlers.UpdateProduct).Methods(http.MethodPost, http.MethodOptions).Name("update prodyct")
	//Categorie Template Boostrap
	router.HandleFunc("/listproducts", handlers.Temp_listproducts).Methods(http.MethodGet, http.MethodOptions).Name("view list products")
	router.HandleFunc("/createproduct", handlers.Temp_createproduct).Methods(http.MethodGet, http.MethodOptions).Name("view create product")
	router.HandleFunc("/updateproduct", handlers.Temp_updateproduct).Methods(http.MethodGet, http.MethodOptions).Name("view update product")
	router.HandleFunc("/deleteproduct", handlers.Temp_deleteproduct).Methods(http.MethodGet, http.MethodOptions).Name("view delete product")

	//Response JSON categories + products by category
	router.HandleFunc("/all", handlers.GetAll).Methods(http.MethodGet, http.MethodOptions).Name("Get all products and categories")

	// srv := &http.Server{
	// 	Addr:    fmt.Sprintf(":%d", 3006),
	// 	Handler: router,
	// }

	fmt.Println("Server started on port ", 3006)
	err := http.ListenAndServeTLS(":3006", "server.pem", "server.key", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
