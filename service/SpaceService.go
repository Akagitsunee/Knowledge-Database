package service

import (
	"kdb/model"
	"kdb/repository"
)

var spaceRepository repository.SpaceRepository

func init() {
	spaceRepository = spaceRepository.InitSpaceRepository()
}

func GetAllSpaces() []model.Space {
	return spaceRepository.GetAllSpaces()
}

func GetSpaceById(id int) {
	spaceRepository.GetSpaceById(id)
}

func GetSpaceByFullName(name string) {
	spaceRepository.GetSpaceByName(name)
}

func CreateSpace(s *model.Space) {
	spaceRepository.AddSpace(s)
}

func ModifySpace(id int, s *model.Space) {
	spaceRepository.ModifySpace(id, s)
}

func DeleteSpace(id int) {
	spaceRepository.DeleteSpace(id)
}
