package queries

import (
	"strconv"

	"../database"
	"../types"

	"log"

	"github.com/graphql-go/graphql"
)

/**
* GetIdeasQuery
*
 */
func GetIdeasQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.IdeaType),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Idea ID",
				Type:        graphql.ID,
			},
			"email": &graphql.ArgumentConfig{
				Description: "Idea email user",
				Type:        graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
            log.Printf("[query] idea\n")
            
            var ideas []types.Idea

			i := params.Args["id"].(string)
			email := params.Args["email"].(string)

			id, err := strconv.Atoi(i)
			if err != nil {
				return "", err
			}

			ideas = database.Select(database.Params{Id: id, Email: email})

            log.Printf("%v", ideas)
			return ideas, nil
		},
	}
}
