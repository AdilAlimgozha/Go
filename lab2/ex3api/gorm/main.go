package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

var db *gorm.DB
var err error

func initDB() {
	dsn := "host=localhost user=postgres password=Adilek2003alimgozha dbname=golab3 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
}

func getUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.IndentedJSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&newUser)
	c.IndentedJSON(http.StatusOK, newUser)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&user)
	c.IndentedJSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	db.Delete(&user)
	c.Status(http.StatusNoContent)
}

func main() {

	initDB()

	router := gin.Default()

	router.GET("/users", getUsers)
	router.POST("/user", createUser)
	router.PUT("/user/:id", updateUser)
	router.DELETE("/user/:id", deleteUser)

	router.Run(":8080")
}
