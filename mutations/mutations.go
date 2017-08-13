package mutations

import (
	"github.com/graphql-go/graphql"
)

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"createUser": GetCreateUserMutation(),
	}
}

var MutationType = graphql.NewObject(graphql.ObjectConfig{
    Name:	"RootMutation",
    Fields: GetRootFields(),
});
