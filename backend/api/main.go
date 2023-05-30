package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)

	log.Println("Server started on http://localhost:8000/graphql")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
