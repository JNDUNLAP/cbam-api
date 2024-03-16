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
	http.HandleFunc(
		"/quarterly-reports/:id/imported-goods/:goodId/emissions",
		emissions.GetGoodsEmissions(client),
	)

}

func ImportedGoodsGroup(client *data.MongoDBClient) {
	http.HandleFunc(
		"/quarterly-reports/:id/imported-goods",
		goods.GetImportedGoods(client),
	)

}

func ImporterGroup(client *data.MongoDBClient) {
	http.HandleFunc(
		"/quarterly-reports/:id/importer",
		importer.GetImporter(client),
	)
}

func RemarksEmissionsGroup(client *data.MongoDBClient) {
	http.HandleFunc(
		"/quarterly-reports/:reportID/imported-goods/:goodID/remarks-emissions",
		remarks.GetRemarksEmissions(client),
	)
}

func NationalCompetentAuthGroup(client *data.MongoDBClient) {
	http.HandleFunc(
		"/quarterly-reports/:id/national-competent-auth",
		authority.GetNationalCompetentAuth(client),
	)
}

func SignaturesGroup(client *data.MongoDBClient) {
	http.HandleFunc(
		"/quarterly-reports/:id/signatures",
		signatures.GetSignatures(client),
	)
}

func SupportingDocumentsGroup(client *data.MongoDBClient) {
	http.HandleFunc(
		"/quarterly-reports/:reportID/imported-goods/:goodID/supporting-documents",
		documents.GetSupportingDocuments(client),
	)
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
