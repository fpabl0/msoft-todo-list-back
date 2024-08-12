package user

// Repository defines user repository methods.
type Repository interface {
	Create(u *User) (*User, error)
	// if not found, should return (nil, nil)
	FindByID(id uint) (*User, error)
	// if not found, should return (nil, nil)
	FindbyEmail(email string) (*User, error)
	FindAll() ([]*User, error)
	Update(u *User) (*User, error)
	Delete(id uint) error
}
