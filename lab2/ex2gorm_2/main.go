package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// User model
type User struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserID"` // One-to-one association with Profile
}

// Profile model
type Profile struct {
	ID                uint   `gorm:"primaryKey"`
	UserID            uint   // Foreign key to associate with User
	Bio               string `json:"bio"`
	ProfilePictureURL string `json:"profile_picture_url"`
}

func (User) TableName() string {
	return "users"
}

var db *gorm.DB
var err error

func initDB() {
	// Connection to postgres
	dsn := "host=localhost user=postgres password=Adilek2003alimgozha dbname=golab3 port=5432 sslmode=disable TimeZone=Asia/Almaty"

	// Gorm connection
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Logging sql queries
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("Error in DB connection: %v", err)
	}

	// Retrieve the SQL connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error in sql connection: %v", err)
	}

	// Polling settings connectinos
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Success")
}

func create_tables() {
	// Auto-migrate the schema (create the tables if they don't exist)
	err = db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate the schema: %v", err)
	}
}

func insert() {
	// Insert User and Profile in a transaction
	err = db.Transaction(func(tx *gorm.DB) error {
		// Create a new User
		newUser := User{Name: "Adil Aidil", Age: 30}
		if err := tx.Create(&newUser).Error; err != nil {
			fmt.Println("Error inserting user:", err)
			return err // Rollback if there's an error
		}

		// Create a new Profile associated with the User
		newProfile := Profile{UserID: newUser.ID, Bio: "Software Developer", ProfilePictureURL: "http://example.com/pic.jpg"}
		if err := tx.Create(&newProfile).Error; err != nil {
			fmt.Println("Error inserting profile:", err)
			return err // Rollback if there's an error
		}

		return nil // Commit the transaction
	})

	if err != nil {
		log.Println("Transaction failed:", err)
	} else {
		log.Println("User and Profile inserted successfully!")
	}

}

func queryUsersWithProfiles() {
	var users []User
	if err := db.Preload("Profile").Find(&users).Error; err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		log.Printf("User: %s, Age: %d, Bio: %s, ProfilePictureURL: %s\n", user.Name, user.Age, user.Profile.Bio, user.Profile.ProfilePictureURL)
	}
}

func updateUserProfile(userID uint, newProfile Profile) error {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}

	// Update profile
	user.Profile.Bio = newProfile.Bio
	user.Profile.ProfilePictureURL = newProfile.ProfilePictureURL

	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func deleteUser(userID uint) error {
	var user User
	if err := db.Preload("Profile").First(&user, userID).Error; err != nil {
		return err
	}

	// Delete profile
	if err := db.Delete(&user.Profile).Error; err != nil {
		return err
	}

	// Delete user
	if err := db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func main() {

	initDB()

	create_tables()

	insert()

	queryUsersWithProfiles()

	newProfile := Profile{Bio: "Updated bio", ProfilePictureURL: "http://example.com/newpic.jpg"}
	if err := updateUserProfile(1, newProfile); err != nil {
		log.Println("Error updating profile:", err)
	} else {
		log.Println("User profile updated successfully.")
	}

	if err := deleteUser(1); err != nil {
		log.Println("Error deleting user:", err)
	} else {
		log.Println("User deleted successfully.")
	}
}
