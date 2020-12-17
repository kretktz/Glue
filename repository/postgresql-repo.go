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

//TODO: Fix the SQL error "pq: invalid reference to FROM-clause entry for table"
func (*repo) ListSpacesWithTicketsPsql() ([]entity.ISpace, []entity.ITicket, error){
	db := PsqlConnect()

	defer db.Close()

	var spaces []entity.ISpace
	//var tickets []entity.ITicket

	query := "SELECT public.ispace.name, public.ispace.uid, public.iticket.space_id, public.iticket.name, public.iticket.price \nFROM public.ispace s\nJOIN public.iticket t ON s.uid = t.space_id\nWHERE s.uid = $1"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error executing a query: %v", err)
	}

	defer rows.Close()

	space := &entity.ISpace{}

	for rows.Next() {
		ticket := entity.ITicket{}
		err = rows.Scan(
			&space.Name,
			&space.UID,
			&ticket.SpaceID,
			&ticket.Name,
			&ticket.Price,
		)
		if err != nil {
			panic(err)
		}
		space.Tickets = append(space.Tickets, ticket)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return spaces, space.Tickets, nil
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
			&space.Tickets,
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

func (*repo) CreateNewTicketPsql(ticket *entity.ITicket) (*entity.ITicket, error) {
	db := PsqlConnect()
	defer db.Close()


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