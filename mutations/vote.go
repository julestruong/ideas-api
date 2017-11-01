package mutations

import (
    "../security"
    "../database"

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
            
            log.Printf("calling vote 1")
            var queryParams database.VoteQueryParams
            queryParams = database.VoteQueryParams{
                Id: params.Args["id"].(int),
                Email: security.User.Email,
            }
            log.Printf("calling vote")
            voteOk, err := database.Vote(queryParams);
            if err != nil {
                log.Printf("err %v", err)
            }
            log.Printf("user %s voted id %d", security.User.Email, params.Args["id"].(int))

            return voteOk, err
        },
    }
}
