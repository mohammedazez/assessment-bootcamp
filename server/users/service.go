package users

import "server/entity"

type Service interface {
	GetAllUser() ([]UserFormatter, error)
	GetUserByIDService (userID string) (UserFormatter, error)
	MakeNewUserService(userInput UserRegisterInput)(UserFormatter, error)
	UserLoginService(userInput UserLoginInput) (UserLoginFormatter, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUser() ([]UserFormatter, error) {
	users, err := s.repository.GetAllUsers()

	if err != nil {
		return []UserFormatter{}, nil
	}

	formatters := UserFormats(users)

	return formatters, nil
}

func (s *service) GetUserByIDService (userID string) (UserFormatter, error) {
	users, err := s.repository.GetUsersByID(userID)

	if err != nil {
		return UserFormatter{}, nil
	}

	formatter := UserFormat(users)

	return formatter, nil
}

func (s *service) MakeNewUserService(userInput UserRegisterInput)(UserFormatter, error)  {
	var NewUser = entity.User{
		Fullname: userInput.FullName,
		Address: userInput.Address,
		Email: userInput.Email,
		Password: userInput.Password,
	}

	users, err := s.repository.MakeNewUser(NewUser)

	if err != nil {
		return UserFormatter{}, err
	}

	formatter := UserFormat(users)
	return formatter, nil
}

func (s *service) UserLoginService(input UserLoginInput) (UserLoginFormatter, error) {
	checkUser, err := s.repository.UserWantToLogin(input.Email)

	if err != nil {
		return UserLoginFormat(), nil
	}
}
