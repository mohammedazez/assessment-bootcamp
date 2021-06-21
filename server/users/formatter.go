package users

import "server/entity"

type UserLoginFormatter struct {
	ID            int    `json:"id"`
	Fullname      string `json:"fullname"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	Authorization string `json:"authorization"`
}

type UserFormatter struct {
	ID            int    `json:"id"`
	Fullname      string `json:"fullname"`
	Address       string `json:"address"`
	Email         string `json:"email"`
}

func UserFormat(user entity.User) UserFormatter {
	return UserFormatter{
		ID:        user.ID,
		Fullname: user.Fullname,
		Email:     user.Email,
		Address:   user.Address,
	}
}

func UserLoginFormat(user entity.User, token string) UserLoginFormatter {
	return UserLoginFormatter{
		ID:            user.ID,
		Fullname:          user.Fullname,
		Email:         user.Email,
		Address:       user.Address,
		Authorization: token,
	}
}


func UserFormats(users []entity.User) []UserFormatter {
	var userFormats []UserFormatter

	for _, user := range users {
		userFormats = append(userFormats, UserFormat(user))
	}

	return userFormats
}
