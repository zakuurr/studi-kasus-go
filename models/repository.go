package models

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]User, error)
	FindById(ID int) (User, error)
	FindOffsetLimit(a int, b int) ([]User, error)
	Create(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) FindById(ID int) (User, error) {
	var user User

	err := r.db.Find(&user, ID).Error

	return user, err
}

func (r *repository) FindOffsetLimit(a int, b int) ([]User, error) {
	var users []User

	err := r.db.Limit(a).Offset(b).Find(&users).Error

	return users, err
}

func (r *repository) Create(user User) (User, error) {

	err := r.db.Create(&user).Error

	return user, err
}
func (r *repository) UpdateUser(user User) (User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user User) (User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
