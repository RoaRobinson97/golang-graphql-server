package graph

import "github.com/roarobinson97/graphql-server/graph/model"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	videos []*model.Video
}
