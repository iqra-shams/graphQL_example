package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/iqra-sham/jokes/assests"
)

var jokeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Joke",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.String},
			"title": &graphql.Field{Type: graphql.String},
			"text":  &graphql.Field{Type: graphql.String},
		},
	},
)

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"joke": &graphql.Field{
				Type: jokeType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if ok {
						// Find joke by ID
						for _, j := range assests.Jokes {
							if j.ID == id {
								return j, nil
							}
						}
					}
					return nil, nil
				},
			},
			"jokes": &graphql.Field{
				Type: graphql.NewList(jokeType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Return all jokes
					return assests.Jokes , nil
				},
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: rootQuery,
	},
)
