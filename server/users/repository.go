package users

import (
	"gorm.io/gorm"
	"server/entity"
)

type Repository interface {
	GetAllUsers() ([]entity.User, error)
	GetUsersByID(ID string)(entity.User, error)
	MakeNewUser(newUser entity.User) (entity.User, error)
	UserWantToLogin (email string) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB)  *repository{
	return &repository{db}
}

func (r *repository)  GetAllUsers() ([]entity.User, error){
	var users []entity.User

	err := r.db.Find(&users).Error

	if err != nil {
		return users, nil
	}

	return users, nil
}

func (r *repository) GetUsersByID(ID string)(entity.User, error){
	var users entity.User

	err := r.db.Where("id = ?", ID).Find(&users).Error

	if err != nil {
		return users, nil
	}

	return users, nil
}

func (r *repository) MakeNewUser(newUser entity.User) (entity.User, error) {
	if err := r.db.Create(&newUser).Error; err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (r repository) UserWantToLogin(email string) (entity.User, error)  {
	var user entity.User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}