package dataloader

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/graph-gophers/dataloader"
	"github.com/troopdev/graphql-poc/graph/storage"
)

const (
	loadersKey = "dataloaders"
)

type Loaders struct {
	UserById *dataloader.Loader
}

func NewLoaders(db storage.Storage) *Loaders {
	loaders := &Loaders{}
	// instantiate the user dataloader
	users := &userBatcher{db: db}
	loaders.UserById = dataloader.NewBatchedLoader(users.get)
	return loaders
}

func Middleware(db storage.Storage, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loaders := NewLoaders(db)
		ctx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

type userBatcher struct {
	db storage.Storage
}

// batchFn implements the dataloader for finding many users by Id
func (u *userBatcher) get(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	fmt.Printf("running dataloader batchFn, keys: %s\n", strings.Join(keys.Keys(), ","))
	// create a map for remembering the order of keys passed in
	keyOrder := make(map[string]int, len(keys))
	// collect the keys to search for
	var userIDs []string
	for ix, key := range keys {
		userIDs = append(userIDs, key.String())
		keyOrder[key.String()] = ix
	}
	// search for those users
	dbRecords, err := u.db.GetUsers(ctx, userIDs)
	// if DB error, return
	if err != nil {
		return []*dataloader.Result{{Data: nil, Error: err}}
	}
	// construct an output array of dataloader results
	results := make([]*dataloader.Result, len(keys))
	// enumerate records, put into output
	for _, record := range dbRecords {
		ix, ok := keyOrder[record.ID]
		// if found, remove from index lookup map so we know elements were found
		if ok {
			results[ix] = &dataloader.Result{Data: record, Error: nil}
			delete(keyOrder, record.ID)
		}
	}
	// fill array positions with errors where not found in DB
	for userID, ix := range keyOrder {
		err := errors.New(fmt.Sprintf("user not found %s", userID))
		results[ix] = &dataloader.Result{Data: nil, Error: err}
	}
	// return results
	return results
}
