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

func (mock *MockITicket) ListAllAvailableTickets() ([]entity.ITicket, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.ITicket), args.Error(1)
}

func TestService_ListAllAvailableTickets(t *testing.T) {
	mockRepo := new(MockITicket)

	ticket := entity.ITicket{
		Availability: 34,
		Colour:       "blue",
		Description:  "cheap",
		Name:         "cheap tick",
		Period:       2,
		Price:        500,
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
	mockRepo.On("ListAllAvailableTickets").Return([]entity.ITicket{ticket}, nil)

	testService := TicketService(mockRepo)

	result, _ := testService.ListAllAvailableTickets()

	//Mock Assertion
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, 34, result[0].Availability)
	assert.Equal(t, "blue", result[0].Colour)
	assert.Equal(t, "cheap", result[0].Description)
	assert.Equal(t, "cheap tick", result[0].Name)
	assert.Equal(t, 2, result[0].Period)
	assert.Equal(t, 500, result[0].Price)
	assert.Equal(t, "007", result[0].SpaceID)
	assert.Equal(t, []entity.ISpace{
		{
			Address:                "address",
			Availability:           "available",
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
