package service

import (
	"errors"
	"glue/glue-backend-golang/entity"
	"glue/glue-backend-golang/repository"
)

type (
	// ISpaceService implements the method ListSpaces
	ISpaceService interface {
		ListSpaces() ([]entity.ISpace, error)
		GetSpaceByID(spaceID string) ([]entity.ISpace, error)
		CreateSpace(space *entity.ISpace) (*entity.ISpace, error)
		ValidateSpace(e *entity.ISpace) error

		ListSpacesPsql() ([]entity.ISpace, error)
		CreateNewSpacePsql(space *entity.ISpace) (*entity.ISpace, error)
		GetSpaceByIDPsql(spaceID string) (entity.ISpace, error)
		ListSpacesWithTicketsPsql() ([]entity.ISpace, []entity.ITicket, error)
	}
)

var (
	spaceRepo repository.ISpaceRepository
	spaceID string
)



//SpacesService creates a new service for ISpace
func SpacesService(repository repository.ISpaceRepository) ISpaceService {
	spaceRepo = repository
	return &service{}
}

func (*service) ValidateSpace(space *entity.ISpace) error {
	if space == nil {
		err := errors.New("the space is not specified")
		return err
	}
	return nil
}

func (*service) CreateSpace(space *entity.ISpace) (*entity.ISpace, error) {
	return spaceRepo.SaveSpace(space)
}

func (*service) ListSpaces() ([]entity.ISpace, error) {
	return spaceRepo.ListSpaces()
}

func (*service) GetSpaceByID(spaceID string) ([]entity.ISpace, error) {
	return spaceRepo.GetSpaceByID(spaceID)
}

func (*service) CreateNewSpace(space *entity.ISpace) (*entity.ISpace, error){
	return spaceRepo.CreateNewSpace(space)
}

func (*service) ListSpacesPsql() ([]entity.ISpace, error) {
	return spaceRepo.ListSpacesPsql()
}

func (*service) CreateNewSpacePsql(space *entity.ISpace) (*entity.ISpace, error) {
	return spaceRepo.CreateNewSpacePsql(space)
}

func (*service) GetSpaceByIDPsql(spaceID string) (entity.ISpace, error) {
	return spaceRepo.GetSpaceByIDPsql(spaceID)
}

func (*service) ListSpacesWithTicketsPsql() ([]entity.ISpace, []entity.ITicket, error) {
	return spaceRepo.ListSpacesWithTicketsPsql()
}