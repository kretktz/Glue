package repository

import (
	entity "glue/glue-backend-golang/entity"
)

// ISpaceRepository implements the method ListSpaces
type ISpaceRepository interface {
	ListSpaces() ([]entity.ISpace, error)
}
