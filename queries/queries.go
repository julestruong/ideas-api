package queries

import (
	"github.com/graphql-go/graphql"
)

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"idea": GetIdeaQuery(),
	}
}

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:	"RootQuery",
	Fields: GetRootFields(),
});
