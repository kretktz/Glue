package service

import (
	"errors"
	entity "glue/glue-backend-golang/entity"
	repository "glue/glue-backend-golang/repository"
)

// PlaceService implements the methods
type PlaceService interface {
	Validate(place *entity.Place) error
	Create(place *entity.Place) (*entity.Place, error)
	FindAll() ([]entity.Place, error)
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

func (*service) Validate(place *entity.Place) error {
	if place == nil {
		err := errors.New("The place is not specified")
		return err
	}
	return nil
}

func (*service) Create(place *entity.Place) (*entity.Place, error) {
	return repo.Save(place)
}

func (*service) FindAll() ([]entity.Place, error) {
	return repo.FindAll()
}
