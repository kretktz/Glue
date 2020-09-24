package repository

import (
	"context"
	entity "glue/glue-backend-golang/entity"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// PlaceRepository implements the methods
type PlaceRepository interface {
	Save(place *entity.IPlace) (*entity.IPlace, error)
	FindAll() ([]entity.IPlace, error)
}

type repo struct{}

//NewPlaceRepository does stuff
func NewPlaceRepository() PlaceRepository {
	return &repo{}
}

const (
	projectID      string = "glue-25e3b"
	collectionName string = "IPlace"
)

func (*repo) Save(place *entity.IPlace) (*entity.IPlace, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ConfirmPageTitle": place.ConfirmPageTitle,
		"PhoneNumber":      place.PhoneNumber,
		"VisitPlaceName":   place.VisitPlaceName,
		"SlackSentMessage": place.SlackSentMessage,
		"SlackWebHookURL":  place.SlackWebHookURL,
	})
	if err != nil {
		log.Fatalf("Failed to add a new place: %v", err)
		return nil, err
	}
	return place, nil
}

func (*repo) FindAll() ([]entity.IPlace, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var places []entity.IPlace
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
		place := entity.IPlace{
			ConfirmPageTitle: doc.Data()["ConfirmPageTitle"].(string),
			PhoneNumber:      doc.Data()["PhoneNumber"].(string),
			VisitPlaceName:   doc.Data()["VisitPlaceName"].(string),
			SlackSentMessage: doc.Data()["SlackSentMessage"].(string),
			SlackWebHookURL:  doc.Data()["SlackWebHookURL"].(string),
		}
		places = append(places, place)
	}
	return places, nil
}
