package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"go-gqlgen/infrastructure/db/conn"
	graph2 "go-gqlgen/infrastructure/graph"
	"go-gqlgen/infrastructure/repository"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {

	ctx := context.Background()
	pool := conn.Connect(ctx)
	repository.InitRepository(pool)
	defer pool.Close()

	srv := handler.NewDefaultServer(graph2.NewExecutableSchema(graph2.Config{Resolvers: &graph2.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		panic("the PORT environment variable is empty or not found")
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
