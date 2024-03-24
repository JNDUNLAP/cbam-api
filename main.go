package main

import (
	"cbam_api/data"
	"cbam_api/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
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
	client := data.SetupDatabase(os.Getenv("MONGODB_URI"), os.Getenv("DB_NAME"))
	server(client)
}
