package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/troopdev/graphql-poc/graph/generated"
	"github.com/troopdev/graphql-poc/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:   uuid.NewString(),
		Name: input.Name,
	}
	// add to state
	r.users = append(r.users, user)
	// return
	return user, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	// generate ID
	newID := uuid.NewString()
	// model item
	todo := &model.Todo{
		Text:   input.Text,
		ID:     newID,
		UserID: input.UserID,
	}
	// add to state
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) GetTodo(ctx context.Context, id string) (*model.Todo, error) {
	for _, t := range r.todos {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, errors.New("not found!")
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	fmt.Printf("user %s\n", obj.UserID)
	// find the user, else error
	for _, u := range r.users {
		if u.ID == obj.UserID {
			return u, nil
		}
	}
	return nil, errors.New("not found")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
