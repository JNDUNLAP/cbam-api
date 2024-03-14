package routes

import (
	"dunlap/database"
	"dunlap/report/authority"
	"dunlap/report/documents"
	"dunlap/report/emissions"
	"dunlap/report/goods"
	"dunlap/report/importer"
	"dunlap/report/remarks"
	"dunlap/report/signatures"
	"dunlap/upload"
	"net/http"
)

func QuarterlyReportGroup(client *database.MongoDBClient) {
	// e.GET("/quarterly-reports/:reportID", reports.GetQuarterlyReports(client))
	// e.GET("/quarterly-reports", reports.GetQuarterlyReports)
	// e.POST("/quarterly-reports", reports.CreateQuarterlyReport)
	// e.PUT("/quarterly-reports/:id", reports.UpdateQuarterlyReport)
	// e.DELETE("/quarterly-reports/:id", reports.DeleteQuarterlyReport)
}

func GoodsEmissionsGroup(client *database.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:id/imported-goods/:goodId/emissions", emissions.GetGoodsEmissions(client))
	// e.PUT("/quarterly-reports/:id/imported-goods/:goodId/emissions", emissions.UpdateGoodsEmissions)
}

func ImportedGoodsGroup(client *database.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:id/imported-goods", goods.GetImportedGoods(client))
	// e.POST("/quarterly-reports/:id/imported-goods", goods.AddImportedGood)
	// e.PUT("/quarterly-reports/:id/imported-goods/:goodId", goods.UpdateImportedGood)
	// e.DELETE("/quarterly-reports/:id/imported-goods/:goodId", goods.DeleteImportedGood)
}

func ImporterGroup(client *database.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:id/importer", importer.GetImporter(client))
	// e.PUT("/quarterly-reports/:id/importer", importer.UpdateImporter)
}

func RemarksEmissionsGroup(client *database.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:reportID/imported-goods/:goodID/remarks-emissions", remarks.GetRemarksEmissions(client))
	// e.PUT("/quarterly-reports/:reportID/imported-goods/:goodID/remarks-emissions", remarks.UpdateRemarksEmissions)
}

func NationalCompetentAuthGroup(client *database.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:id/national-competent-auth", authority.GetNationalCompetentAuth(client))
	// e.PUT("/quarterly-reports/:id/national-competent-auth", signatures.UpdateNationalCompetentAuth)
}

func SignaturesGroup(client *database.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:id/signatures", signatures.GetSignatures(client))
	// e.PUT("/quarterly-reports/:id/signatures", signatures.UpdateSignatures)
}

func SupportingDocumentsGroup(client *database.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents", documents.GetSupportingDocuments(client))
	// // e.POST("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents", documents.AddSupportingDocument)
	// // e.GET("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents/:docID", documents.GetSupportingDocument)
	// // e.PUT("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents/:docID", documents.UpdateSupportingDocument)
	// // e.DELETE("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents/:docID", documents.DeleteSupportingDocument)
}

func UploadHandler(client *database.MongoDBClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		report, err := upload.XML(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		RespondWithJSON(w, http.StatusOK, report)
	}
}
