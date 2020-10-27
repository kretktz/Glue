package repository

import (
	entity "glue/glue-backend-golang/entity"
)

// ISpaceRepository implements the methods Save and FindAll
type ISpaceRepository interface {
	ListSpaces() ([]entity.ISpace, error)
}
