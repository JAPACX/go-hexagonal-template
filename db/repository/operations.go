package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"

	"go-gqlgen/db/entities"
)

type Repository struct {
	Pool *pgxpool.Pool
}

func (r *Repository) CreateUser(ctx context.Context, user entities.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2)`

	_, err := r.Pool.Exec(ctx, query, user.Name, user.Email)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	fmt.Println("User created successfully!")
	return nil
}
