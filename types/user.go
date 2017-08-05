package types

import (
	"github.com/graphql-go/graphql"
)

type User struct {
	ID		  int 	 `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	LastName  string `db:"lastname" json:"lastname"`
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":		 &graphql.Field{Type: graphql.Int},
		"firstname": &graphql.Field{Type: graphql.String},
		"lastname":	 &graphql.Field{Type: graphql.String},
		"ideas": 	 &graphql.Field{
			Type: graphql.NewList(IdeaType),
			Resolve: func (params graphql.ResolveParams) (interface{}, error) {
				var ideas []Idea

				return ideas, nil
			},
		},
	},
})