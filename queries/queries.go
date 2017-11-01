package queries

import (
	"github.com/graphql-go/graphql"
)

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"ideas": GetIdeasQuery(),
	}
}

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:	"RootQuery",
	Fields: GetRootFields(),
});
