package mutations

import (
	"github.com/graphql-go/graphql"
)

func GetRootFields() graphql.Fields {
	return graphql.Fields{
        "createUser": GetCreateUserMutation(),
        "deleteUser": GetDeleteUserMutation(),
        "updateUser": GetUpdateUserMutation(),
	}
}

var MutationType = graphql.NewObject(graphql.ObjectConfig{
    Name:	"RootMutation",
    Fields: GetRootFields(),
});
