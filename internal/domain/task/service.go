package task

// Service defines task service.
type Service struct {
	repo Repository
}

// NewService creates a new task service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateTask creates a new task.
func (s *Service) CreateTask(input *CreateInput) (*Task, error) {
	defCompleted := false
	return s.repo.Create(&Task{
		Description: input.Description,
		Completed:   &defCompleted,
		UserID:      input.UserID,
	})
}

// UpdateTask updates the task with the specified id.
func (s *Service) UpdateTask(id uint, input *UpdateInput) (*Task, error) {
	return s.repo.Update(&Task{
		ID:          id,
		Description: input.Description,
		Completed:   input.Completed,
	})
}

// DeleteTask deletes the task with the specified id.
func (s *Service) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}

// GetTaskByID returns the task with the specified id.
func (s *Service) GetTaskByID(id uint) (*Task, error) {
	return s.repo.FindByID(id)
}

// GetAllTasks returns all the tasks.
func (s *Service) GetAllTasks() ([]*Task, error) {
	return s.repo.FindAll()
}

// GetTasksByUser returns all the tasks owned by the user with the specified id.
func (s *Service) GetTasksByUser(userID uint) ([]*Task, error) {
	return s.repo.FindAllByUser(userID)
}
