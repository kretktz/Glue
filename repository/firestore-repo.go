package repository

import (
	"context"
	"encoding/json"
	"fmt"
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
			// ticket := entity.Ticket{
			// 	TicketType:         doc.Data()["TicketType"].(string),
			// 	NumberTicketsAvail: doc.Data()["NumberTicketsAvail"].(int64),
			// }

			// tickets = append(tickets, ticket)

			aMap := make(map[string]interface{})
			aMap = doc.Data()
			b, err := json.Marshal(aMap)
			if err != nil {
				panic(err)
			}
			var ticket entity.Ticket
			err23 := json.Unmarshal([]byte(b), &ticket)
			if err23 != nil {
				fmt.Println("error:", err)
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
