package repositories

import (
	"context"
	"go-gqlgen/domain/entities"
)

type userRepository interface {
	Users(ctx context.Context) ([]entities.User, error)
	UserById(ctx context.Context) (entities.User, error)
	CreateUser(ctx context.Context, user entities.User) (string, error)
	UpdateUser(ctx context.Context, id string, user entities.User) error
	DeleteUser(ctx context.Context, id string) (string, error)
}
