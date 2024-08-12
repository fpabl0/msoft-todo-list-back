package task

// Task defines a single Task in a Todo-List
type Task struct {
	ID          uint
	Description string
	Completed   *bool
	UserID      uint
}
