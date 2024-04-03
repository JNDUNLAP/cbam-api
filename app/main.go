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

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file !!", err)
	}
}

func server(client *data.MongoDBClient) {
	router := routes.NewRouter()
	routes.QuarterlyReportGroup(router, client)
	routes.DeclarantGroup(router, client)
	routes.NationalCompetentAuthGroup(router, client)
	routes.ImporterGroup(router, client)
	routes.ImportedGoodsGroup(router, client)
	routes.SignaturesGroup(router, client)
	routes.GoodsEmissionsGroup(router, client)
	routes.SupportingDocumentsGroup(router, client)
	routes.FileGroup(router, client)
	router.ListRoutes()
	port := routes.GetServerPort()
	log.Printf("CBAM API is Running on PORT %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	loadEnv()
	mongoURI := fmt.Sprintf("mongodb://%s:%s@mongodb:27017/%s?authSource=admin",
		os.Getenv("MONGO_ROOT_USERNAME"),
		os.Getenv("MONGO_ROOT_PASSWORD"),
		os.Getenv("DB_NAME"))

	fmt.Println(mongoURI)
	client, err := data.SetupDatabase(mongoURI)
	if err != nil {
		log.Fatalf("Failed to set up MongoDB database: %v", err)
	}
	defer client.Client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Failed to set up MongoDB database: %v", err)
	}
	server(client)
}
