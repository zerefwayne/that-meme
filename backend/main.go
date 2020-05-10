package main

import (
	"log"
	"net/http"
	"os"

	"github.com/zerefwayne/that-meme/config"
	"github.com/zerefwayne/that-meme/routes"
)

func init() {
	config.Config.LoadEnv()

	config.Config.ConnectDatabase()
	config.Config.ConnectS3()
	config.Config.ConnectElasticSearch()
}

func generatePort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":5000"
	}
	return ":" + port
}

func main() {

	listenAddress := generatePort()

	handler := routes.NewHandler()

	log.Fatal(http.ListenAndServe(listenAddress, handler))

}
