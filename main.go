package main 

import (
	"./queries"

	"net/http"
	"log"

	"github.com/graphql-go/handler"
	"github.com/graphql-go/graphql"
)

func httpHandler() {
	log.Print("oui");
}

func main() {

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: 	"RootQuery",
			Fields: queries.GetRootFields(),
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

	http.Handle("/", httpHandler)
	log.Print("ready: listening...\n")

	http.ListenAndServe(":8383", nil)
}