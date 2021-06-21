package users

import (
	"gorm.io/gorm"
	"server/entity"
)

type Repository interface {
	GetAllUsers() ([]entity.User, error)
	GetUsersByID(ID string)([]entity.User, error)
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

func (r *repository) GetUsersByID(ID string)([]entity.User, error){
	var users []entity.User

	err := r.db.Where("id = ?", ID).Find(&users).Error

	if err != nil {
		return users, nil
	}

	return users, nil
}