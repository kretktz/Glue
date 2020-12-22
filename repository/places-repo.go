package repository

import "glue/glue-backend-golang/entity"

// PlaceRepository implements the methods concerning places
type PlaceRepository interface {
		FireStoreSave(place *entity.Place) (*entity.Place, error)
		FireStoreFindAll() ([]entity.Place, error)
	}