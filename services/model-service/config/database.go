package config

import (
	"fmt"
	"log"
	"model-service/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func loadEnvVariables() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
}

func getDatabaseConfig() (string, string, string, string, string) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("❌ One or more database connection details are not set in .env file")
	}

	return host, port, user, password, dbname
}

func ConnectDB() {
	loadEnvVariables()

	host, port, user, password, dbname := getDatabaseConfig()

	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Database Connected Successfully!")

	err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		log.Fatal("❌ Failed to create uuid-ossp extension:", err)
	}

	err = db.AutoMigrate(
		&models.Models{},
	)

	if err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	} else {
		fmt.Println("✅ AutoMigrate completed!")
	}

	DB = db
}

// func InitDB() {
// 	var err error
// 	dsn := os.Getenv("DATABASE_URL")
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}

// 	// Run migrations
// 	err = DB.AutoMigrate(&models.Brands{})
// 	if err != nil {
// 		log.Fatalf("Failed to migrate database: %v", err)
// 	}
// }
