package model

import (
	"fmt"
	"log"
	"time"

	"github.com/graphql-go/graphql"
)

type Message struct {
	id        int       `json:"id"`
	text      string    `json:"text"`
	timestamp time.Time `json:"timestamp"`
}

var (
	schema graphql.Schema

	messageType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Message",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"timestamp": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	})

	rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"Messages": &graphql.Field{
				Type:        graphql.NewList(messageType),
				Description: "Messages in the last 10 minutes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return []Message{}, nil
				},
			},
		},
	})

	rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createMessage": &graphql.Field{
				Type:        messageType,
				Description: "Create new message",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"text": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var (
						ok   bool
						id   int
						text string
					)
					// marshall and cast the argument value
					id, ok = params.Args["id"].(int)
					if !ok {
						log.Fatal("Error marshall and cast the argument value id")
					}
					text, ok = params.Args["text"].(string)
					if !ok {
						log.Fatal("Error marshall and cast the argument value text")
					}

					// perform mutation operation here
					// for e.g. create a Todo and save to DB.
					newMessage := Message{
						id:   id,
						text: text,
					}

					//TodoList = append(TodoList, newTodo)

					// return the new Todo object that we supposedly save to DB
					// Note here that
					// - we are returning a `Todo` struct instance here
					// - we previously specified the return Type to be `todoType`
					// - `Todo` struct maps to `todoType`, as defined in `todoType` ObjectConfig`
					return newMessage, nil
				},
			},
		},
	})
)

func init() {
	var err error
	// define schema, with our rootQuery and rootMutation
	schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		log.Fatalln("Error define GraphQL schema", err)
	}
}

func GraphQL(query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
