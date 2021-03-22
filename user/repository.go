package user

import (
	"gorm.io/gorm"
)

// Repository User with type interface
type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository to instance a new variable from struct repository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
