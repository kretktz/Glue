package repository

import (
	entity "glue/glue-backend-golang/entity"
)

// PlaceRepository implements the methods Save and FindAll
type PlaceRepository interface {
	Save(place *entity.Place) (*entity.Place, error)
	FindAll() ([]entity.Place, error)
}
