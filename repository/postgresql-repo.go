package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"glue/glue-backend-golang/entity"
)

//NewPostgreRepository creates a new repository
func NewPostgreRepository() ISpaceRepository {
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

func (*repo) ListSpacesPostgre() ([]entity.ISpace, error){
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

	rows, err := db.Query("SELECT * FROM ISpace LIMIT $1", 1)
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