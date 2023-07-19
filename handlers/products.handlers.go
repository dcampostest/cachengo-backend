package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"

	"github.com/pedidosya/golan-rest-simple/datasource"
	"github.com/pedidosya/golan-rest-simple/models"
)

type TemplateData struct {
	Product    models.Product
	Categories []models.Category
}

var templateProducts = template.Must(template.ParseGlob("./templates/*"))

func GetProducts(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	listAllProducts := datasource.GetListProducts()
	json.NewEncoder(w).Encode(listAllProducts)
}

func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	listAllProducts := datasource.GetListProductsByCategory()
	json.NewEncoder(w).Encode(listAllProducts)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "POST" {
		name := r.FormValue("productName")
		description := r.FormValue("productDescription")
		price := r.FormValue("productPrice")
		categoryName := r.FormValue("productNameCategory")
		idCategory := datasource.GetCategoryByName(categoryName)
		datasource.CreateProduct(name, description, price, strconv.Itoa(idCategory.ID))
		http.Redirect(w, r, "/listproducts", 301)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "POST" {
		id := r.FormValue("id")
		datasource.DeleteProduct(id)
		http.Redirect(w, r, "/listproducts", 301)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("productName")
		description := r.FormValue("productDescription")
		price := r.FormValue("productPrice")
		categoryName := r.FormValue("productNameCategory")
		idCategory := datasource.GetCategoryByName(categoryName)
		datasource.UpdateProduct(id, name, description, price, strconv.Itoa(idCategory.ID))
		http.Redirect(w, r, "/listproducts", 301)
	}
}

// function Products template boostrap
func Temp_createproduct(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	listAllCategories := datasource.GetListCategories()
	templateProducts.ExecuteTemplate(w, "create_product", listAllCategories)
}
func Temp_listproducts(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	listAllProducts := datasource.GetListProducts()
	templateProducts.ExecuteTemplate(w, "list_products", listAllProducts)
}

func Temp_updateproduct(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	m := map[string]interface{}{}
	m["product"] = models.Product{}
	m["categories"] = []models.Category{}

	idProduct := r.URL.Query().Get("id")
	m["product"] = datasource.GetProductByID(idProduct)
	m["categories"] = datasource.GetListCategories()
	templateProducts.ExecuteTemplate(w, "update_product", m)
}

func Temp_deleteproduct(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	idProduct := r.URL.Query().Get("id")
	getProductByID := datasource.GetProductByID(idProduct)
	templateProducts.ExecuteTemplate(w, "delete_product", getProductByID)
}
