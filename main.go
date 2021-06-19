package main

import (
	"chatroom/Config"
	"chatroom/Routes"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error load .env file")
	}
	Config.DB, err = gorm.Open(postgres.Open(Config.DBURL(Config.BuildConfig())), &gorm.Config{})
	if err != nil {
		log.Println("DB error : ", err)
	}

	sqlDB, err := Config.DB.DB()

	defer sqlDB.Close()

	r := Routes.SetupRouter()
	r.Run()
}
