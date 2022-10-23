package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	FindAll() ([]User, error)
	FindById(ID int) (User, error)
	FindOffsetLimit(a int, b int) ([]User, error)
	Create(user User) (User, error)
	UpdateUser(ID int, user User) (User, error)
	DeleteUser(ID int) (User, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]User, error) {

	users, err := s.repository.FindAll()

	return users, err
}

func (s *service) FindById(ID int) (User, error) {
	user, err := s.repository.FindById(ID)
	return user, err
}

func (s *service) FindOffsetLimit(a int, b int) ([]User, error) {
	users, err := s.repository.FindOffsetLimit(a, b)

	return users, err
}

func (s *service) Create(user User) (User, error) {

	user = User{
		Username: user.Username,
		Password: hashAndSalt(user.Password),
		Name:     user.Name,
	}
	newUser, err := s.repository.Create(user)

	return newUser, err
}

func (s *service) UpdateUser(ID int, user User) (User, error) {

	u, err := s.repository.FindById(ID)

	u.Username = user.Username
	u.Password = hashAndSalt(user.Password)
	u.Name = user.Name

	newUser, err := s.repository.UpdateUser(u)

	return newUser, err
}

func (s *service) DeleteUser(ID int) (User, error) {

	u, err := s.repository.FindById(ID)
	newUser, err := s.repository.DeleteUser(u)

	return newUser, err
}

func hashAndSalt(pwd string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
