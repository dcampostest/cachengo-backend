package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pedidosya/golan-rest-simple/data"
	"github.com/pedidosya/golan-rest-simple/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid product.")
	}
	json.Unmarshal(reqBody, &newProduct)
	newProduct.ID = len(data.Products) + 1
	data.Products = append(data.Products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for i, product := range data.Products {
		if product.ID == productId {
			w.WriteHeader(http.StatusOK)
			data.Products = append(data.Products[:i], data.Products[i+1:]...)
			fmt.Fprintf(w, "Delete product ID:%v - Name:%s", product.ID, product.Name)
			return
		} else if i == len(data.Products)-1 {
			fmt.Fprintf(w, "No found ID:%v", productId)
		}
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	var updateProduct models.Product
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert enter valid Data")
	}
	json.Unmarshal(reqBody, &updateProduct)
	for i, product := range data.Products {
		if product.ID == productId {
			updateProduct.ID = productId
			data.Products[i] = updateProduct
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if updateProduct.ID != 0 {
		json.NewEncoder(w).Encode(updateProduct)
	} else {
		fmt.Fprintf(w, "No se encontro el id")
	}
}

func GetProductId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Id Invalido")
		return
	}
	for i, product := range data.Products {
		if product.ID == productId {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(product)
			return
		} else if i == len(data.Products)-1 {
			fmt.Fprintf(w, "No found ID:%v", productId)
		}
	}
}

// function categorie template boostrap
func Temp_createproduct(w http.ResponseWriter, r *http.Request) {
	templateCreateCategorie.ExecuteTemplate(w, "create_product", nil)
}
func Temp_listproducts(w http.ResponseWriter, r *http.Request) {
	templateCreateCategorie.ExecuteTemplate(w, "list_products", nil)
}
