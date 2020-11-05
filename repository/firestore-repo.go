package repository

import (
	"context"
	entity "glue/glue-backend-golang/entity"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type repo struct{}

//NewFirestoreRepository creates a new repository
func NewFirestoreRepository() PlaceRepository {
	return &repo{}
}

//NewISpaceRepository creates a new repository
func NewISpaceRepository() ISpaceRepository {
	return &repo{}
}

const (
	projectID      string = "glue-25e3b"
	collectionName string = "Place"
)

func (*repo) Save(place *entity.Place) (*entity.Place, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	//TODO: Wrap data into json unmarshal func
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
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

func (*repo) FindAll() ([]entity.Place, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var places []entity.Place
	it := client.Collection(collectionName).Documents(ctx)
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
		/*
			var place entity.Place
			if err := doc.DataTo(&place); err != nil {
				log.Fatalf("Failed to fetch place data: %v", err)
			}

			var tickets []entity.Ticket
			something := client.Collection(collectionName)
			for _, t := range tickets {
				if _, err := something.Doc(t.PlaceName).Collection("Ticket").NewDoc().Set(ctx, map[string]interface{}{
					"TicketType":         t.TicketType,
					"NumberTicketsAvail": t.NumberTicketsAvail,
				}); err != nil {
					log.Fatalf("Failed to fetch ticket data: %v", err)
				}
			}
		*/
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

func (*repo) ListSpaces() ([]entity.ISpace, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var (
		spaces []entity.ISpace
	)
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

		// getting ticket data
		//method 1 - hardcoded
		/*
		ticketSnap, err := client.Collection("ITicket").Doc("RX9YD3fPXoKDesjbxxEI").Get(ctx)
		if err != nil {
			return nil, err
		}
		var ticket entity.ITicket
		ticketSnap.DataTo(&ticket)
		*/
		/*
		// method 2 - in development
		var tempID string
		tempID = space.UID
		ticketSnap := client.Collection("ITicket").Where("UID", "==", tempID).Snapshots(ctx)
		var ticket entity.ITicket
		ticketSnap.
		*/


		// method 3 - nested for loop

		var (
			tempID string
			ticket entity.ITicket
			tickets []entity.ITicket
		)
		tempID = space.UID
		ticketSnap := client.Collection("ITicket").Where("UID", "==", tempID).Documents(ctx)
		for {
			doc, err := ticketSnap.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed to iterate over tickets: %v", err)
				return nil, err
			}
			if err := doc.DataTo(&ticket); err != nil {
				log.Fatalf("Failed to fetch ticket data: %v", err)
			}
			tickets = append(tickets, ticket)
		}

		// insert ticket data into space
		space.Tickets = ticket

		spaces = append(spaces, space)
	}
	return spaces, nil
}
