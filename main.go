package main

import (
	"cbam_api/data"
	"cbam_api/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	loadEnv()

	mongoURI := fmt.Sprintf("mongodb://%s:%s@mongodb:27017/%s?authSource=admin", os.Getenv("MONGO_ROOT_USERNAME"), os.Getenv("MONGO_ROOT_PASSWORD"), os.Getenv("DB_NAME"))

	client, err := data.SetupDatabase(mongoURI)
	if err != nil {
		log.Fatalf("Failed to set up MongoDB database: %v", err)
	}
	defer client.Client.Disconnect(context.Background())

	server(client)
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file !!", err)
	}
}

func server(client *data.MongoDBClient) {
	router := routes.NewRouter()
	routes.QuarterlyReportGroup(router)
	router.ListRoutes()
	port := routes.GetServerPort()
	log.Printf("CBAM API is Running on PORT %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
