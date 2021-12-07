package service

import (
	"kdb/model"
	"kdb/repository"
)

var userRepository repository.UserRepository

func init() {
	userRepository = userRepository.InitUserRepository()
}

func GetAllUsers() []model.User {
	return userRepository.GetAllUsers()
}

func GetUserById(id int) *model.User{
	userRepository.GetUserById(id)
}

func GetUserByFullName(surname string, name string) *model.User{
	userRepository.GetUserByFullName(surname, name)
}

func CreateUser(u *model.User) *model.User{
	userRepository.AddUser(u)
}

func ModifyUser(id int, u *model.User) *model.User{
	userRepository.ModifyUser(id, u)
}

func DeleteUser(id int) {
	userRepository.DeleteUser(id)
}

