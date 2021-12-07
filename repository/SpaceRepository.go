package repository

import "kdb/model"

type SpaceRepository struct {

}

func (repository *SpaceRepository) GetAllSpaces() []model.Space {

}

func (repository *SpaceRepository) GetSpaceById(id int) model.Space  {

}

func (repository *SpaceRepository) GetSpaceByName(name string) model.Space {

}


func (repository *SpaceRepository) AddSpace(space *model.Space) model.Space {

}

func (repository *SpaceRepository) DeleteSpace(id int) {

}

func (repository *SpaceRepository) ModifySpace(id int,space *model.Space) model.Space {

}

func (repository *SpaceRepository) InitSpaceRepository() SpaceRepository {
	return SpaceRepository{}
}