package application

import (
	"context"
	"errors"
	"go-gqlgen/domain/entities"
	"go-gqlgen/domain/repositories"
)

type UserUseCase struct {
	repo repositories.UserInterface
}

func NewUserUseCase(repo repositories.UserInterface) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (ui *UserUseCase) GetUsers(ctx context.Context) (*[]entities.User, error) {
	return ui.repo.Users(ctx)
}

func (ui *UserUseCase) GetUserById(ctx context.Context, id string) (*entities.User, error) {
	return ui.repo.UserById(ctx, id)
}

func (ui *UserUseCase) CreateUser(ctx context.Context, user entities.User) (*entities.User, error) {

	return ui.repo.CreateUser(ctx, user)
}

func (ui *UserUseCase) UpdateUser(ctx context.Context, id string, user entities.User) (*entities.User, error) {

	return ui.repo.UpdateUser(ctx, id, user)
}

func (ui *UserUseCase) DeleteUser(ctx context.Context, id string) (bool, error) {

	if id == "" {
		return false, errors.New("missing user ID for delete")
	}

	return ui.repo.DeleteUser(ctx, id)
}
