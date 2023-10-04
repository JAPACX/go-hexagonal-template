package graph

import (
	"go-gqlgen/application"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserUseCase *application.UserUseCase
}
