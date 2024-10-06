package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB
var err error

func initDB() {
	cfg := "user=postgres password=Adilek2003alimgozha dbname=golab3 host=localhost sslmode=disable"
	db, err = sql.Open("postgres", cfg)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(20)                 // Set the maximum number of open connections
	db.SetMaxIdleConns(5)                  // Set the maximum number of idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Set the maximum lifetime of a connection

	// Ensure the database connection is valid
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// Table creating func
func createTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT UNIQUE NOT NULL,
        age INT NOT NULL
    );`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// Function to insert multiple users within a transaction
func insertUsersWithTransaction(users []User) error {
	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}

	// Prepare the insert statement
	stmt, err := tx.Prepare("INSERT INTO users (name, age) VALUES ($1, $2)")
	if err != nil {
		tx.Rollback() // Rollback if preparing the statement fails
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()

	// Loop over users and execute the statement for each
	for _, user := range users {
		_, err := stmt.Exec(user.Name, user.Age)
		if err != nil {
			tx.Rollback() // Rollback if an error occurs during insertion
			return fmt.Errorf("error inserting user %v: %v", user.Name, err)
		}
	}

	// Commit the transaction if all insertions are successful
	err = tx.Commit()
	if err != nil {
		tx.Rollback() // Rollback if committing the transaction fails
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func QueryUsers(db *sql.DB, ageFilter *int, limit, offset int) ([]User, error) {
	query := "SELECT id, name, age FROM users WHERE 1=1"
	args := []interface{}{}

	// Apply age filter if provided
	if ageFilter != nil {
		query += " AND age = $1"
		args = append(args, *ageFilter)
	}

	// Apply pagination
	query += " LIMIT $2 OFFSET $3"
	args = append(args, limit, offset)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// UpdateUser updates a user's details (name and age) by their ID.
func UpdateUser(db *sql.DB, userID int, newName string, newAge int) error {
	query := "UPDATE users SET name = $1, age = $2 WHERE id = $3"
	result, err := db.Exec(query, newName, newAge, userID)
	if err != nil {
		return fmt.Errorf("could not update user: %v", err)
	}

	// Check if any rows were affected (i.e., user exists)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get affected rows: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with ID: %d", userID)
	}

	fmt.Printf("User with ID %d updated successfully\n", userID)
	return nil
}

// DeleteUser deletes a user by their ID.
func DeleteUser(db *sql.DB, userID int) error {
	query := "DELETE FROM users WHERE id = $1"
	result, err := db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("could not delete user: %v", err)
	}

	// Check if any rows were affected (i.e., user exists)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get affected rows: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with ID: %d", userID)
	}

	fmt.Printf("User with ID %d deleted successfully\n", userID)
	return nil
}

func main() {
	initDB()

	createTable(db)

	new_users := []User{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 22},
	}

	// Insert users with transaction
	err := insertUsersWithTransaction(new_users)
	if err != nil {
		log.Fatalf("Failed to insert users: %v", err)
	} else {
		fmt.Println("Users inserted successfully")
	}

	// Example: Query users with age filter and pagination (age = 30, page 1, 5 results per page)
	ageFilter := 30
	limit := 5
	page := 1
	offset := (page - 1) * limit

	users, err := QueryUsers(db, &ageFilter, limit, offset)
	if err != nil {
		log.Fatal(err)
	}

	// Print users
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.Id, user.Name, user.Age)
	}

	// Example: Update user with ID 1
	err = UpdateUser(db, 3, "New Name", 28)
	if err != nil {
		log.Printf("Error updating user: %v\n", err)
	}

	// Example: Delete user with ID 2
	err = DeleteUser(db, 3)
	if err != nil {
		log.Printf("Error deleting user: %v\n", err)
	}

}
