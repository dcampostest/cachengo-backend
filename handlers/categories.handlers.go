package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/pedidosya/golan-rest-simple/datasource"
)

var templateCreateCategorie = template.Must(template.ParseGlob("./templates/*"))

// Helper functions for respond with 200 or 500 code
func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func stringToInt64(s string) (int64, error) {
	numero, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}
	return numero, err
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	listAllCategories := datasource.GetListCategories()
	json.NewEncoder(w).Encode(listAllCategories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("categoryName")
		description := r.FormValue("categoryDescription")
		datasource.CreateCategory(name, description)
		http.Redirect(w, r, "/listcategories", 301)
	}
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	datasource.DeleteCategory(id)
	http.Redirect(w, r, "/listcategories", 301)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("categoryName")
		description := r.FormValue("categoryDescription")
		datasource.UpdateCategory(id, name, description)
		http.Redirect(w, r, "/listcategories", 301)
	}
}

// function category template boostrap
func Temp_createcategorie(w http.ResponseWriter, r *http.Request) {
	templateCreateCategorie.ExecuteTemplate(w, "create_category", nil)
}
func Temp_listcategories(w http.ResponseWriter, r *http.Request) {
	listAllCategories := datasource.GetListCategories()
	templateCreateCategorie.ExecuteTemplate(w, "list_categories", listAllCategories)
}

func Temp_updatecategory(w http.ResponseWriter, r *http.Request) {
	idCategory := r.URL.Query().Get("id")
	getCategoryByID := datasource.GetCategoryByID(idCategory)
	templateCreateCategorie.ExecuteTemplate(w, "update_category", getCategoryByID)
}

func Temp_deletecategory(w http.ResponseWriter, r *http.Request) {
	idCategory := r.URL.Query().Get("id")
	getCategoryByID := datasource.GetCategoryByID(idCategory)
	templateCreateCategorie.ExecuteTemplate(w, "delete_category", getCategoryByID)
}
