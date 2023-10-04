package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"go-gqlgen/application"
	graph "go-gqlgen/infrastructure/api/graph/generated"
	resolvers2 "go-gqlgen/infrastructure/api/graph/resolvers"
	"go-gqlgen/infrastructure/db/connect"
	"go-gqlgen/infrastructure/repository"
	"log"
	"net/http"
	"os"
)

func init() {
	fmt.Println("first this ")
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	ctx := context.Background()
	pool := connect.Connect(ctx)
	defer pool.Close()

	// Create an instance of the repository using the connection pool.
	repo := &repository.Repository{Pool: pool}

	// Instantiate a user use case by providing the previously created repository instance.
	// Thanks to Golang's implicit interfaces, we can pass the new repo which implements the required interface.
	userUseCase := application.NewUserUseCase(repo)

	// Initialize resolvers with the use case.
	resolvers := &resolvers2.Resolver{UserUseCase: userUseCase}

	// Set up the resolvers for GraphQL.
	// If additional resolvers are needed, instantiate more repository and use case instances, and initialize a new resolver.
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolvers}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		panic("the PORT environment variable is empty or not found")
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
