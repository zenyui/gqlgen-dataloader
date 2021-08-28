package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/troopdev/graphql-poc/graph/dataloader"
	"github.com/troopdev/graphql-poc/graph/generated"
	"github.com/troopdev/graphql-poc/graph/resolver"
	"github.com/troopdev/graphql-poc/graph/storage"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// instantiate the DB client
	db := storage.NewMemoryStorage()
	// instantiate the gqlgen Graph Resolver
	graphResolver := resolver.NewResolver(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graphResolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloader.Middleware(db, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
