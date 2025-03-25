package config

import (
	"defect-service/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=db user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("✅ Database Connected Successfully!")

	err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		log.Fatal("❌ Failed to create uuid-ossp extension:", err)
	}

	err = db.AutoMigrate(
		&models.Defects{},
	)

	if err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	} else {
		fmt.Println("✅ AutoMigrate completed!")
	}

	DB = db
}
