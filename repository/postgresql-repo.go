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

func PostgreConnect() {

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
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func (*repo) ListSpacesPsql() ([]entity.ISpace, error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqlStatement := `
		INSERT INTO public.ispace (address, availability, coordinates, description, image_urls, location, name, number_of_visitors, telephone_number, top_image_url, uid, visitor_greeting, visitor_slack_message, visitor_slack_webhook_url, website)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		RETURNING uid`
	id := 0
	err = db.QueryRow(sqlStatement, space.Name, space.Availability, space.Coordinates, space.Description, space.ImageURLS, space.Location, space.Name, space.NumberOfVisitors, space.TelephoneNumber, space.TopImageURL, space.UID, space.VisitorGreeting, space.VisitorSlackMessage, space.VisitorSlackWebhookURL, space.Website).Scan(&id)
	if err != nil {
		log.Fatalf("Failed to add a new space: %v", err)
		return nil, err
	}

	return space, nil

}