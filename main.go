package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/iqra-sham/jokes/handler"
)

// Joke struct representing a joke

func main() {
	r := chi.NewRouter()

	r.Post("/graphql", handler.GraphqlHandler)
	fmt.Println("server starting")

	// Use chi router to handle all requests
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
