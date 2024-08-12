package task

// CreateInput defines input to create a task.
type CreateInput struct {
	Description string
	UserID      uint
}

// UpdateInput defines input to update a task.
type UpdateInput struct {
	Description string
	Completed   *bool
}
