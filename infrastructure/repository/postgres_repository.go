package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-gqlgen/domain/entities"
)

type Repository struct {
	Pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{Pool: pool}
}

func (r *Repository) Users(ctx context.Context) (*[]entities.User, error) {

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
	return &users, nil
}

func (r *Repository) UserById(ctx context.Context, id string) (*entities.User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	var user entities.User
	err := r.Pool.QueryRow(ctx, query, id).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return &user, fmt.Errorf("error fetching user by ID: %v", err)
	}
	return &user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user entities.User) (*entities.User, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email`
	var newUser entities.User

	err := r.Pool.QueryRow(ctx, query, user.Name, user.Email).Scan(&newUser.Id, &newUser.Name, &newUser.Email)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return &newUser, nil
}

func (r *Repository) UpdateUser(ctx context.Context, id string, user entities.User) (*entities.User, error) {

	// Update the user data and fetch the updated data in the same query
	query := `
		UPDATE users 
		SET name = $1, email = $2 
		WHERE id = $3 
		RETURNING id, name, email
	`
	row := r.Pool.QueryRow(ctx, query, user.Name, user.Email, id)

	var updatedUser entities.User
	err := row.Scan(&updatedUser.Id, &updatedUser.Name, &updatedUser.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("error updating and fetching user: %v", err)
	}

	return &updatedUser, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id string) (bool, error) {

	query := `DELETE FROM users WHERE id = $1`
	tag, err := r.Pool.Exec(ctx, query, id)
	if err != nil {
		return false, fmt.Errorf("error deleting user: %v", err)
	}
	if tag.RowsAffected() == 0 {
		return false, errors.New("user not found")
	}
	return true, nil
}
