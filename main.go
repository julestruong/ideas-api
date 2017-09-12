package main 

import (
	"./queries"
	"./mutations"
    "./security"
    "./database"

	"net/http"
    "log"
    "database/sql"
    "fmt"

	"github.com/graphql-go/handler"
    "github.com/graphql-go/graphql"
    _ "github.com/lib/pq"
)

const (  
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "example"
    dbname   = "postgres"
)

func main() {

    var err error

    postgreSqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    database.DBCon, err = sql.Open("postgres", postgreSqlInfo)

	if err != nil {
        panic(err)
    }

    defer database.DBCon.Close()
    err = database.DBCon.Ping()

    if err != nil {
        panic(err)
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
	log.Printf("ready: listening...\n")

	http.ListenAndServe(":8383", nil)
}
