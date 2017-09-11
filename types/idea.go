package types

import (
    "time"

	"github.com/graphql-go/graphql"
)

type Idea struct {
	ID 		    int 	    `db:"id" json:"id"`
	Email	    string	    `db:"email" json:"email"`
	Body	    string	    `db:"body" json:"body"`
	CreatedAt	time.Time 	`db:"created_at" json:"created_at"`
}

var IdeaType = graphql.NewObject(graphql.ObjectConfig{
	Name: 	"Idea",
	Fields: graphql.Fields{
        "id": 	        &graphql.Field{Type: graphql.Int},
        "email":        &graphql.Field{Type: graphql.String},
        "body":         &graphql.Field{Type: graphql.String},
        "created_at":   &graphql.Field{Type: graphql.DateTime},
	},
})
