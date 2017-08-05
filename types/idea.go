package types

import (
	"github.com/graphql-go/graphql"
)

type Idea struct {
	ID 		int 	`db:"id" json:"id"`
	Body	string	`db:"body" json:"body"`
}

var IdeaType = graphql.NewObject(graphql.ObjectConfig{
	Name: 	"Idea",
	Fields: graphql.Fields{
		"id": 	&graphql.Field{Type: graphql.Int},
		"body": &graphql.Field{Type: graphql.String},
	},
})