package queries

import (
	"../types"

	"log"
	
	"github.com/graphql-go/graphql"
)

func GetIdeaQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.IdeaType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			log.Printf("[query] idea\n")
			var ideas []types.Idea

			return ideas, nil 
		},
	}
}
