package queries

import (
	"github.com/graphql-go/graphql"
)

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"user": GetUserQuery(),
	}
}

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:	"RootQuery",
	Fields: GetRootFields(),
});