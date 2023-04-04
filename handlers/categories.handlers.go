package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/pedidosya/golan-rest-simple/data"
	"github.com/pedidosya/golan-rest-simple/models"
)

var templateCreateCategorie = template.Must(template.ParseGlob("./templates/*"))

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid category.")
	}
	json.Unmarshal(reqBody, &newCategory)
	newCategory.ID = len(data.Categories) + 1
	data.Categories = append(data.Categories, newCategory)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for i, category := range data.Categories {
		if category.ID == categoryId {
			w.WriteHeader(http.StatusOK)
			data.Categories = append(data.Categories[:i], data.Categories[i+1:]...)
			fmt.Fprintf(w, "Delete product ID:%v - Name:%s", category.ID, category.Name)
			return
		} else if i == len(data.Categories)-1 {
			fmt.Fprintf(w, "No found ID:%v", categoryId)
		}
	}
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	var updateCategory models.Category
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert enter valid Data")
	}
	json.Unmarshal(reqBody, &updateCategory)
	for i, category := range data.Categories {
		if category.ID == categoryId {
			updateCategory.ID = categoryId
			data.Categories[i] = updateCategory
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if updateCategory.ID != 0 {
		json.NewEncoder(w).Encode(updateCategory)
	} else {
		fmt.Fprintf(w, "No se encontro el id")
	}
}

func GetCategoryID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Id Invalido")
		return
	}
	for i, category := range data.Categories {
		if category.ID == categoryId {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(category)
			return
		} else if i == len(data.Products)-1 {
			fmt.Fprintf(w, "No found ID:%v", categoryId)
		}
	}
}

// function categorie template boostrap
func Temp_createcategorie(w http.ResponseWriter, r *http.Request) {
	templateCreateCategorie.ExecuteTemplate(w, "create_category", nil)
}
func Temp_listcategories(w http.ResponseWriter, r *http.Request) {
	templateCreateCategorie.ExecuteTemplate(w, "list_categories", nil)
}
