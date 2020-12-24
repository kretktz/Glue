package repository

import (
	"context"
	"glue/glue-backend-golang/entity"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type repo struct{}

//NewFirestoreRepository creates a new repository
func NewFirestoreRepository() PlaceRepository {
	return &repo{}
}

//NewISpaceRepository creates a new repository to accommodate ISpace related functions
func NewISpaceRepository() ISpaceRepository {
	return &repo{}
}

//NewITicketRepository creates a new repository to accommodate ITicket related functions
func NewITicketRepository() ITicketRepository {
	return &repo{}
}


const (
	projectID      string = "glue-25e3b"
)

// NewFirestoreClient creates a client to connect to Firestore DB
func NewFirestoreClient() (*firestore.Client, context.Context){
	contx := context.Background()
	client, err := firestore.NewClient(contx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
	}
	return client, contx
}

//Place related functions
func (*repo) FireStoreSave(place *entity.Place) (*entity.Place, error) {

	client, ctx := NewFirestoreClient()
	defer client.Close()

	//TODO: Wrap data into json unmarshal func
	_, _, err := client.Collection("Place").Add(ctx, map[string]interface{}{
		"PlaceName":     place.PlaceName,
		"PlaceLocation": place.PlaceLocation,
		"PhoneNumber":   place.PhoneNumber,
	})
	if err != nil {
		log.Fatalf("Failed to add a new place: %v", err)
		return nil, err
	}
	return place, nil
}

func (*repo) FireStoreFindAll() ([]entity.Place, error) {

	client, ctx := NewFirestoreClient()
	defer client.Close()

	var places []entity.Place
	it := client.Collection("Place").Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}

		ticketRef := doc.Ref.Collection("Ticket")
		var tickets []entity.Ticket
		it := ticketRef.Documents(ctx)

		for {
			doc, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed to iterate over tickets: %v", err)
				return nil, err
			}
			var ticket entity.Ticket
			if err := doc.DataTo(&ticket); err != nil {
				log.Fatalf("Failed to fetch ticket data: %v", err)
			}

			tickets = append(tickets, ticket)

		}

		place := entity.Place{
			PlaceName:     doc.Data()["PlaceName"].(string),
			PlaceLocation: doc.Data()["PlaceLocation"].(string),
			PhoneNumber:   doc.Data()["PhoneNumber"].(string),
			NumTickets:    tickets,
		}

		places = append(places, place)
	}
	return places, nil

}

// ISpace related functions

// FireStoreListSpaces lists all spaces along with tickets in the Firestore DB
func (*repo) FireStoreListSpaces() ([]entity.ISpace, error) {

	client, ctx:= NewFirestoreClient()
	defer client.Close()

	var (
		spaces []entity.ISpace
	)
	// fetching space data
	it := client.Collection("ISpace").Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}
		var space entity.ISpace
		if err := doc.DataTo(&space); err != nil {
			log.Fatalf("Failed to fetch space data: %v", err)
		}

		// fetching ticket data

		var (
			ticket entity.ITicket
			tickets []entity.ITicket
		)

		ticketSnap := client.Collection("ITicket").Where("UID", "==", space.UID)
		docs, err := ticketSnap.Documents(ctx).GetAll()
		if err != nil {
			log.Fatalf(
				"Failed to iterate over tickets: %v",
				err,
			)
		}
		for _, doc := range docs {
			doc.DataTo(&ticket)
			tickets = append(tickets, ticket)
		}

		// insert ticket data into a corresponding space
		space.Tickets = tickets
		// add the resulting space with its tickets to a list of spaces
		spaces = append(spaces, space)
	}
	return spaces, nil
}

// FireStoreGetSpaceByID fetches a space as specified by provided ID (UID)
func (*repo) FireStoreGetSpaceByID(spaceID string) ([]entity.ISpace, error) {

	client, ctx := NewFirestoreClient()
	defer client.Close()

	var (
		space entity.ISpace
		spaces []entity.ISpace
	)
	// fetching space data with its ticket options as specified by ID
	it := client.Collection("ISpace").Where("UID", "==", spaceID).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}
		if err := doc.DataTo(&space); err != nil {
			log.Fatalf("Failed to fetch space data: %v", err)
		}

		var (
			ticket entity.ITicket
			tickets []entity.ITicket
		)
		// fetching associated ticket data
		ticketSnap := client.Collection("ITicket").Where("UID", "==", space.UID)
		docs, err := ticketSnap.Documents(ctx).GetAll()
		if err != nil {
			log.Fatalf(
				"Failed to iterate over tickets: %v",
				err,
			)
		}
		for _, doc := range docs {
			doc.DataTo(&ticket)
			tickets = append(tickets, ticket)
		}

		// insert ticket data into space
		space.Tickets = tickets

		spaces = append(spaces, space)
	}
	return spaces, nil
}

// FireStoreCreateNewSpaces writes a new record to the Firestore DB
func (*repo) FireStoreCreateNewSpace(space *entity.ISpace) (*entity.ISpace, error){
	client, ctx := NewFirestoreClient()
	defer client.Close()

	doc, wr, err := client.Collection("ISpace").Add(ctx, entity.ISpace{
		Address:                space.Address,
		Availability:           space.Availability,
		Coordinates:            space.Coordinates,
		Description:            space.Description,
		ImageURLS:              space.ImageURLS,
		Location:               space.Location,
		Name:                   space.Name,
		NumberOfVisitors:       space.NumberOfVisitors,
		TelephoneNumber:        space.TelephoneNumber,
		Tickets:                space.Tickets,
		TopImageURL:            space.TopImageURL,
		UID:                    space.UID,
		VisitorGreeting:        space.VisitorGreeting,
		VisitorSlackMessage:    space.VisitorSlackMessage,
		VisitorSlackWebhookURL: space.VisitorSlackWebhookURL,
		Website:                space.Website,
	})
	doc.Snapshots(ctx)
	log.Printf("Data added: %v and %v", wr, doc)
	if err != nil {
		log.Fatalf("Failed to add a new space: %v", err)
		return nil, err
	}
	return space, nil
}

// ITickets Related functions

// FireStoreListAllAvailableTickets lists all tickets with availability > 0 along with associated ISpace
func (r *repo) FireStoreListAllAvailableTickets() ([]entity.ITicket, error) {

	client, ctx := NewFirestoreClient()
	defer client.Close()

	var (
		ticket entity.ITicket
		tickets []entity.ITicket
	)
	// fetching tickets data
	it := client.Collection("ITicket").Where("Availability", ">", 0).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}
		if err := doc.DataTo(&ticket); err != nil {
			log.Fatalf("Failed to fetch available tickets data: %v", err)
		}
			var (
				space entity.ISpace
				spaces []entity.ISpace
			)
			// fetching space data associated with the ticket
			spaceSnap := client.Collection("ISpace").Where("UID", "==", ticket.SpaceID)
			docs, err := spaceSnap.Documents(ctx).GetAll()
			if err != nil {
				log.Fatalf(
					"Failed to iterate over spaces: %v",
					err,
				)
			}
			for _, doc := range docs {
				doc.DataTo(&space)
				spaces = append(spaces, space)
			}

			// insert space data into ticket
			ticket.Space = spaces

		tickets = append(tickets, ticket)
	}
	return tickets, nil
}