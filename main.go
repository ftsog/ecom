package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ftsog/ecom/config"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	router := config.NewRouter(r)
	router.Routing()

	server := &http.Server{
		Handler:      router.Route,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("listening on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
