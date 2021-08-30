# GqlGen graphql server with dataloader

This repo demonstrates a simple [GqlGen](https://gqlgen.com/) graphql server using the [graph-gophers/dataloader](https://github.com/graph-gophers/dataloader). This service is based on the GqlGen "Todo" sample app, and the dataloader middleware is inspired by [vektah/dataloaden](https://github.com/vektah/dataloaden). Unlike vekta/dataloaden, this dataloader does not rely on code generation.

### Project Structure
```sh
.
├── gqlgen.yml # GqlGen configuration
├── graph
│   ├── dataloader # implements the user dataloader
│   ├── generated # GqlGen generated files
│   ├── model # defines User and Todo model structs
│   ├── resolver # implements the graphql queries/mutations
│   └── storage # mock datastore for use in dataloader
├── schema.graphqls # graphql schema definition
└── server.go # runnable server
```


### Run the app
```sh
# run the app, then navigate to http://localhost:8080/
go run ./server.go
```

Once the app server is running you can create the user...
```graphql
mutation createUser {
  createUser(input: {name: "foo", userId:"1"}) {
    id
  }
}
```

Create a todo...
```graphql
mutation createTodo {
  createTodo(input:{
    text:"bar",
    userId:"1"
  }) {
    id
  }
}
```

And observe the data loader in action!
```graphql
query {
  listTodos {
    id
    text
    user {
      name
    }
  }
}
```

### Building the app

```sh
# run GqlGen (to generate new resolvers & models)
go get github.com/99designs/gqlgen
go generate ./...
```

*NOTE:* there are [issues](https://github.com/99designs/gqlgen/issues/800) running the GqlGen generator if you have vendored dependencies. The easiest workaround is to delete the `vendor/` folder when you generate.
