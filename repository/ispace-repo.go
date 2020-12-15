package repository

import (
	"glue/glue-backend-golang/entity"
)

type (
	// ISpaceRepository implements the method ListSpaces and GetSpaceByID
	ISpaceRepository interface {
		ListSpaces() ([]entity.ISpace, error)
		GetSpaceByID(spaceID string) ([]entity.ISpace, error)
		SaveSpace(space *entity.ISpace) (*entity.ISpace, error)
		CreateNewSpace(space *entity.ISpace) (*entity.ISpace, error)

		ListSpacesPsql() ([]entity.ISpace, error)
		CreateNewSpacePsql(space *entity.ISpace) (*entity.ISpace, error)
	}
)

