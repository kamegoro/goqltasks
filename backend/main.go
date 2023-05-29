package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryType,
	Mutation: mutationType,
})

var todos []Todo

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func main() {
	http.Handle("/graphql", handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}))

	log.Println("Server started on http://localhost:8000/graphql")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"todos": &graphql.Field{
			Type: graphql.NewList(todoType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return todos, nil
			},
		},
	},
})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createTodo": &graphql.Field{
			Type: todoType,
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				title := p.Args["title"].(string)
				todo := Todo{
					ID:    "1",
					Title: title,
				}
				todos = append(todos, todo)
				return todo, nil
			},
		},
	},
})

var todoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
	},
})
