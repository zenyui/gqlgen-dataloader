package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zenyui/gqlgen-dataloader/graph/dataloader"
	"github.com/zenyui/gqlgen-dataloader/graph/generated"
	"github.com/zenyui/gqlgen-dataloader/graph/model"
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

func (r *queryResolver) ListTodos(ctx context.Context) ([]*model.Todo, error) {
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
	fmt.Printf("todoResolver.User, todo=%s, user=%s\n", obj.ID, obj.UserID)
	return dataloader.For(ctx).GetUser(obj.UserID)
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.db.GetAllTodos(ctx)
}
