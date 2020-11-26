package repository

import (
	entity "glue/glue-backend-golang/entity"
)

type (
	// PlaceRepository implements the methods Save and FindAll
	PlaceRepository interface {
		Save(place *entity.Place) (*entity.Place, error)
		FindAll() ([]entity.Place, error)
	}
)

