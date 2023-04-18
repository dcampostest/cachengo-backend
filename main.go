package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Define routes
	router := mux.NewRouter()
	SetupRoutes(router)
	port := ":8000"

	server := &http.Server{
		Handler: router,
		Addr:    port,
		// timeouts so the server never waits forever...
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server started at %s", port)
	log.Fatal(server.ListenAndServeTLS("/etc/letsencrypt/csr/0000_csr-certbot.pem", "/etc/letsencrypt/csr/0000_csr-certbot.pem"))
}
