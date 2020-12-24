package service

import (
	"glue/glue-backend-golang/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockITicket struct {
	mock.Mock
}

func (mock *MockITicket) PsqlCreateNewTicket(ticket *entity.ITicket) (*entity.ITicket, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.ITicket), args.Error(1)
}

func TestService_PsqlCreateNewTicket(t *testing.T) {
	mockRepo := new(MockITicket)

	var(
		avail int64 = 2
		period int64 = 2
		price int64 = 2
	)

	ticket := entity.ITicket{
		Availability: avail,
		Colour:       "colour",
		Description:  "description",
		Name:         "name",
		Period:       period,
		Price:        price,
		SpaceID:      "space_id",
		UID:          "uid",
	}

	//Setup Expectations
	mockRepo.On("PsqlCreateNewTicket").Return(&ticket, nil)

	testService := TicketService(mockRepo)

	result, err := testService.PsqlCreateNewTicket(&ticket)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, avail, result.Availability)
	assert.Equal(t, "colour", result.Colour)
	assert.Equal(t, "description", result.Description)
	assert.Equal(t, "name", result.Name)
	assert.Equal(t, period, result.Period)
	assert.Equal(t, price, result.Price)
	assert.Equal(t, "space_id", result.SpaceID)
	assert.Equal(t, "uid", result.UID)
	assert.Nil(t, err)
}

func (mock *MockITicket) FireStoreListAllAvailableTickets() ([]entity.ITicket, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.ITicket), args.Error(1)
}

func TestService_FireStoreListAllAvailableTickets(t *testing.T) {
	mockRepo := new(MockITicket)

	var(
		availability int64 = 34
		period       int64 = 2
		price        int64 = 500
	)

	ticket := entity.ITicket{
		Availability: availability,
		Colour:       "blue",
		Description:  "cheap",
		Name:         "cheap tick",
		Period:       period,
		Price:        price,
		SpaceID:      "007",
		Space: []entity.ISpace{
			{
				Address:                "address",
				Availability:           "availability",
				Coordinates:            "coordinates",
				Description:            "description",
				ImageURLS:              "image urls",
				Location:               "location",
				Name:                   "name",
				NumberOfVisitors:       "number of visitors",
				TelephoneNumber:        "telephone number",
				Tickets:                nil,
				TopImageURL:            "top image url",
				UID:                    "uid",
				VisitorGreeting:        "visitor greeting",
				VisitorSlackMessage:    "visitor slack message",
				VisitorSlackWebhookURL: "visitor slack webhook url",
							Website:                "website",
						},
		},
		UID: "007",
	}

	//Setup expectations
	mockRepo.On("FireStoreListAllAvailableTickets").Return([]entity.ITicket{ticket}, nil)

	testService := TicketService(mockRepo)

	result, _ := testService.FireStoreListAllAvailableTickets()

	//Mock Assertion
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, availability, result[0].Availability)
	assert.Equal(t, "blue", result[0].Colour)
	assert.Equal(t, "cheap", result[0].Description)
	assert.Equal(t, "cheap tick", result[0].Name)
	assert.Equal(t, period, result[0].Period)
	assert.Equal(t, price, result[0].Price)
	assert.Equal(t, "007", result[0].SpaceID)
	assert.Equal(t, []entity.ISpace{
		{
			Address:                "address",
			Availability:           "availability",
			Coordinates:            "coordinates",
			Description:            "description",
			ImageURLS:              "image urls",
			Location:               "location",
			Name:                   "name",
			NumberOfVisitors:       "number of visitors",
			TelephoneNumber:        "telephone number",
			Tickets:                nil,
			TopImageURL:            "top image url",
			UID:                    "uid",
			VisitorGreeting:        "visitor greeting",
			VisitorSlackMessage:    "visitor slack message",
			VisitorSlackWebhookURL: "visitor slack webhook url",
			Website:                "website",
				},
	}, result[0].Space)
	assert.Equal(t, "007", result[0].UID)
}