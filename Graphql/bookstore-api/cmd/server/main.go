package main

import (
    "bookstore-api/graph/generated"
    "bookstore-api/graph/resolvers"
    "bookstore-api/internal/db"
    "log"
    "net/http"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
)

func main() {
    // Connect to Postgres using your existing internal/db code
    if err := db.Connect(); err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }

    resolver := &resolvers.Resolver{}
    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
        Resolvers: resolver,
    }))

    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    log.Println("connect to http://localhost:8080/ for GraphQL playground")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
