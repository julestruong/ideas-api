package mutations

import (
	"../types"
    "../database"
    "../security"

    "log"
    "time"
    "strconv"

	"github.com/graphql-go/graphql"
)

func GetCreateIdeaMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.IdeaType,
		Args: graphql.FieldConfigArgument{
			"body": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
            },
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
            now := time.Now()
            year, week := now.ISOWeek()
            
			idea := &types.Idea{
                Body: params.Args["body"].(string),
                Week: strconv.Itoa(year) + strconv.Itoa(week),
                Email: security.User.Email, 
			}

            err := database.InsertIdea(idea);

            if err != nil {
                log.Printf("Error while trying to create an idea")
                return "", err
            }

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
