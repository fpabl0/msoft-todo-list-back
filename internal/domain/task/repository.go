package task

// Repository defines task repository methods.
type Repository interface {
	Create(t *Task) (*Task, error)
	// if not found, should return (nil, nil)
	FindByID(id uint) (*Task, error)
	FindAll() ([]*Task, error)
	FindAllByUser(userID uint) ([]*Task, error)
	Update(t *Task) (*Task, error)
	Delete(id uint) error
}
