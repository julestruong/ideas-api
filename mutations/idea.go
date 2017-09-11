package mutations

import (
	"../types"
    "../database"

    "log"

	"github.com/graphql-go/graphql"
)

func GetCreateIdeaMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.IdeaType,
		Args: graphql.FieldConfigArgument{
			"body": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
            },
            "email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			idea := &types.Idea{
				Body: params.Args["firstname"].(string),
				Email: params.Args["email"].(string),
			}

            database.InsertIdea(idea);

			log.Printf("idea has been created");

			return idea, nil
		},
    }
}

func GetUpdateIdeaMutation() *graphql.Field {
    return &graphql.Field{
        Type: types.IdeaType,
        Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"body": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
        },
        Resolve: func(params graphql.ResolveParams) (interface{}, error) {

            id := params.Args["id"].(int)
            body := params.Args["body"].(string)

            count := database.UpdateIdea(id, body)

            return count, nil
        },
    }
}
