package mutations

import (
    "../security"

    "log"

    "github.com/graphql-go/graphql"
)

func GetVoteIdeaMutation() *graphql.Field {
    return &graphql.Field{
        Type: graphql.Boolean,
        Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{
                Type: graphql.NewNonNull(graphql.Int),
            },
        },
        Resolve: func(params graphql.ResolveParams) (interface{}, error) {
            log.Printf("user %s voted id %d", security.User.Email, params.Args["id"].(int))

            return true, nil
        },
    }
}
