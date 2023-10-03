package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-gqlgen/domain/entities"
)

type Repository struct {
	Pool *pgxpool.Pool
}

var instance *Repository

func InitRepository(pool *pgxpool.Pool) {
	instance = &Repository{Pool: pool}
}

func GetRepository() *Repository {
	return instance
}

func (r *Repository) Users(ctx context.Context) ([]entities.User, error) {

	query := `select * from users`
	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %v", err)
	}
	defer rows.Close()
	var users []entities.User

	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning user: %v", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}
	return users, nil
}

func (r *Repository) UserById(ctx context.Context, id string) (entities.User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	var user entities.User
	err := r.Pool.QueryRow(ctx, query, id).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, fmt.Errorf("error fetching user by ID: %v", err)
	}
	return user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user entities.User) (string, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	var id string
	err := r.Pool.QueryRow(ctx, query, user.Name, user.Email).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error creating user: %v", err)
	}
	return id, nil
}

func (r *Repository) UpdateUser(ctx context.Context, id string, user entities.User) error {
	if id == "" {
		return errors.New("missing user ID for update")
	}

	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	tag, err := r.Pool.Exec(ctx, query, user.Name, user.Email, id)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	if tag.RowsAffected() == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, id string) (string, error) {
	if id == "" {
		return "", errors.New("missing user ID for delete")
	}

	query := `DELETE FROM users WHERE id = $1`
	tag, err := r.Pool.Exec(ctx, query, id)
	if err != nil {
		return "", fmt.Errorf("error deleting user: %v", err)
	}
	if tag.RowsAffected() == 0 {
		return "", errors.New("user not found")
	}
	return id, nil
}
