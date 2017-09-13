package queries

import (
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
			"week": &graphql.ArgumentConfig{
				Description: "Idea Week",
				Type:        graphql.String,
			},
			"email": &graphql.ArgumentConfig{
				Description: "Idea email user",
				Type:        graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
            log.Printf("[query] idea\n")
            
            var ideas []types.Idea
            
            var queryParams database.QueryParams
            if params.Args["week"] != nil {
                queryParams.Week = params.Args["week"].(string) 
            }
            
            if params.Args["email"] != nil {
                queryParams.Email = params.Args["email"].(string) 
            }

			ideas = database.Select(queryParams)

            log.Printf("%v", ideas)
			return ideas, nil
		},
	}
}
