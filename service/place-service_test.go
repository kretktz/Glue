package service

import (
	"glue/glue-backend-golang/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) FireStoreSave(place *entity.Place) (*entity.Place, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Place), args.Error(1)
}

func (mock *MockRepository) FireStoreFindAll() ([]entity.Place, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Place), args.Error(1)
}

func Test_FireStoreFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	place := entity.Place{PlaceName: "Name", PlaceLocation: "Takamatsu", PhoneNumber: "5678"}
	// Setup Expectations
	mockRepo.On("FireStoreFindAll").Return([]entity.Place{place}, nil)

	testService := NewPlaceService(mockRepo)

	result, _ := testService.FireStoreFindAll()

	//Mock Assertion
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, "Name", result[0].PlaceName)
	assert.Equal(t, "Takamatsu", result[0].PlaceLocation)
	assert.Equal(t, "5678", result[0].PhoneNumber)
}

func Test_FireStoreCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	place := entity.Place{PlaceName: "Name", PlaceLocation: "Takamatsu", PhoneNumber: "5678"}

	//Setup expectations
	mockRepo.On("FireStoreSave").Return(&place, nil)

	testService := NewPlaceService(mockRepo)

	result, err := testService.FireStoreCreate(&place)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "Name", result.PlaceName)
	assert.Equal(t, "Takamatsu", result.PlaceLocation)
	assert.Equal(t, "5678", result.PhoneNumber)
	assert.Nil(t, err)
}

func Test_FireStoreValidateEmptyPlace(t *testing.T) {
	testService := NewPlaceService(nil)

	err := testService.FireStoreValidate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "the place is not specified", err.Error())
}
