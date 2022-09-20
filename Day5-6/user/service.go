package user

import "errors"

type Service interface {
	GetAllUser() ([]User, error)
	GetUserById(userID int) (User, error)
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
