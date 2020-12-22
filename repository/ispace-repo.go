package repository

import (
	"glue/glue-backend-golang/entity"
)

type (
	// ISpaceRepository implements the methods concerning spaces
	ISpaceRepository interface {
		FireStoreListSpaces() ([]entity.ISpace, error)
		FireStoreGetSpaceByID(spaceID string) ([]entity.ISpace, error)
		FireStoreSaveSpace(space *entity.ISpace) (*entity.ISpace, error)
		FireStoreCreateNewSpace(space *entity.ISpace) (*entity.ISpace, error)

		PsqlListSpaces() ([]entity.ISpace, error)
		PsqlCreateNewSpace(space *entity.ISpace) (*entity.ISpace, error)
		PsqlGetSpaceByID(spaceID string) (entity.ISpace, error)
		PsqlListSpacesWithTickets() ([]entity.ISpace, []entity.ITicket, error)
	}
)

