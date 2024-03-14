package main

import (
	"dunlap/database"
	"dunlap/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	client := database.SetupDatabase()
	server(client)
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func server(client *database.MongoDBClient) {
	http.HandleFunc("/api/ping", routes.PingHandler)
	http.HandleFunc("/api/upload", routes.UploadHandler(client))
	routes.SupportingDocumentsGroup(client)
	// http.HandleFunc("/quarterly-reports/:id/declarant", declarant.UpdateDeclarant)
	port := routes.GetServerPort()
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
