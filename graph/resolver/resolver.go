package resolver

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/zenyui/gqlgen-dataloader/graph/storage"
)

// Resolver serves as dependency injection for your app

type Resolver struct {
	// db is an interface for reading/writing to the datastore
	db storage.Storage
}

// NewResolver returns a Resolver
func NewResolver(db storage.Storage) *Resolver {
	output := &Resolver{db: db}
	return output
}
