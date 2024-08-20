// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package modelgen

import (
	"github.com/msoft-g1/todo-list-backend/internal/domain/task"
	"github.com/msoft-g1/todo-list-backend/internal/domain/user"
	"github.com/msoft-g1/todo-list-backend/internal/errs"
)

// Mutation operations
type Mutation struct {
}

// Query operations
type Query struct {
}

type TaskCreatePayload struct {
	Task  *task.Task  `json:"task,omitempty"`
	Error *errs.Error `json:"error,omitempty"`
}

type TaskDeletePayload struct {
	Error *errs.Error `json:"error,omitempty"`
}

type TaskUpdatePayload struct {
	Task  *task.Task  `json:"task,omitempty"`
	Error *errs.Error `json:"error,omitempty"`
}

type UserAccessTokenCreatePayload struct {
	User            *user.User  `json:"user,omitempty"`
	UserAccessToken *string     `json:"userAccessToken,omitempty"`
	Error           *errs.Error `json:"error,omitempty"`
}

type UserAccessTokenRenewPayload struct {
	UserAccessToken *string     `json:"userAccessToken,omitempty"`
	Error           *errs.Error `json:"error,omitempty"`
}

type UserCreatePayload struct {
	User  *user.User  `json:"user,omitempty"`
	Error *errs.Error `json:"error,omitempty"`
}
