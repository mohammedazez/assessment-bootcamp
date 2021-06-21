package lists

import (
	"gorm.io/gorm"
	"server/entity"
)

type Repository interface {
	GetAllLists() ([]entity.List, error)
	GetListssByID(ID string)([]entity.List, error)
	MakeNewList(newList entity.List) (entity.List, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB)  *repository{
	return &repository{db}
}

func (r *repository)  GetAllLists() ([]entity.List, error){
	var lists []entity.List

	err := r.db.Find(&lists).Error

	if err != nil {
		return lists, nil
	}

	return lists, nil
}


func (r *repository) GetListssByID(ID string)([]entity.List, error){
	var lists []entity.List

	err := r.db.Where("id = ?", ID).Find(&lists).Error

	if err != nil {
		return lists, nil
	}

	return lists, nil
}

func (r *repository) MakeNewList(newList entity.List) (entity.List, error) {
	if err := r.db.Create(&newList).Error; err != nil {
		return newList, err
	}

	return newList, nil
}