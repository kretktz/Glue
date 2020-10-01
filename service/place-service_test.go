package service

import (
	entity "glue/glue-backend-golang/entity"
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

	place := entity.Place{ConfirmPageTitle: "Title", PhoneNumber: "5678", VisitPlaceName: "Some place Name", SlackSentMessage: "msg", SlackWebHookURL: "someURL"}
	// Setup Expectations
	mockRepo.On("FindAll").Return([]entity.Place{place}, nil)

	testService := NewPlaceService(mockRepo)

	result, _ := testService.FindAll()

	//Mock Assertion
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, "Title", result[0].ConfirmPageTitle)
	assert.Equal(t, "5678", result[0].PhoneNumber)
	assert.Equal(t, "Some place Name", result[0].VisitPlaceName)
	assert.Equal(t, "msg", result[0].SlackSentMessage)
	assert.Equal(t, "someURL", result[0].SlackWebHookURL)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	place := entity.Place{ConfirmPageTitle: "Title", PhoneNumber: "5678", VisitPlaceName: "Some place Name", SlackSentMessage: "msg", SlackWebHookURL: "someURL"}

	//Setup expectations
	mockRepo.On("Save").Return(&place, nil)

	testService := NewPlaceService(mockRepo)

	result, err := testService.Create(&place)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "Title", result.ConfirmPageTitle)
	assert.Equal(t, "5678", result.PhoneNumber)
	assert.Equal(t, "Some place Name", result.VisitPlaceName)
	assert.Equal(t, "msg", result.SlackSentMessage)
	assert.Equal(t, "someURL", result.SlackWebHookURL)
	assert.Nil(t, err)
}

func TestValidateEmptyPlace(t *testing.T) {
	testService := NewPlaceService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The place is not specified", err.Error())
}
