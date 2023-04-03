package handlers

import (
	"fmt"
	"net/http"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome API cachengo")
}
