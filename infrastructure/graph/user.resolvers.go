package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"go-gqlgen/domain/entities"
	"go-gqlgen/infrastructure/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	userToCreate := entities.User{
		Name:  input.Name,
		Email: input.Email,
	}

	id, err := r.UserUseCase.CreateUser(ctx, userToCreate)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    id,
		Name:  userToCreate.Name,
		Email: userToCreate.Email,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *model.UpdateUser) (*model.User, error) {
	userToUpdate := entities.User{
		Name:  *input.Name,
		Email: *input.Email,
	}

	if err := r.UserUseCase.UpdateUser(ctx, id, userToUpdate); err != nil {
		return nil, err
	}

	return &model.User{
		ID:    id,
		Name:  userToUpdate.Name,
		Email: userToUpdate.Email,
	}, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	_, err := r.UserUseCase.DeleteUser(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	entityUsers, err := r.UserUseCase.GetUsers(ctx)

	if err != nil {
		return nil, err
	}

	var modelUsers []*model.User
	for _, eu := range *entityUsers {
		mu := &model.User{
			ID:    eu.Id,
			Name:  eu.Name,
			Email: eu.Email,
		}
		modelUsers = append(modelUsers, mu)
	}
	return modelUsers, nil
}

// UserByID is the resolver for the userById field.
func (r *queryResolver) UserByID(ctx context.Context, id string) (*model.User, error) {

	userEntity, err := r.UserUseCase.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	mu := &model.User{
		ID:    userEntity.Id,
		Name:  userEntity.Name,
		Email: userEntity.Email,
	}

	return mu, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
