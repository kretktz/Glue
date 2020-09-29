package repository

import (
	entity "glue/glue-backend-golang/entity"
)

// PlaceRepository implements the methods Save and FindAll
type PlaceRepository interface {
	Save(place *entity.IPlace) (*entity.IPlace, error)
	FindAll() ([]entity.IPlace, error)
}
