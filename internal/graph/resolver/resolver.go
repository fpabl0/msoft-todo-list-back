package resolver

import (
	"github.com/msoft-g1/todo-list-backend/internal/domain/task"
	"github.com/msoft-g1/todo-list-backend/internal/domain/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver defines resolver object.
type Resolver struct {
	UsersService *user.Service
	TasksService *task.Service
}
