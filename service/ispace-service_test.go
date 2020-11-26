package service

import (
	"glue/glue-backend-golang/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockISpace struct {
	mock.Mock
}

func (mock *MockISpace) ListSpaces() ([]entity.ISpace, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.ISpace), args.Error(1)
}

func (mock *MockISpace) GetSpaceByID(spaceID string) ([]entity.ISpace, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.ISpace), args.Error(1)
}

func TestListSpaces(t *testing.T) {
	mockRepo := new(MockISpace)

	space := entity.ISpace{Address: "some address",
		Availability:     "available",
		Coordinates:      "some coordinates",
		Description:      "some description",
		ImageURLS:        "some url.com",
		Location:         "some location",
		Name:             "some name",
		NumberOfVisitors: "some number",
		TelephoneNumber:  "some number",
		Tickets: []entity.ITicket{
			entity.ITicket{
				Colour:      "some colour",
				Description: "some description",
				Name:        "some name",
				Period:      64,
				Price:       64,
				SpaceID:     "some id",
				UID:         "some uid",
			},
		},
		TopImageURL:            "some url.com",
		UID:                    "some id",
		VisitorGreeting:        "some greeting",
		VisitorSlackMessage:    "some message",
		VisitorSlackWebhookURL: "some url",
		Website:                "some website"}

	// Setup Expectations
	mockRepo.On("ListSpaces").Return([]entity.ISpace{space}, nil)

	testService := SpacesService(mockRepo)

	result, _ := testService.ListSpaces()

	//Mock Assertion
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, "some address", result[0].Address)
	assert.Equal(t, "available", result[0].Availability)
	assert.Equal(t, "some coordinates", result[0].Coordinates)
	assert.Equal(t, "some description", result[0].Description)
	assert.Equal(t, "some url.com", result[0].ImageURLS)
	assert.Equal(t, "some location", result[0].Location)
	assert.Equal(t, "some name", result[0].Name)
	assert.Equal(t, "some number", result[0].NumberOfVisitors)
	assert.Equal(t, "some number", result[0].TelephoneNumber)
	assert.Equal(t, []entity.ITicket{
		entity.ITicket{
			Colour:      "some colour",
			Description: "some description",
			Name:        "some name",
			Period:      64,
			Price:       64,
			SpaceID:     "some id",
			UID:         "some uid",
		},
	}, result[0].Tickets)
	assert.Equal(t, "some url.com", result[0].TopImageURL)
	assert.Equal(t, "some id", result[0].UID)
	assert.Equal(t, "some greeting", result[0].VisitorGreeting)
	assert.Equal(t, "some message", result[0].VisitorSlackMessage)
	assert.Equal(t, "some url", result[0].VisitorSlackWebhookURL)
	assert.Equal(t, "some website", result[0].Website)
}

func TestService_GetSpaceByID(t *testing.T) {
	mockRepo := new(MockISpace)

	space := entity.ISpace{Address: "some address",
		Availability:     "available",
		Coordinates:      "some coordinates",
		Description:      "some description",
		ImageURLS:        "some url.com",
		Location:         "some location",
		Name:             "some name",
		NumberOfVisitors: "some number",
		TelephoneNumber:  "some number",
		Tickets: []entity.ITicket{
			entity.ITicket{
				Colour:      "some colour",
				Description: "some description",
				Name:        "some name",
				Period:      64,
				Price:       64,
				SpaceID:     "some id",
				UID:         "some uid",
			},
		},
		TopImageURL:            "some url.com",
		UID:                    "some id",
		VisitorGreeting:        "some greeting",
		VisitorSlackMessage:    "some message",
		VisitorSlackWebhookURL: "some url",
		Website:                "some website"}

	// Setup Expectations
	mockRepo.On("GetSpaceByID").Return([]entity.ISpace{space}, nil)

	testService := SpacesService(mockRepo)

	result, _ := testService.GetSpaceByID(spaceID)

	//Mock Assertion
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, "some address", result[0].Address)
	assert.Equal(t, "available", result[0].Availability)
	assert.Equal(t, "some coordinates", result[0].Coordinates)
	assert.Equal(t, "some description", result[0].Description)
	assert.Equal(t, "some url.com", result[0].ImageURLS)
	assert.Equal(t, "some location", result[0].Location)
	assert.Equal(t, "some name", result[0].Name)
	assert.Equal(t, "some number", result[0].NumberOfVisitors)
	assert.Equal(t, "some number", result[0].TelephoneNumber)
	assert.Equal(t, []entity.ITicket{
		entity.ITicket{
			Colour:      "some colour",
			Description: "some description",
			Name:        "some name",
			Period:      64,
			Price:       64,
			SpaceID:     "some id",
			UID:         "some uid",
		},
	}, result[0].Tickets)
	assert.Equal(t, "some url.com", result[0].TopImageURL)
	assert.Equal(t, "some id", result[0].UID)
	assert.Equal(t, "some greeting", result[0].VisitorGreeting)
	assert.Equal(t, "some message", result[0].VisitorSlackMessage)
	assert.Equal(t, "some url", result[0].VisitorSlackWebhookURL)
	assert.Equal(t, "some website", result[0].Website)
}