package repository

import (
	"kdb/model"
)

type UserRepository struct {

}

func (repository *UserRepository) GetAllUsers() []model.User {

}

func (repository *UserRepository) GetUserById(id int) *model.User{

}

func (repository *UserRepository) GetUserByFullName(surname string, name string) []model.User {

}

func (repository *UserRepository) AddUser(user *model.User) *model.User{

}

func (repository *UserRepository) DeleteUser(id int) {

}

func (repository *UserRepository) ModifyUser(id int,user *model.User) *model.User{

}

func (repository *UserRepository) InitUserRepository() UserRepository {
	return UserRepository{}
}

