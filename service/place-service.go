package service

import (
	"errors"
	entity "glue/glue-backend-golang/entity"
	repository "glue/glue-backend-golang/repository"
)

// PlaceService implements the methods
type PlaceService interface {
	Validate(place *entity.IPlace) error
	Create(place *entity.IPlace) (*entity.IPlace, error)
	FindAll() ([]entity.IPlace, error)
}

type service struct{}

var (
	repo repository.PlaceRepository
)

//NewPlacesService creates a new service
func NewPlacesService(repository repository.PlaceRepository) PlaceService {
	repo = repository
	return &service{}
}

func (*service) Validate(place *entity.IPlace) error {
	if place == nil {
		err := errors.New("The place is not specified")
		return err
	}
	return nil
}

func (*service) Create(place *entity.IPlace) (*entity.IPlace, error) {
	return repo.Save(place)
}

func (*service) FindAll() ([]entity.IPlace, error) {
	return repo.FindAll()
}
