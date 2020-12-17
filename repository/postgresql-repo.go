package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"glue/glue-backend-golang/entity"
	"log"
)

//NewPostgreRepository creates a new repository
func NewPsqlRepository() ISpaceRepository {
	return &repo{}
}

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "sedes123"
	dbname = "calhounio_demo"
)

func PsqlConnect() *sql.DB {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func (*repo) ListSpacesPsql() ([]entity.ISpace, error){
	db := PsqlConnect()

	defer db.Close()

	var spaces []entity.ISpace

	rows, err := db.Query("SELECT * FROM public.ispace LIMIT $1", 10)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var space entity.ISpace
		err = rows.Scan(
			&space.Address,
			&space.Availability,
			&space.Coordinates,
			&space.Description,
			&space.ImageURLS,
			&space.Location,
			&space.Name,
			&space.NumberOfVisitors,
			&space.TelephoneNumber,
			&space.TopImageURL,
			&space.UID,
			&space.VisitorGreeting,
			&space.VisitorSlackMessage,
			&space.VisitorSlackWebhookURL,
			&space.Website,
		)
		if err != nil {
			panic(err)
		}
		spaces = append(spaces, space)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return spaces, nil
}

func (*repo) CreateNewSpacePsql(space *entity.ISpace) (*entity.ISpace, error) {
	db := PsqlConnect()
	defer db.Close()


	sqlStatement := `
		INSERT INTO public.ispace (address, availability, coordinates, description, image_urls, location, name, number_of_visitors, telephone_number, top_image_url, uid, visitor_greeting, visitor_slack_message, visitor_slack_webhook_url, website)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		RETURNING uid`
	id := 0
	err := db.QueryRow(sqlStatement, space.Address, space.Availability, space.Coordinates, space.Description, space.ImageURLS, space.Location, space.Name, space.NumberOfVisitors, space.TelephoneNumber, space.TopImageURL, space.UID, space.VisitorGreeting, space.VisitorSlackMessage, space.VisitorSlackWebhookURL, space.Website).Scan(&id)
	if err != nil {
		log.Fatalf("Failed to add a new space: %v", err)
		return nil, err
	}

	return space, nil

}

func (*repo) GetSpaceByIDPsql(spaceID string) (entity.ISpace, error) {
	db := PsqlConnect()

	defer db.Close()

	var space entity.ISpace

	rows, err := db.Query("SELECT * FROM public.ispace WHERE uid = $1", spaceID)
	if err != nil {
		log.Fatalf("Couldn't fetch the space: %v", err)
		return space, err
	}
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(
			&space.Address,
			&space.Availability,
			&space.Coordinates,
			&space.Description,
			&space.ImageURLS,
			&space.Location,
			&space.Name,
			&space.NumberOfVisitors,
			&space.TelephoneNumber,
			&space.TopImageURL,
			&space.UID,
			&space.VisitorGreeting,
			&space.VisitorSlackMessage,
			&space.VisitorSlackWebhookURL,
			&space.Website,
		)
		if err != nil {
			log.Fatalf("Could not fetch the Space details: %v", err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error GetSpacebyIDPsql: %v", err)
	}

	return space, nil
}