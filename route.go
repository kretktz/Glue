package main

import (
	"encoding/json"
	"net/http"
)

// IPlace struct with details of a single space
type IPlace struct {
	ConfirmPageTitle string `json:"confirmPageTitle"`
	PhoneNumber      string `json:"phoneNumber"`
	VisitPlaceName   string `json:"visitPlaceName"`
	SlackSentMessage string `json:"slackSentMessage"`
	SlackWebHookURL  string `json:"slackWebHookURL"`
}

var (
	places []IPlace
)

func init() {
	places = []IPlace{{ConfirmPageTitle: "confirmed",
		PhoneNumber:      "9907e2245678",
		VisitPlaceName:   "Cambodia",
		SlackSentMessage: "no man no",
		SlackWebHookURL:  "someURL",
	}}
}

// GetPlaces gets places
func GetPlaces(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(places)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling the places array}`))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

// AddPlace adds a place
func AddPlace(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var place IPlace
	err := json.NewDecoder(req.Body).Decode(&place)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling the request}`))
		return
	}
	// place.Id = len(places) + 1
	places = append(places, place)
	res.WriteHeader(http.StatusOK)
	result, err := json.Marshal(places)
	res.Write(result)

}
