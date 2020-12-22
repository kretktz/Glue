package repository

import (
	"glue/glue-backend-golang/entity"
)

type (
	// PlaceRepository implements the methods Save and FindAll
	PlaceRepository interface {
		FireStoreSave(place *entity.Place) (*entity.Place, error)
		FireStoreFindAll() ([]entity.Place, error)
	}
)

