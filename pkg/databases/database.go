package databases

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env")
	}

	port, _ := strconv.Atoi(os.Getenv("APP_DB_PORT"))

	var (
		host = os.Getenv("APP_DB_HOST")
		name = os.Getenv("APP_DB_NAME")
		user = os.Getenv("APP_DB_USER")
		pass = os.Getenv("APP_DB_PASS")
	)

	sqlinfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", host, port, name, user, pass)

	db, err := gorm.Open(postgres.Open(sqlinfo), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	return db
}
