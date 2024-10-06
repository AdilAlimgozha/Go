package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Users struct {
	Id   int64
	Name string
	Age  int64
}

// Table creating func
func createTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT,
        age INT
    );`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table Created")
}

// Data insertion func
func insertData(db *sql.DB, name string, age int) {
	query := `INSERT INTO users (name, age) VALUES ($1, $2)`

	_, err := db.Exec(query, name, age)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User inserted succesfully")
}

// Select all
func selectAll(db *sql.DB) {

	var users []Users

	query := `SELECT * FROM users`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal()
	}

	defer rows.Close()

	for rows.Next() {
		var user Users

		if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal()
	}

	fmt.Println(users)
}

func main() {
	// Connection
	cfg := "user=postgres password=Adilek2003alimgozha dbname=golab3 host=localhost sslmode=disable"
	// Get a database handle.
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// Creating table
	createTable(db)

	// Inserting data
	insertData(db, "Max", 25)
	insertData(db, "Peter", 27)

	// Secelting all
	selectAll(db)
}
