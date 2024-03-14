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
	http.HandleFunc("/api/ping", routes.PingHandler)
	http.HandleFunc("/api/upload", routes.UploadHandler(client))
	routes.QuarterlyReportGroup(client)
	routes.SupportingDocumentsGroup(client)

	port := routes.GetServerPort()
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	loadEnv()

	client := data.SetupDatabase(os.Getenv("URI"), os.Getenv("DB_NAME"))
	server(client)
}
