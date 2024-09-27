package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB

func initDB() {
	var err error
	cfg := "user=postgres password=Adilek2003alimgozha dbname=golab3 host=localhost sslmode=disable"
	// Get a database handle.
	db, err = sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func getUsers(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.IndentedJSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.QueryRow("INSERT INTO users(name, age) VALUES($1, $2) RETURNING id", newUser.Name, newUser.Age).Scan(&newUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

func updateUser(c *gin.Context) {
	idStr := c.Param("id") // Получаем id как строку
	var updatedUser User

	// Преобразуем строку в uint
	id, err := strconv.ParseUint(idStr, 10, 32) // Преобразуем в 32-битный беззнаковый целочисленный тип
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Используем преобразованный id
	_, err = db.Exec("UPDATE users SET name = $1, age = $2 WHERE id = $3", updatedUser.Name, updatedUser.Age, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedUser.ID = int(id) // Присваиваем id обратно в структуру как uint
	c.IndentedJSON(http.StatusOK, updatedUser)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent) // No content response for successful deletion
}

func main() {
	initDB()
	router := gin.Default()

	// Define routes
	router.GET("/users", getUsers)
	router.POST("/user", createUser)
	router.PUT("/user/:id", updateUser)
	router.DELETE("/user/:id", deleteUser)

	router.Run(":8080") // Start the server on port 8080
}
