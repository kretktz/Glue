package main

import (
	"encoding/json"
	entity "glue/glue-backend-golang/entity"
	repository "glue/glue-backend-golang/repository"
	"net/http"
)

// IPlace struct with details of a single space
// type IPlace struct {
// 	ConfirmPageTitle string `json:"confirmPageTitle"`
// 	PhoneNumber      string `json:"phoneNumber"`
// 	VisitPlaceName   string `json:"visitPlaceName"`
// 	SlackSentMessage string `json:"slackSentMessage"`
// 	SlackWebHookURL  string `json:"slackWebHookURL"`
// }

var (
	repo repository.PlaceRepository = repository.NewPlaceRepository()
)

// GetPlaces gets places
func GetPlaces(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	places, err := repo.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error fetching the places}`))
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(places)
}

// AddPlace adds a place
func AddPlace(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var place entity.IPlace
	err := json.NewDecoder(req.Body).Decode(&place)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling the request}`))
		return
	}
	// place.ID, err = rand.Int()
	// needs math/rand in imports
	repo.Save(&place)
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(place)

}
