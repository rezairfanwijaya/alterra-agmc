package database

import (
	"altera/Day2/config"
	"altera/Day2/models"
)

// get all user
func GetUsers() (interface{}, error) {
	// init stuct user
	var users []models.User

	// orm to get users
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// add new user
func AddUser(user models.User) (models.User, error) {
	if err := config.DB.Create(&user); err != nil {
		return user, err.Error
	}

	return user, nil
}

// find user by email
func FindUserByEmail(email string) (models.User, error) {
	var user models.User

	if err := config.DB.Where("email = ?", email).Find(&user); err != nil {
		return user, err.Error
	}

	return user, nil
}

// find user by id
func FindUserById(id int) (models.User, error) {
	var user models.User

	if err := config.DB.Where("id = ?", id).Find(&user); err != nil {
		return user, err.Error
	}

	return user, nil
}

// update user by id
func UpdateUserById(user models.User) (models.User, error) {
	if err := config.DB.Save(&user); err != nil {
		return user, err.Error
	}

	return user, nil
}

// delete user by id
func DeleteUserById(id int) error {
	var user models.User
	if err := config.DB.Where("id = ?", id).Delete(&user); err != nil {
		return err.Error
	}

	return nil
}
