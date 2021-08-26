package resolver

//go:generate go run github.com/99designs/gqlgen generate

import "github.com/troopdev/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	users []*model.User
}
