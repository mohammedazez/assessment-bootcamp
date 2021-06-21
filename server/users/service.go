package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"server/auth"
	"server/entity"
)

type Service interface {
	GetAllUser() ([]UserFormatter, error)
	GetUserByIDService(userID string) (UserFormatter, error)
	MakeNewUserService(userInput UserRegisterInput) (UserFormatter, error)
	UserLoginService(userInput UserLoginInput) (UserLoginFormatter, error)
}

type service struct {
	repository  Repository
	authService auth.Service
}

func NewService(repository Repository, authService auth.Service) *service {
	return &service{repository, authService}
}

func (s *service) GetAllUser() ([]UserFormatter, error) {
	users, err := s.repository.GetAllUsers()

	if err != nil {
		return []UserFormatter{}, nil
	}

	formatters := UserFormats(users)

	return formatters, nil
}

func (s *service) GetUserByIDService(userID string) (UserFormatter, error) {
	users, err := s.repository.GetUsersByID(userID)

	if err != nil {
		return UserFormatter{}, nil
	}

	formatter := UserFormat(users)

	return formatter, nil
}

func (s *service) MakeNewUserService(userInput UserRegisterInput) (UserFormatter, error) {
	var NewUser = entity.User{
		Fullname: userInput.FullName,
		Address:  userInput.Address,
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	users, err := s.repository.MakeNewUser(NewUser)

	if err != nil {
		return UserFormatter{}, err
	}

	formatter := UserFormat(users)
	return formatter, nil
}

func (s *service) UserLoginService(userInput UserLoginInput) (UserLoginFormatter, error) {

	checkUser, err := s.repository.UserWantToLogin(userInput.Email)

	if err != nil {
		return UserLoginFormatter{}, err
	}

	if checkUser.ID == 0 || len(checkUser.Fullname) <= 1 {
		return UserLoginFormatter{}, errors.New("user email or password invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(userInput.Password))

	if err != nil {
		return UserLoginFormatter{}, errors.New("user email or password invalid")
	}

	token, _ := s.authService.GenerateToken(checkUser.ID)

	formatter := UserLoginFormat(checkUser, token)

	return formatter, nil
}
