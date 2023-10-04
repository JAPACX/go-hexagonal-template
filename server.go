package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"go-gqlgen/application"
	resolvers2 "go-gqlgen/infrastructure/api/graph/resolvers"
	"go-gqlgen/infrastructure/db/connect"
	graph "go-gqlgen/infrastructure/graph"
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

	// Crear instancia del repositorio con el pool de conexión
	repo := &repository.Repository{Pool: pool}

	// Crear instancia del caso de uso y pasar el repositorio
	userUseCase := application.NewUserUseCase(repo) // Asume que tienes un caso de uso llamado "UserUseCase" en tu paquete application

	// Inicializar resolvers con el caso de uso
	resolvers := &resolvers2.Resolver{UserUseCase: userUseCase}

	// Ahora, usa estos resolvers para tu servidor GraphQL
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
