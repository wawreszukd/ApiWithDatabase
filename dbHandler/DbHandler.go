package dbHandler

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "password"
	Dbname   = "test"
)

type MyDb struct {
	Db *sql.DB
}
type Person struct {
	ID          int
	FirstName   string
	Lastname    string
	Gender      string
	DateOfBirth string
}

func InitiateDataBase() (*MyDb, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Database connected")
	return &MyDb{Db: db}, nil
}
func (db MyDb) HandleSelect() (error, string) {
	sqlStatement := "SELECT * FROM person;"

	rows, err := db.Db.Query(sqlStatement)
	if err != nil {
		return err, ""
	}

	defer rows.Close()

	var person Person
	var Fullout string
	for rows.Next() {
		err := rows.Scan(&person.ID, &person.FirstName, &person.Lastname, &person.Gender, &person.DateOfBirth)
		if err != nil {
			return err, ""
		}
		outstring := fmt.Sprintf("%d %s %s %s %s \n", person.ID, person.FirstName, person.Lastname, person.Gender, person.DateOfBirth)
		Fullout += outstring
	}

	return nil, Fullout
}
func (db MyDb) HandleInsert(firstname, lastname, gender, date string) {
	sqlStatement := `
	INSERT INTO person (first_name, last_name, gender, date_of_birth)
	VALUES ($1, $2, $3, $4)`
	_, err := db.Db.Exec(sqlStatement, firstname, lastname, gender, date)
	if err != nil {
		log.Fatal(err)
	}
}
