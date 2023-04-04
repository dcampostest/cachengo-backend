package handlers

import (
	"net/http"
	"text/template"
)

var plantilla = template.Must(template.ParseGlob("./templates/*"))

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome API cachengo")
	plantilla.ExecuteTemplate(w, "principal", nil)
}
