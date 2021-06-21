package lists

import "server/entity"

type Service interface {
	GetAllLists() ([]entity.List, error)
	GetListByIDService (listID string) ([]entity.List, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllLists() ([]entity.List, error) {
	lists, err := s.repository.GetAllLists()

	if err != nil {
		return lists, nil
	}

	return lists, nil
}

func (s *service) GetListByIDService (listID string) ([]entity.List, error) {
	lists, err := s.repository.GetListssByID(listID)

	if err != nil {
		return lists, nil
	}

	return lists, nil
}


