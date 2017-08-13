package main 

import (
	"./queries"
	"./mutations"
    "./security"
    "./database"

	"net/http"
    "log"
    "database/sql"

	"github.com/graphql-go/handler"
    "github.com/graphql-go/graphql"
    _ "github.com/lib/pq"
)

func main() {
    var err error
    database.DBCon, err = sql.Open("postgres", "postgres://postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	schemaConfig := graphql.SchemaConfig{
		Query:      queries.QueryType,
		Mutation:   mutations.MutationType,
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
