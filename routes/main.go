package routes

import (
	"dunlap/data"
	"dunlap/files"
	"dunlap/report/authority"
	"dunlap/report/documents"
	"dunlap/report/emissions"
	"dunlap/report/goods"
	"dunlap/report/importer"
	"dunlap/report/remarks"
	"dunlap/report/reports"
	"dunlap/report/signatures"
	"net/http"
)

func QuarterlyReportGroup(client *data.MongoDBClient) {

	http.HandleFunc(
		"/quarterly-reports/:reportID",
		reports.GetQuarterlyReport(client),
	)

	http.HandleFunc(
		"/quarterly-reports",
		reports.GetAllQuarterlyReports(client),
	)

	http.HandleFunc(
		"/quarterly-reports/create",
		reports.CreateQuarterlyReport(client),
	)

	http.HandleFunc("/quarterly-reports/delete:id",
		reports.DeleteAllQuarterlyReports(client))
}

func GoodsEmissionsGroup(client *data.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:id/imported-goods/:goodId/emissions", emissions.GetGoodsEmissions(client))
	// e.PUT("/quarterly-reports/:id/imported-goods/:goodId/emissions", emissions.UpdateGoodsEmissions)
}

func ImportedGoodsGroup(client *data.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:id/imported-goods", goods.GetImportedGoods(client))
	// e.POST("/quarterly-reports/:id/imported-goods", goods.AddImportedGood)
	// e.PUT("/quarterly-reports/:id/imported-goods/:goodId", goods.UpdateImportedGood)
	// e.DELETE("/quarterly-reports/:id/imported-goods/:goodId", goods.DeleteImportedGood)
}

func ImporterGroup(client *data.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:id/importer", importer.GetImporter(client))
	// e.PUT("/quarterly-reports/:id/importer", importer.UpdateImporter)
}

func RemarksEmissionsGroup(client *data.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:reportID/imported-goods/:goodID/remarks-emissions", remarks.GetRemarksEmissions(client))
	// e.PUT("/quarterly-reports/:reportID/imported-goods/:goodID/remarks-emissions", remarks.UpdateRemarksEmissions)
}

func NationalCompetentAuthGroup(client *data.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:id/national-competent-auth", authority.GetNationalCompetentAuth(client))
	// e.PUT("/quarterly-reports/:id/national-competent-auth", signatures.UpdateNationalCompetentAuth)
}

func SignaturesGroup(client *data.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:id/signatures", signatures.GetSignatures(client))
	// e.PUT("/quarterly-reports/:id/signatures", signatures.UpdateSignatures)
}

func SupportingDocumentsGroup(client *data.MongoDBClient) {
	http.HandleFunc("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents", documents.GetSupportingDocuments(client))
	// // e.POST("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents", documents.AddSupportingDocument)
	// // e.GET("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents/:docID", documents.GetSupportingDocument)
	// // e.PUT("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents/:docID", documents.UpdateSupportingDocument)
	// // e.DELETE("/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents/:docID", documents.DeleteSupportingDocument)
}

func UploadHandler(client *data.MongoDBClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		report, err := files.XML(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		RespondWithJSON(w, http.StatusOK, report)
	}
}
