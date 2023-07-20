package storage

import (
	"context"
	"fmt"

	"github.com/zenyui/gqlgen-dataloader/graph/model"
)

// Storage provides methods for reading/writing Users and Todos to a datastore
type Storage interface {
	// PutUser persists a User
	PutUser(ctx context.Context, usr *model.User) error
	// PutTodo persists a Todo
	PutTodo(ctx context.Context, todo *model.Todo) error
	// GetUsers accepts many user IDs and returns an array of matching Users
	GetUsers(ctx context.Context, ids []string) ([]*model.User, error)
	// GetTodos accepts many Todo IDs and returns an array of matching Todos
	GetTodos(ctx context.Context, ids []string) ([]*model.Todo, error)
	// GetAllTodos lists all Todo's in the database
	GetAllTodos(ctx context.Context) ([]*model.Todo, error)
}

// MemoryStorage implements the Storage methods in memory as golang maps
type MemoryStorage struct {
	todos map[string]*model.Todo
	users map[string]*model.User
}

// NewMemoryStorage returns a MemoryStorage with internal maps initialized
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		todos: make(map[string]*model.Todo),
		users: make(map[string]*model.User),
	}
}

func (m *MemoryStorage) PutUser(ctx context.Context, usr *model.User) error {
	m.users[usr.ID] = usr
	return nil
}

func (m *MemoryStorage) PutTodo(ctx context.Context, todo *model.Todo) error {
	m.todos[todo.ID] = todo
	return nil
}

func (m *MemoryStorage) GetUsers(ctx context.Context, ids []string) ([]*model.User, error) {
	fmt.Printf("GetUsers %v\n", ids)
	output := make([]*model.User, 0, len(ids))
	for _, id := range ids {
		if usr, ok := m.users[id]; ok {
			output = append(output, usr)
		}
	}
	return output, nil
}

func (m *MemoryStorage) GetTodos(ctx context.Context, ids []string) ([]*model.Todo, error) {
	output := make([]*model.Todo, 0, len(ids))
	for _, id := range ids {
		if todo, ok := m.todos[id]; ok {
			output = append(output, todo)
		}
	}
	return output, nil
}

func (m *MemoryStorage) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
	output := make([]*model.Todo, 0, len(m.todos))
	for _, todo := range m.todos {
		output = append(output, todo)
	}
	return output, nil
}
