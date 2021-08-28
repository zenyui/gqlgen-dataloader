package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	gopher_dataloader "github.com/graph-gophers/dataloader"
	"github.com/troopdev/graphql-poc/graph/dataloader"
	"github.com/troopdev/graphql-poc/graph/generated"
	"github.com/troopdev/graphql-poc/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:   input.UserID,
		Name: input.Name,
	}
	if user.ID == "" {
		user.ID = uuid.NewString()
	}
	// add to state
	if err := r.db.PutUser(ctx, user); err != nil {
		return nil, err
	}
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
	if err := r.db.PutTodo(ctx, todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.db.GetAllTodos(ctx)
}

func (r *queryResolver) GetTodo(ctx context.Context, id string) (*model.Todo, error) {
	results, err := r.db.GetTodos(ctx, []string{id})
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("todo not found %s", id)
	}
	return results[0], nil
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	results, err := r.db.GetUsers(ctx, []string{id})
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("user not found %s", id)
	}
	return results[0], nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	searchKey := obj.UserID
	fmt.Printf("calling user loader %s\n", searchKey)
	thunk := dataloader.For(ctx).UserById.Load(ctx, gopher_dataloader.StringKey(searchKey))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.User), nil
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
