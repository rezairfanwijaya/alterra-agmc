package database

import (
	"altera/Day2/config"
	"altera/Day2/models"
)

func GetUsers() (interface{}, error) {
	// init stuct user
	var users []models.User

	// orm to get users
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func AddUser(user models.User) (models.User, error) {
	if err := config.DB.Create(&user); err != nil {
		return user, err.Error
	}

	return user, nil
}
