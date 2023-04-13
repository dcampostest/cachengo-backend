package handlers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/pedidosya/golan-rest-simple/datasource"
)

var templateCreateCategorie = template.Must(template.ParseGlob("./templates/*"))

func GetCategory(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	listAllCategories := datasource.GetListCategories()
	json.NewEncoder(w).Encode(listAllCategories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "POST" {
		name := r.FormValue("categoryName")
		description := r.FormValue("categoryDescription")
		datasource.CreateCategory(name, description)
		http.Redirect(w, r, "/listcategories", 301)
	}
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "POST" {
		id := r.FormValue("id")
		datasource.DeleteCategory(id)
		http.Redirect(w, r, "/listcategories", 301)
	}
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
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
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	templateCreateCategorie.ExecuteTemplate(w, "create_category", nil)
}
func Temp_listcategories(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	listAllCategories := datasource.GetListCategories()
	templateCreateCategorie.ExecuteTemplate(w, "list_categories", listAllCategories)
}

func Temp_updatecategory(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	idCategory := r.URL.Query().Get("id")
	getCategoryByID := datasource.GetCategoryByID(idCategory)
	templateCreateCategorie.ExecuteTemplate(w, "update_category", getCategoryByID)
}

func Temp_deletecategory(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	idCategory := r.URL.Query().Get("id")
	getCategoryByID := datasource.GetCategoryByID(idCategory)
	templateCreateCategorie.ExecuteTemplate(w, "delete_category", getCategoryByID)
}
