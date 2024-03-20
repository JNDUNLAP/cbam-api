package main

import (
	"dunlap/data"
	"dunlap/routes"
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
	// http.HandleFunc("/api/ping", routes.PingHandler)
	// http.HandleFunc("/api/upload", routes.UploadHandler(client))
	router := routes.NewRouter()
	routes.QuarterlyReportGroup(router, client)
	routes.DeclarantGroup(router, client)
	routes.NationalCompetentAuthGroup(router, client)
	routes.ImporterGroup(router, client)
	routes.ImportedGoodsGroup(router, client)
	routes.SignaturesGroup(router, client)
	routes.GoodsEmissionsGroup(router, client)
	routes.SupportingDocumentsGroup(router, client)
	router.ListRoutes()
	port := routes.GetServerPort()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	loadEnv()
	client := data.SetupDatabase(os.Getenv("URI"), os.Getenv("DB_NAME"))
	server(client)
}
