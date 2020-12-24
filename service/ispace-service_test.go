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

// Create New Space functions test -------------------------------------------------------------------

func (mock *MockISpace) FireStoreCreateNewSpace(space *entity.ISpace) (*entity.ISpace, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.ISpace), args.Error(1)
}

func TestService_FireStoreCreateNewSpace(t *testing.T) {
	mockRepo := new(MockISpace)

	space := entity.ISpace{
		Address:                "address",
		Availability:           "availability",
		Coordinates:            "coordinates",
		Description:            "description",
		ImageURLS:              "image_urls",
		Location:               "location",
		Name:                   "name",
		NumberOfVisitors:       "number",
		TelephoneNumber:        "telephone",
		TopImageURL:            "top_image",
		UID:                    "uid",
		VisitorGreeting:        "greeting",
		VisitorSlackMessage:    "message",
		VisitorSlackWebhookURL: "webhook",
		Website:                "website",
	}

	//Setup Expectations
	mockRepo.On("FireStoreCreateNewSpace").Return(&space, nil)

	testService := SpacesService(mockRepo)

	result, err := testService.FireStoreCreateNewSpace(&space)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "address", result.Address)
	assert.Equal(t, "availability", result.Availability)
	assert.Equal(t, "coordinates", result.Coordinates)
	assert.Equal(t, "description", result.Description)
	assert.Equal(t, "image_urls", result.ImageURLS)
	assert.Equal(t, "location", result.Location)
	assert.Equal(t, "name", result.Name)
	assert.Equal(t, "number", result.NumberOfVisitors)
	assert.Equal(t, "telephone", result.TelephoneNumber)
	assert.Equal(t, "top_image", result.TopImageURL)
	assert.Equal(t, "uid", result.UID)
	assert.Equal(t, "greeting", result.VisitorGreeting)
	assert.Equal(t, "message", result.VisitorSlackMessage)
	assert.Equal(t, "webhook", result.VisitorSlackWebhookURL)
	assert.Equal(t, "website", result.Website)
	assert.Nil(t, err)
}

func (mock *MockISpace) PsqlCreateNewSpace(space *entity.ISpace) (*entity.ISpace, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.ISpace), args.Error(1)
}

func TestService_PsqlCreateNewSpace(t *testing.T) {
	mockRepo := new(MockISpace)

	space := entity.ISpace{
		Address:                "address",
		Availability:           "availability",
		Coordinates:            "coordinates",
		Description:            "description",
		ImageURLS:              "image_urls",
		Location:               "location",
		Name:                   "name",
		NumberOfVisitors:       "number",
		TelephoneNumber:        "telephone",
		TopImageURL:            "top_image",
		UID:                    "uid",
		VisitorGreeting:        "greeting",
		VisitorSlackMessage:    "message",
		VisitorSlackWebhookURL: "webhook",
		Website:                "website",
	}

	//Setup Expectations
	mockRepo.On("FireStoreCreateNewSpace").Return(&space, nil)

	testService := SpacesService(mockRepo)

	result, err := testService.FireStoreCreateNewSpace(&space)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "address", result.Address)
	assert.Equal(t, "availability", result.Availability)
	assert.Equal(t, "coordinates", result.Coordinates)
	assert.Equal(t, "description", result.Description)
	assert.Equal(t, "image_urls", result.ImageURLS)
	assert.Equal(t, "location", result.Location)
	assert.Equal(t, "name", result.Name)
	assert.Equal(t, "number", result.NumberOfVisitors)
	assert.Equal(t, "telephone", result.TelephoneNumber)
	assert.Equal(t, "top_image", result.TopImageURL)
	assert.Equal(t, "uid", result.UID)
	assert.Equal(t, "greeting", result.VisitorGreeting)
	assert.Equal(t, "message", result.VisitorSlackMessage)
	assert.Equal(t, "webhook", result.VisitorSlackWebhookURL)
	assert.Equal(t, "website", result.Website)
	assert.Nil(t, err)
}

// List Spaces functions tests ------------------------------------------------------------------------

func (mock *MockISpace) PsqlListSpaces() ([]entity.ISpace, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.ISpace), args.Error(1)
}

func TestService_PsqlListSpaces(t *testing.T) {
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
			{
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
	mockRepo.On("PsqlListSpaces").Return([]entity.ISpace{space}, nil)

	testService := SpacesService(mockRepo)

	result, _ := testService.PsqlListSpaces()

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
		{
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

func (mock *MockISpace) PsqlListSpacesWithTickets() ([]entity.ISpace, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.ISpace), args.Error(1)
}

func TestService_PsqlListSpacesWithTickets(t *testing.T) {
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
			{
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
	mockRepo.On("PsqlListSpaces").Return([]entity.ISpace{space}, nil)

	testService := SpacesService(mockRepo)

	result, _ := testService.PsqlListSpaces()

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
		{
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

func (mock *MockISpace) FireStoreListSpaces() ([]entity.ISpace, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.ISpace), args.Error(1)
}

func TestService_FireStoreListSpaces(t *testing.T) {
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
			{
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
	mockRepo.On("FireStoreListSpaces").Return([]entity.ISpace{space}, nil)

	testService := SpacesService(mockRepo)

	result, _ := testService.FireStoreListSpaces()

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
		{
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

// Fetch a particular space tests ---------------------------------------------------------------------

func (mock *MockISpace) PsqlGetSpaceByID(spaceID string) ([]entity.ISpace, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.ISpace), args.Error(1)
}

func TestService_PsqlGetSpaceByID(t *testing.T) {
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
			{
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


	mockRepo.On("PsqlGetSpaceByID").Return([]entity.ISpace{space}, nil)

	testService := SpacesService(mockRepo)

	result, _ := testService.PsqlGetSpaceByID(spaceID)

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
		{
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

func (mock *MockISpace) FireStoreGetSpaceByID(spaceID string) ([]entity.ISpace, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.ISpace), args.Error(1)
}

func TestService_FireStoreGetSpaceByID(t *testing.T) {
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
			{
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
	mockRepo.On("FireStoreGetSpaceByID").Return([]entity.ISpace{space}, nil)

	testService := SpacesService(mockRepo)

	result, _ := testService.FireStoreGetSpaceByID(spaceID)

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
		{
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