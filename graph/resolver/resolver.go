package resolver

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/troopdev/graphql-poc/graph/storage"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db storage.Storage
}

func NewResolver(db storage.Storage) *Resolver {
	output := &Resolver{db: db}
	return output
}
