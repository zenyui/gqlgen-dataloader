# Graphql POC

This repo demonstrates a simple graphql server making calls to donwstream gRPC services using the [GqlGen](https://gqlgen.com/) library.

### Quickstart
TODO

*NOTE:* there are [issues](https://github.com/99designs/gqlgen/issues/800) running the GqlGen generator if you have vendored dependencies. The easiest workaround is to delete the `vendor/` folder when you generate.

### Concepts

*Dataloader*

From the authors:
> DataLoader is a generic utility to be used as part of your application's data fetching layer to provide a consistent API over various backends and reduce requests to those backends via batching and caching.

Dataloaders are used in graphql to reduce the number of round-trips to a given data store. For example, consider you are modeling graph with a `Team` object that contains many `User` objects and a `listTeams` query that returns an array of teams and their nested users. Instead of fetching each user from the database individually, a dataloader could implement bulk-fetching of `User` records by key, and the graphql server can even cache those user records in case they appear in many teams.
