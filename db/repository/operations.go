package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Repository struct {
	Pool *pgxpool.Pool
}

type User struct {
	Id        string
	Name      string
	Email     string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (r *Repository) CreateUser(ctx context.Context, u User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2)`

	_, err := r.Pool.Exec(ctx, query, u.Name, u.Email)
	if err != nil {
		return fmt.Errorf("Error creating user: %v", err)
	}
	fmt.Println("User created successfully!")
	return nil
}
