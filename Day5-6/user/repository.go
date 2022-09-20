package user

import (
	"gorm.io/gorm"
)

type Respository interface {
	ShowAllUser() ([]User, error)
	ShowUserById(userID int) (User, error)
	DeleteUserById(userID int) error
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

// function get user by id
func (r *repository) ShowUserById(userID int) (User, error) {
	var user User
	if err := r.db.First(&user, userID).Error; err != nil {
		return user, err
	}

	return user, nil
}

// function delete user by id
func (r *repository) DeleteUserById(userID int) error {
	var user User
	if err := r.db.First(&user, userID).Error; err != nil {
		return err
	}

	if err := r.db.Where("id = ?", userID).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
