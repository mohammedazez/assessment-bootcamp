package users

import "server/entity"

type Service interface {
	GetAllUser() ([]entity.User, error)
	GetUserByIDService (userID string) ([]entity.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUser() ([]entity.User, error) {
	users, err := s.repository.GetAllUsers()

	if err != nil {
		return users, nil
	}

	return users, nil
}

func (s *service) GetUserByIDService (userID string) ([]entity.User, error) {
	users, err := s.repository.GetUsersByID(userID)

	if err != nil {
		return users, nil
	}

	return users, nil
}


