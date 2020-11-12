package entity

//IFavourite struct containing data for favourites space by user
type IFavourite struct {
	SpaceID string `json:"space_id"`
	UID     string `json:"uid"`
	UserID  string `json:"user_id"`
}
