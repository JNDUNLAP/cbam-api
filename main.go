package main

import (
	"dunlap/app/declarant"
	"dunlap/app/documents"
	"dunlap/app/emissions"
	"dunlap/app/goods"
	"dunlap/app/importer"
	"dunlap/app/remarks"
	"dunlap/app/reports"
	"dunlap/app/signatures"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	e := echo.New()
	// e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))

	e.GET("/api/ping", func(c echo.Context) error { return c.String(http.StatusOK, "Pong") })

	e.GET("/quarterly-reports", reports.GetQuarterlyReports)
	e.GET("/quarterly-reports/:id", reports.GetQuarterlyReport)
	e.POST("/quarterly-reports", reports.CreateQuarterlyReport)
	e.PUT("/quarterly-reports/:id", reports.UpdateQuarterlyReport)
	e.DELETE("/quarterly-reports/:id", reports.DeleteQuarterlyReport)

	e.GET("/quarterly-reports/:id/declarant", declarant.GetDeclarant)
	e.PUT("/quarterly-reports/:id/declarant", declarant.UpdateDeclarant)

	// Importer routes
	e.GET("/quarterly-reports/:id/importer", importer.GetImporter)
	e.PUT("/quarterly-reports/:id/importer", importer.UpdateImporter)

	// NationalCompetentAuth routes
	e.GET("/quarterly-reports/:id/national-competent-auth", signatures.GetNationalCompetentAuth)
	e.PUT("/quarterly-reports/:id/national-competent-auth", signatures.UpdateNationalCompetentAuth)
	e.GET("/quarterly-reports/:id/signatures", signatures.GetSignatures)
	e.PUT("/quarterly-reports/:id/signatures", signatures.UpdateSignatures)

	// ImportedGood routes
	e.GET("/quarterly-reports/:id/imported-goods", goods.GetImportedGoods)
	e.POST("/quarterly-reports/:id/imported-goods", goods.AddImportedGood)
	e.PUT("/quarterly-reports/:id/imported-goods/:goodId", goods.UpdateImportedGood)
	e.DELETE("/quarterly-reports/:id/imported-goods/:goodId", goods.DeleteImportedGood)

	// GoodsEmissions routes
	e.GET("/quarterly-reports/:id/imported-goods/:goodId/emissions", emissions.GetGoodsEmissions)
	e.PUT("/quarterly-reports/:id/imported-goods/:goodId/emissions", emissions.UpdateGoodsEmissions)

	// SupportingDocuments routes
	e.GET("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents", documents.GetSupportingDocuments)
	e.POST("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents", documents.AddSupportingDocument)
	e.GET("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents/:docID", documents.GetSupportingDocument)
	e.PUT("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents/:docID", documents.UpdateSupportingDocument)
	e.DELETE("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents/:docID", documents.DeleteSupportingDocument)

	// RemarksEmissions routes
	e.GET("/quarterly-reports/:reportID/imported-goods/:goodID/remarks-emissions", remarks.GetRemarksEmissions)
	e.PUT("/quarterly-reports/:reportID/imported-goods/:goodID/remarks-emissions", remarks.UpdateRemarksEmissions)

	e.Logger.Fatal(e.Start(PORT))
}
