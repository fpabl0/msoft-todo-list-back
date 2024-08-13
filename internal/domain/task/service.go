package task

import "github.com/msoft-g1/todo-list-backend/internal/errs"

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
	completed := false
	return s.repo.Create(&Task{
		Description: input.Description,
		Completed:   &completed,
		UserID:      input.UserID,
	})
}

// UpdateTask updates the task with the specified id. The userID is the user who is requesting
// the update, if that userID does not match with userID in the task, an error is returned.
func (s *Service) UpdateTask(id uint, userID uint, input *UpdateInput) (*Task, error) {
	t, err := s.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	if t.UserID != userID {
		return nil, errs.New(errs.CodeAccessDenied, "Solo el usuario dueño de esta tarea puede editarla")
	}
	return s.repo.Update(&Task{
		ID:          id,
		Description: input.Description,
		Completed:   input.Completed,
	})
}

// DeleteTask deletes the task with the specified id. The userID is the user who is requesting
// the delete, if that userID does not match with userID in the task, an error is returned.
func (s *Service) DeleteTask(id uint, userID uint) error {
	t, err := s.GetTaskByID(id)
	if err != nil {
		return err
	}
	if t.UserID != userID {
		return errs.New(errs.CodeAccessDenied, "Solo el usuario dueño de esta tarea puede borrarla")
	}
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
