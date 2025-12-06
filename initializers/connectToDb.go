package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("CONN")
	fmt.Println("DSN: ", dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("gorm err: ", err)
		panic("Failed to connect to DB")
	}
	fmt.Println("Connected to DB successfully: ", DB)
}
