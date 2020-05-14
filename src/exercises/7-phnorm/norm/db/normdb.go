package db

import (
	"database/sql"
	"fmt"
	"strings"

	// have to import this so that the package gets initialized properly
	_ "github.com/lib/pq"
)

const (
	dbtype = "postgres"

	basicCon = "host=%s port=%d user=%s password=%s sslmode=disable"
	dbCon    = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"

	host     = "localhost"
	port     = 5432
	user     = "dev"
	pass     = "devpass"
	database = "go_exercises_7"

	createPNTableSQL = `
	CREATE TABLE IF NOT EXISTS phone_numbers (
	  id SERIAL,
	  number VARCHAR(255)
	)
	`

	insertPNTableSQL = `
	INSERT INTO phone_numbers (number) VALUES ($1) RETURNING id
	`
	selectPNsSQL = `
	SELECT id, number FROM phone_numbers
	`
	updatePNsSQL = `
	UPDATE phone_numbers SET number=$1 WHERE id=$2
	`
	deletePNsSQL = `
	DELETE FROM phone_numbers WHERE id=$1
	`

	testPNNumbers = `

	1234567890
	123 456 7891
	(123) 456 7892
	(123) 456-7893
	123-456-7894
	123-456-7890
	1234567892
	(123)456-7892

	1234567890
	123 456 7891
	(123) 456 7892
	(123) 456-7893
	123-456-7894
	123-456-7890
	1234567892
	(123)456-7892

	
	`
)

type PhoneNumber struct {
	Key       int
	Number    string
	Modified  bool
	Duplicate bool
}

func GetAllPNs(db *sql.DB) []PhoneNumber {
	// db := connectTo()
	rows, err := db.Query(selectPNsSQL)
	must(err)
	ret := make([]PhoneNumber, 0)
	for rows.Next() {
		next := PhoneNumber{}
		rows.Scan(&next.Key, &next.Number)
		ret = append(ret, next)
	}
	defer rows.Close()
	// defer db.Close()
	return ret
}

func UpdatePNs(pns []PhoneNumber, db *sql.DB) {
	// db := connectTo()
	stmt, err := db.Prepare(updatePNsSQL)
	must(err)
	for _, pn := range pns {
		if pn.Modified {
			_, err = stmt.Exec(pn.Number, pn.Key)
			must(err)
		}
	}
	defer stmt.Close()
	// defer db.Close()
}

// TODO this function only eliminates duplicates in the current set, not across the database
func RemoveDuplicatePNs(pns []PhoneNumber, db *sql.DB) {
	// db := connectTo()
	stmt, err := db.Prepare(deletePNsSQL)
	must(err)
	for _, pn := range pns {
		if pn.Duplicate {
			_, err = stmt.Exec(pn.Key)
			must(err)
		}
	}
	defer stmt.Close()
	// defer db.Close()
}

func connectTo() *sql.DB {
	db, err := sql.Open(dbtype, fmt.Sprintf(dbCon, host, port, user, pass, database))
	must(err)
	return db
}

func SetupDB() *sql.DB {
	psqlInfo := fmt.Sprintf(basicCon, host, port, user, pass)
	db, err := sql.Open(dbtype, psqlInfo)
	must(err)
	err = resetDB(db, database)
	must(err)
	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, database)
	db, err = sql.Open(dbtype, psqlInfo)
	must(err)
	must(createPNTable(db))
	testnums := strings.Split(testPNNumbers, "\n")
	for _, num := range testnums {
		num = strings.TrimSpace(num)
		if num != "" {
			_, err := insertPN(db, num)
			must(err)
		}
	}
	return db
}

func insertPN(db *sql.DB, pn string) (int, error) {
	var id int
	err := db.QueryRow(insertPNTableSQL, pn).Scan(&id)
	must(err)
	// DOESN'T WORK FOR POSTGRES?
	//id, err := res.LastInsertId()
	return int(id), err
}

func createPNTable(db *sql.DB) error {
	_, err := db.Exec(createPNTableSQL)
	return err
}

func createDB(db *sql.DB, name string) error {
	// DANGEROUS SQL INJECTION
	_, err := db.Exec("CREATE DATABASE " + name)
	must(err)
	return nil
}

func resetDB(db *sql.DB, name string) error {
	// DANGEROUS SQL INJECTION
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	must(err)
	return createDB(db, name)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
