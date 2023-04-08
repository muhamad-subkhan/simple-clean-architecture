package config

import (
	"clean/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func Database() *gorm.DB {


	err := godotenv.Load()
	if err == nil {
		fmt.Println("environment Successfully Load")
	} else {
		log.Fatal("environment Load Failed")
	}

	PGHOST := os.Getenv("PGHOST")
	PGUSER := os.Getenv("PGUSER")
	PGPASSWORD := os.Getenv("PGPASSWORD")
	PGDATABASE := os.Getenv("PGDATABASE")
	PGPORT := os.Getenv("PGPORT")


	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", PGHOST, PGUSER, PGPASSWORD, PGDATABASE, PGPORT)
	DB, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't open database")
	}

	DB.Debug().AutoMigrate(
		models.User{},
	)

	return DB
}