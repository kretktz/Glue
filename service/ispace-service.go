package service

import (
	entity "glue/glue-backend-golang/entity"
	repository "glue/glue-backend-golang/repository"
)

// ISpaceService implements the methods
type ISpaceService interface {
	ListSpaces() ([]entity.ISpace, error)
}

var (
	spaceRepo repository.ISpaceRepository
)

//ListSpacesService creates a new service
func ListSpacesService(repository repository.ISpaceRepository) ISpaceService {
	spaceRepo = repository
	return &service{}
}

func (*service) ListSpaces() ([]entity.ISpace, error) {
	return spaceRepo.ListSpaces()
}
