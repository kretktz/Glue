package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"glue/glue-backend-golang/entity"
	"log"
)

//NewPsqlRepository creates a new repository
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

// PsqlConnect creates a new connection to the PostgreSQL DB
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

// PsqlListSpaces returns a list of spaces (up to 10) from PostgreSQL DB
func (*repo) PsqlListSpaces() ([]entity.ISpace, error){
	// connecting to DB
	db := PsqlConnect()
	// keeping the connection open
	defer db.Close()

	var spaces []entity.ISpace
	// Psql query build
	rows, err := db.Query("SELECT * FROM public.ispace LIMIT $1", 10)
	if err != nil {
		log.Fatalf("Error fetching rows: %v", err)
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
		log.Fatalf("Error scanning rows: %v", err)
	}

	return spaces, nil
}

// PsqlListSpacesWithTickets returns a list of spaces with corresponding tickets
func (*repo) PsqlListSpacesWithTickets() ([]entity.ISpace, error){
	// connecting to the DB
	db := PsqlConnect()
	// keeping the connection open
	defer db.Close()

	var spaces []entity.ISpace
	// query build
	query := "SELECT *\nFROM public.ispace JOIN public.iticket ON (space_id = ispace.uid);"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error executing a query: %v", err)
	}

	defer rows.Close()

	space := entity.ISpace{}
	ticket := entity.ITicket{}

	for rows.Next() {

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
			&ticket.Availability,
			&ticket.Colour,
			&ticket.Description,
			&ticket.Name,
			&ticket.Period,
			&ticket.Price,
			&ticket.SpaceID,
			&ticket.UID,
			)
		if err != nil {
			log.Fatalf("Error scanning rows in ITicket: %v", err)
		}
		space.Tickets = append(space.Tickets, ticket)
		spaces = append(spaces, space)
	}


	err = rows.Err()
	if err != nil {
		log.Fatalf("Error in the for loop: %v", err)
	}



	return spaces, nil
}

// PsqlCreateNewSpace writes a new ISpace record to the PostgreSQL DB
func (*repo) PsqlCreateNewSpace(space *entity.ISpace) (*entity.ISpace, error) {
	// connecting to the DB
	db := PsqlConnect()
	// keeping the connection open
	defer db.Close()

	// query build
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

// PsqlGetSpaceByID returns a single space as specified by the provided ID
func (*repo) PsqlGetSpaceByID(spaceID string) (entity.ISpace, error) {
	// connecting to the DB
	db := PsqlConnect()
	// keeping the connection open
	defer db.Close()

	var space entity.ISpace
	// query build
	rows, err := db.Query("SELECT ispace.address, ispace.availability, ispace.coordinates, ispace.description, ispace.image_urls, ispace.location, ispace.name, ispace.number_of_visitors, ispace.telephone_number, ispace.top_image_url, ispace.uid, ispace.visitor_greeting, ispace.visitor_slack_message, ispace.visitor_slack_webhook_url, ispace.website \nFROM public.ispace WHERE uid = $1", spaceID)
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
		log.Fatalf("Error PsqlGetSpaceByID: %v", err)
	}

	return space, nil
}

// PsqlCreateNewTicket writes a new ITicket record to the PostgreSQL DB
func (*repo) PsqlCreateNewTicket(ticket *entity.ITicket) (*entity.ITicket, error) {
	// connecting to the DB
	db := PsqlConnect()
	// keeping the connection open
	defer db.Close()

	// query build
	sqlStatement := `
		INSERT INTO public.iticket (availability, colour, description, name, period, price, space_id, uid)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING uid`
	id := 0
	err := db.QueryRow(sqlStatement, ticket.Availability, ticket.Colour, ticket.Description, ticket.Name, ticket.Period, ticket.Price, ticket.SpaceID, ticket.UID).Scan(&id)
	if err != nil {
		log.Fatalf("Failed to add a new ticket: %v", err)
		return nil, err
	}

	return ticket, nil
}