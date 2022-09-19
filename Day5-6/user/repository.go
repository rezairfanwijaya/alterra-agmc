package user

import (
	"gorm.io/gorm"
)

type Respository interface {
	ShowAllUser() ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// funtion get all user
func (r *repository) ShowAllUser() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
