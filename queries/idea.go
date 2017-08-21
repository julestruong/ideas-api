package queries

import (
	"strconv"

	"../database"
	"../types"

	"log"

	"github.com/graphql-go/graphql"
)

/**
* GetIdeaQuery
*
 */
func GetIdeaQuery() *graphql.Field {
	return &graphql.Field{
		Type: types.IdeaType,
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

			i := params.Args["id"].(string)
			email := params.Args["email"].(string)

			id, err := strconv.Atoi(i)
			if err != nil {
				return nil, err
			}

			ideas := database.Select(database.Params{Id: id, Email: email})

			return ideas, nil
		},
	}
}
