package sqliterepo

import (
	"errors"

	"github.com/msoft-g1/todo-list-backend/internal/domain/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ user.Repository = (*usersRepo)(nil)

// usersRepo sqlite users repository.
type usersRepo struct {
	db *gorm.DB
}

// NewUsersRepo creates a new users repository.
func NewUsersRepo(db *gorm.DB) user.Repository {
	db.AutoMigrate(&user.User{})
	return &usersRepo{db}
}

func (r *usersRepo) Create(u *user.User) (*user.User, error) {
	ret := *u
	err := r.db.Create(&ret).Error
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (r *usersRepo) FindByID(id uint) (*user.User, error) {
	var u user.User
	err := r.db.First(&u, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *usersRepo) FindbyEmail(email string) (*user.User, error) {
	var u user.User
	err := r.db.First(&u, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *usersRepo) FindAll() ([]*user.User, error) {
	var users []*user.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *usersRepo) Update(u *user.User) (*user.User, error) {
	ret := user.User{ID: u.ID}
	err := r.db.Model(&ret).Clauses(clause.Returning{}).Updates(u).Error
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (r *usersRepo) Delete(id uint) error {
	return r.db.Delete(&user.User{}, id).Error
}
