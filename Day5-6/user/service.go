package user

import "errors"

type Service interface {
	GetAllUser() ([]User, error)
	GetUserById(userID int) (User, error)
	DeleteUserById(userID int) error
	AddNewUser(userInput UserInput) (User, error)
	UpdateUserById(userInput UserInput, userID int) (User, error)
}

type service struct {
	repo Respository
}

func NewService(repo Respository) *service {
	return &service{repo}
}

func (s *service) GetAllUser() ([]User, error) {
	users, err := s.repo.ShowAllUser()
	if err != nil {
		return users, errors.New("error while get data all users")
	}

	return users, nil
}

func (s *service) GetUserById(userID int) (User, error) {
	user, err := s.repo.ShowUserById(userID)
	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (s *service) DeleteUserById(userID int) error {
	err := s.repo.DeleteUserById(userID)
	if err != nil {
		return errors.New("user not found")
	}

	return nil
}

func (s *service) AddNewUser(userInput UserInput) (User, error) {
	// cek apakah user tersebut sudah terdaftar di db
	existingUser, err := s.repo.ShowUserByEmail(userInput.Email)
	if err != nil {
		return User{}, err
	}

	if existingUser.Id != 0 {
		return User{}, errors.New("user is already registered")
	}

	// add user to db
	var user User
	user.Email = userInput.Email
	user.Password = userInput.Password
	user.Name = userInput.Name

	newUser, err := s.repo.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) UpdateUserById(userInput UserInput, userID int) (User, error) {

	// cek apakah user ada
	userExisting, err := s.repo.ShowUserById(userID)
	if err != nil {
		return userExisting, errors.New("user not found")
	}

	// cek apakah email yang diupdate sudah ada, jika ada maka tolak update
	emailExisting, err := s.repo.ShowUserByEmail(userInput.Email)
	if err != nil {
		return emailExisting, err
	}

	if emailExisting.Id != 0 {
		return emailExisting, errors.New("email alredy taken")
	}

	// update
	userExisting.Email = userInput.Email
	userExisting.Password = userInput.Password
	userExisting.Name = userInput.Name

	// save
	userUpdated, err := s.repo.Save(userExisting)
	if err != nil {
		return userUpdated, errors.New("failed to update user")
	}

	return userUpdated, nil
}
