package service

import (
	"errors"
	"glue/glue-backend-golang/entity"
	"glue/glue-backend-golang/repository"
)

type (
	// ISpaceService implements the method ListSpaces
	ISpaceService interface {
		FireStoreListSpaces() ([]entity.ISpace, error)
		FireStoreGetSpaceByID(spaceID string) ([]entity.ISpace, error)
		FireStoreCreateSpace(space *entity.ISpace) (*entity.ISpace, error)
		FireStoreValidateSpace(e *entity.ISpace) error

		PsqlListSpaces() ([]entity.ISpace, error)
		PsqlCreateNewSpace(space *entity.ISpace) (*entity.ISpace, error)
		PsqlGetSpaceByID(spaceID string) (entity.ISpace, error)
		PsqlListSpacesWithTickets() ([]entity.ISpace, []entity.ITicket, error)
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

func (*service) FireStoreValidateSpace(space *entity.ISpace) error {
	if space == nil {
		err := errors.New("the space is not specified")
		return err
	}
	return nil
}

func (*service) FireStoreCreateSpace(space *entity.ISpace) (*entity.ISpace, error) {
	return spaceRepo.FireStoreSaveSpace(space)
}

func (*service) FireStoreListSpaces() ([]entity.ISpace, error) {
	return spaceRepo.FireStoreListSpaces()
}

func (*service) FireStoreGetSpaceByID(spaceID string) ([]entity.ISpace, error) {
	return spaceRepo.FireStoreGetSpaceByID(spaceID)
}

func (*service) FireStoreCreateNewSpace(space *entity.ISpace) (*entity.ISpace, error){
	return spaceRepo.FireStoreCreateNewSpace(space)
}

func (*service) PsqlListSpaces() ([]entity.ISpace, error) {
	return spaceRepo.PsqlListSpaces()
}

func (*service) PsqlCreateNewSpace(space *entity.ISpace) (*entity.ISpace, error) {
	return spaceRepo.PsqlCreateNewSpace(space)
}

func (*service) PsqlGetSpaceByID(spaceID string) (entity.ISpace, error) {
	return spaceRepo.PsqlGetSpaceByID(spaceID)
}

func (*service) PsqlListSpacesWithTickets() ([]entity.ISpace, []entity.ITicket, error) {
	return spaceRepo.PsqlListSpacesWithTickets()
}