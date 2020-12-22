package service

import (
	"errors"
	"glue/glue-backend-golang/entity"
	"glue/glue-backend-golang/repository"
)

// PlaceService implements the methods concerning places
type PlaceService interface {
	FireStoreValidate(place *entity.Place) error
	FireStoreCreate(place *entity.Place) (*entity.Place, error)
	FireStoreFindAll() ([]entity.Place, error)
}

type service struct{}

var (
	repo repository.PlaceRepository
)

//NewPlaceService creates a new service
func NewPlaceService(repository repository.PlaceRepository) PlaceService {
	repo = repository
	return &service{}
}

func (*service) FireStoreValidate(place *entity.Place) error {
	if place == nil {
		err := errors.New("the place is not specified")
		return err
	}
	return nil
}

func (*service) FireStoreCreate(place *entity.Place) (*entity.Place, error) {
	return repo.FireStoreSave(place)
}

func (*service) FireStoreFindAll() ([]entity.Place, error) {
	return repo.FireStoreFindAll()
}
