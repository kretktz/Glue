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

func (mock *MockRepository) Save(place *entity.Place) (*entity.Place, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Place), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Place, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Place), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	place := entity.Place{PlaceName: "Name", PlaceLocation: "Takamatsu", PhoneNumber: "5678"}
	// Setup Expectations
	mockRepo.On("FindAll").Return([]entity.Place{place}, nil)

	testService := NewPlaceService(mockRepo)

	result, _ := testService.FindAll()

	//Mock Assertion
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, "Name", result[0].PlaceName)
	assert.Equal(t, "Takamatsu", result[0].PlaceLocation)
	assert.Equal(t, "5678", result[0].PhoneNumber)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	place := entity.Place{PlaceName: "Name", PlaceLocation: "Takamatsu", PhoneNumber: "5678"}

	//Setup expectations
	mockRepo.On("Save").Return(&place, nil)

	testService := NewPlaceService(mockRepo)

	result, err := testService.Create(&place)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "Name", result.PlaceName)
	assert.Equal(t, "Takamatsu", result.PlaceLocation)
	assert.Equal(t, "5678", result.PhoneNumber)
	assert.Nil(t, err)
}

func TestValidateEmptyPlace(t *testing.T) {
	testService := NewPlaceService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The place is not specified", err.Error())
}
