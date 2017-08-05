package queries

import (
	"../types"

	"log"
	
	"github.com/graphql-go/graphql"
)

func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			log.Printf("[query] user\n")
			var users []types.User

			return users, nil 
		},
	}
}