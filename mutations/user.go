package mutations

import (
	"../types"
    "../database"

    "log"

	"github.com/graphql-go/graphql"
)

func GetCreateUserMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"firstname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lastname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
            },
            "email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			user := &types.User{
				Firstname: params.Args["firstname"].(string),
				Lastname: params.Args["lastname"].(string),
				Email: params.Args["email"].(string),
			}

            database.InsertUser(user);

			log.Printf("user has been created");

			return user, nil
		},
    }
}

func GetDeleteUserMutation() *graphql.Field {
    return &graphql.Field{
        Type: types.UserType,
        Args: graphql.FieldConfigArgument{
            "email": &graphql.ArgumentConfig{
                Type: graphql.NewNonNull(graphql.String),
            },
        },
        Resolve: func(params graphql.ResolveParams) (interface{}, error) {
            
            email := params.Args["email"].(string)

            log.Printf(email)
            count := database.DeleteUser(email)

            return count, nil
        },
    }
}


func GetUpdateUserMutation() *graphql.Field {
    return &graphql.Field{
        Type: types.UserType,
        Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"firstname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lastname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
            },
        },
        Resolve: func(params graphql.ResolveParams) (interface{}, error) {

            id := params.Args["id"].(int)
            firstname := params.Args["firstname"].(string)
            lastname := params.Args["lastname"].(string)

            count := database.UpdateUser(id, firstname, lastname)

            return count, nil
        },
    }
}
