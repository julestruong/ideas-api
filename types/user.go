package types

import (
	"github.com/graphql-go/graphql"
)

type User struct {
	ID		  int 	 `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":		 &graphql.Field{Type: graphql.Int},
		"email":	 &graphql.Field{Type: graphql.String},
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
