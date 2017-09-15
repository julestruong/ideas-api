package mutations

import (
	"github.com/graphql-go/graphql"
)

func GetRootFields() graphql.Fields {
	return graphql.Fields{
        "createIdea": GetCreateIdeaMutation(),
        "updateIdea": GetUpdateIdeaMutation(),
        "voteIdea":   GetVoteIdeaMutation(),
	}
}

var MutationType = graphql.NewObject(graphql.ObjectConfig{
    Name:	"RootMutation",
    Fields: GetRootFields(),
});
