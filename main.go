package main 

import (
	"./queries"
	"./mutations"
	"./security"

	"net/http"
	"log"

	"github.com/graphql-go/handler"
	"github.com/graphql-go/graphql"
)

func main() {

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: 	"RootQuery",
			Fields: queries.GetRootFields(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:	"RootMutation",
			Fields: mutations.GetRootFields(),
		}),
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create new schema, error : %v", err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/", security.Handle(httpHandler))
	log.Print("ready: listening...\n")

	http.ListenAndServe(":8383", nil)
}