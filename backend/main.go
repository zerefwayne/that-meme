package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/zerefwayne/that-meme/config"
	"github.com/zerefwayne/that-meme/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	config.Config.ConnectDatabase()

	config.Config.ConnectCache()
	defer config.Config.Cache.Close()

	config.Config.ConnectS3()

	handler := routes.NewHandler()

	log.Fatal(http.ListenAndServe(":5000", handler))

}
