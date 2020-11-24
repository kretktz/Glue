package service

import (
	"glue/glue-backend-golang/entity"
	"glue/glue-backend-golang/repository"
)

// ISpaceService implements the method ListSpaces
type ISpaceService interface {
	ListSpaces() ([]entity.ISpace, error)
	GetSpaceByID(spaceID string) ([]entity.ISpace, error)
}

var spaceRepo repository.ISpaceRepository
var spaceID string


//SpacesService creates a new service for ISpace
func SpacesService(repository repository.ISpaceRepository) ISpaceService {
	spaceRepo = repository
	return &service{}
}


func (*service) ListSpaces() ([]entity.ISpace, error) {
	return spaceRepo.ListSpaces()
}

func (*service) GetSpaceByID(spaceID string) ([]entity.ISpace, error) {
	return spaceRepo.GetSpaceByID(spaceID)
}
