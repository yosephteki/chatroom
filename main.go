package main

import (
	"chatroom/Config"
	"chatroom/Routes"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error load .env file")
	}

	Config.DB, err = gorm.Open("postgres", Config.DBURL(Config.BuildConfig()))
	if err != nil {
		log.Println("DB error : ", err)
	}

	defer Config.DB.Close()

	r := Routes.SetupRouter()
	r.Run()
}
