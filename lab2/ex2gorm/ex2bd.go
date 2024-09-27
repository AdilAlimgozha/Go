package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id   uint
	Name string
	Age  int64
}

func main() {

	dsn := "host=localhost user=postgres password=Adilek2003alimgozha dbname=golab3 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	// Creating table
	db.AutoMigrate(&User{})

	// Inserting data
	db.Create(&User{Name: "Alex", Age: 26})
	db.Create(&User{Name: "Adil", Age: 21})

	// Select all users
	var users []User
	result := db.Find(&users)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println(users)

}
