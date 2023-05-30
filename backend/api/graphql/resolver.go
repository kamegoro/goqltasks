package graphql

import (
	"github.com/graphql-go/graphql"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var todos []Todo

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
					ID:    "1", // 仮のIDを設定
					Title: title,
				}
				todos = append(todos, todo)
				return todo, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryType,
	Mutation: mutationType,
})
