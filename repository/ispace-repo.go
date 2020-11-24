package repository

import (
	"glue/glue-backend-golang/entity"
)

// ISpaceRepository implements the method ListSpaces and GetSpaceByID
type ISpaceRepository interface {
	ListSpaces() ([]entity.ISpace, error)
	GetSpaceByID(spaceID string) ([]entity.ISpace, error)
}
