package repositories

import (
	"context"
	"go-gqlgen/domain/entities"
)

type UserInterface interface {
	Users(ctx context.Context) (*[]entities.User, error)
	UserById(ctx context.Context, id string) (*entities.User, error)
	CreateUser(ctx context.Context, user entities.User) (*entities.User, error)
	UpdateUser(ctx context.Context, id string, user entities.User) (*entities.User, error)
	DeleteUser(ctx context.Context, id string) (bool, error)
}
