package storage

import (
	"context"
	"fmt"

	"github.com/troopdev/graphql-poc/graph/model"
)

type Storage interface {
	PutUser(ctx context.Context, usr *model.User) error
	PutTodo(ctx context.Context, todo *model.Todo) error
	GetUsers(ctx context.Context, ids []string) ([]*model.User, error)
	GetTodos(ctx context.Context, ids []string) ([]*model.Todo, error)
	GetAllTodos(ctx context.Context) ([]*model.Todo, error)
}

type MemoryStorage struct {
	todos map[string]*model.Todo
	users map[string]*model.User
}

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
	fmt.Println("calling MemoryStorage.GetUsers")
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
