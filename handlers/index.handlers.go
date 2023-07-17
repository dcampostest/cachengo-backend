package handlers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/pedidosya/golan-rest-simple/datasource"
)

var plantilla = template.Must(template.ParseGlob("./templates/*"))

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome API cachengo")
	plantilla.ExecuteTemplate(w, "principal", nil)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "Origin,X-Requested-With,Content-Type,Accept,content-type,application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	listAllProductsByCategories := datasource.GetAll()
	json.NewEncoder(w).Encode(listAllProductsByCategories)
}
